# REPORT

## 技術的決定

### 1. アーキテクチャ: 3層構造

**決定**: handler / service / repository の3層構造を採用

```
handler/    - HTTPリクエスト/レスポンス処理
service/    - ビジネスロジック（距離計算、フィルタリング）
repository/ - データアクセス（sqlc生成コードのラッパー）
```

**理由**:
- 各層の責務が明確に分離される
- テスタビリティが高い（各層をモック可能）
- 課題要件の「テスタビリティを意識した設計」に適合

**トレードオフ**:
- シンプルな課題に対してはやや過剰な構造かもしれない
- ファイル数が増える

---

### 2. 開発手法: TDD (Test-Driven Development)

**決定**: t-wadaさん方式のTDD（Red-Green-Refactor）で実装

**理由**:
- テストファーストにより、テスト可能な設計が自然と生まれる
- 小さなステップで確実に進められる
- リグレッションを防げる

---

### 3. 距離計算: Haversine公式

**決定**: 球面三角法に基づくHaversine公式を採用

```go
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
    const earthRadius = 6371.0 // km
    // ...
}
```

**理由**:
- 地球を球体として近似し、2点間の大圏距離を計算
- 日本国内の充電ステーション検索では十分な精度
- 外部ライブラリ不要でシンプルに実装可能

**トレードオフ**:
- 地球の扁平率を考慮していない（Vincenty公式より精度が劣る）
- 極端に長い距離では誤差が大きくなる可能性がある

---

### 4. データベース設計

**決定**: 緯度経度をDECIMAL型で別々に保存

```sql
latitude DECIMAL(10, 7) NOT NULL,
longitude DECIMAL(11, 7) NOT NULL,
```

**理由**:
- 課題要件「緯度経度は数値として別々に保存」に準拠
- CSVデータの精度（小数点以下7桁）を保持
- SQLでの範囲検索やソートが可能

**トレードオフ**:
- PostGISのようなGIS拡張を使えば、より効率的な空間検索が可能だが、要件に沿ってGoでフィルタリングを実装

---

### 5. フィルタリングロジック: Goで実装

**決定**: 距離によるフィルタリングはGoで実装

**理由**:
- 課題要件「フィルタリングロジックは主にGoで実装」に準拠
- DBに依存しないテストが可能
- ビジネスロジックがService層に集約される

**トレードオフ**:
- 大量データの場合、全件取得後にフィルタリングするためパフォーマンスが低下する可能性
- 本番環境では、まず矩形範囲でSQLフィルタリングしてからGoで精密な距離計算をするハイブリッド方式が望ましい

---

### 6. date_from / date_to フィルタリング

**決定**: SQLレベルでフィルタリング

```sql
-- date_from: inclusive (>=)
WHERE last_updated >= ?

-- date_to: exclusive (<)
WHERE last_updated < ?
```

**理由**:
- OCPI仕様に準拠（date_from: inclusive, date_to: exclusive）
- 日時フィルタリングはSQLが効率的
- 距離フィルタリング前にデータ量を削減できる

---

### 7. CSVデータのインポート方式

**決定**: INSERTステートメントに変換してDockerの初期化スクリプトで実行

**理由**:
- `LOAD DATA LOCAL INFILE` はMySQL 8.0のセキュリティ制限により使用困難
- INSERTステートメントは確実に動作し、デバッグも容易
- docker-entrypoint-initdb.d による自動初期化

**トレードオフ**:
- 大量データの場合はバルクインサートやLOAD DATAの方が高速
- CSVファイルとSQLファイルの二重管理が発生

---

### 8. Status enum の文字列変換

**決定**: Service層でint→string変換を実装

```go
func statusToString(status int32) string {
    switch status {
    case 1: return "AVAILABLE"
    case 2: return "BLOCKED"
    // ...
    }
}
```

**理由**:
- OCPI仕様のステータス文字列をレスポンスで返す
- DBには数値で保存し、APIレスポンスでのみ文字列化
- 変換ロジックがService層に集約され、テスト可能

---

### 9. エラーレスポンスの形式

**決定**: 仕様書に定義がないため、独自にエラーレスポンス形式を設計

```go
// バリデーションエラー
c.JSON(http.StatusBadRequest, gin.H{"error": "latitude is required"})

// 内部エラー
c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
```

**HTTPステータスコード**:
- 200 OK: 成功
- 400 Bad Request: パラメータ不正（必須項目欠落、フォーマットエラー）
- 500 Internal Server Error: DB接続エラーなど内部エラー

**理由**:
- RESTful APIの一般的な慣習に従った
- エラーメッセージをJSONで返すことでクライアント側での処理が容易

---

### 10. DateTimeパースのフォールバック

**決定**: RFC3339に加え、タイムゾーンなし形式もサポート

```go
t, err := time.Parse(time.RFC3339, dateFromStr)
if err != nil {
    // タイムゾーンなし形式にフォールバック
    t, err = time.Parse("2006-01-02T15:04:05", dateFromStr)
}
```

**理由**:
- 仕様書のDateTime例に `2015-06-29T20:39:09`（タイムゾーンなし）が含まれている
- クライアントの利便性を考慮し、両形式を受け付ける
- タイムゾーンなしの場合はUTCとして扱う（OCPI仕様に準拠）

---

### 11. radius の型拡張

**決定**: 仕様書では `int` だが、実装では `float64` を受け付ける

**理由**:
- より柔軟な検索半径の指定が可能（例: 0.5km）
- `int` への暗黙的な変換も可能なため後方互換性あり

---

### 12. EVSEインデックスの追加

**決定**: `location_id` にインデックスを追加

```sql
CREATE INDEX idx_evses_location_id ON evses(location_id);
```

**理由**:
- EVSE取得時の検索パフォーマンス向上
- LocationとEVSEの結合クエリが頻繁に発生するため必須

---

### 13. latitude/longitude のバリデーション

**決定**: 仕様書のregexパターンによる厳密な検証は省略し、数値変換と地理的範囲チェックで検証

```go
lat, err := strconv.ParseFloat(latStr, 64)
if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "invalid latitude format"})
    return
}
if lat < -90 || lat > 90 {
    c.JSON(http.StatusBadRequest, gin.H{"error": "latitude must be between -90 and 90"})
    return
}
```

**仕様書のregex**: `-?[0-9]{1,2}\.[0-9]{5,7}`（小数点以下5-7桁）

**理由**:
- 実用上、数値として解釈できれば問題ない
- 過度に厳密なバリデーションはクライアントの利便性を損なう
- 地理的に無効な座標（緯度90超、経度180超）は明らかなエラーなので範囲チェックを追加

**トレードオフ**:
- 仕様書のregexに完全準拠していない
- 不正な精度の座標も受け付けてしまう（小数点以下5-7桁の制約なし）

---

## 問題点・改善案

### 1. N+1問題 (解決済み)

**問題**: 当初の実装では、ロケーションごとにEVSEを取得するためN+1クエリが発生していた。

```go
// Before: N+1クエリ
for _, loc := range locations {
    evses, err := s.repo.GetEVSEsByLocationID(ctx, loc.ID) // ループ内でクエリ実行
}
```

**解決策**: EVSEを一括取得してメモリ上でマッピングする方式に変更

```go
// After: 2クエリ
allEVSEs, err := s.repo.GetAllEVSEs(ctx) // 1回で全件取得

evseMap := make(map[int32][]repository.EVSEWithLocationID)
for _, e := range allEVSEs {
    evseMap[e.LocationID] = append(evseMap[e.LocationID], e)
}

for _, loc := range locations {
    locEVSEs := evseMap[loc.ID] // O(1)でアクセス
}
```

**効果**:
- クエリ数: 1 + N → 2 に削減
- ロケーション数が増えてもクエリ数は一定

**トレードオフ**:
- 全EVSEをメモリに保持するため、データ量が極端に多い場合はメモリ使用量が増加
- 今回の課題規模では問題なし

### 2. 距離フィルタリングのパフォーマンス
全ロケーションを取得してからGoでフィルタリングしているため、データ量が増えると遅くなる。

**改善案**:
- まず矩形範囲（bounding box）でSQLフィルタリング
- その後Goで正確な距離計算

### 3. テストカバレッジ
現在のテストは主要なパスをカバーしているが、エッジケースのテストが不足。

**追加すべきテスト**:
- 境界値テスト（緯度-90/90、経度-180/180）
- 空のEVSEを持つロケーション
- 大量データでのパフォーマンステスト

---

## 参考

- [Go-Challenge Interface Specification Document](./Go-Challenge%20Interface%20Specification%20Document.pdf)
- [Haversine formula - Wikipedia](https://en.wikipedia.org/wiki/Haversine_formula)
- [OCPI 2.2.1 Specification](https://github.com/ocpi/ocpi)
