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

## 問題点・改善案

### 1. N+1問題
現在の実装では、ロケーションごとにEVSEを取得するためN+1クエリが発生する。

**改善案**:
- EVSEを一括取得してメモリ上でマッピング
- JOINクエリで一度に取得

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
