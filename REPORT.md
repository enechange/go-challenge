## オニオンアーキテクチャ概要

オニオンアーキテクチャを採用しております。採用理由は下記に基づいています。

- **層の明確な分離**: 層を責務ごとに明確に分離し、疎結合を実現すること。これにより、内部から外部への依存を防ぎ、各層が独立して機能することが可能となり、保守性と拡張性を担保します。
- **保守性と拡張性の向上**: 分離された層により、将来的な機能拡張や変更が容易になります。層ごとに責任範囲が限定されるため、変更による影響が局所化され、保守性が向上すると考えています。
- **テストの容易さ**: 各層が独立しているため、単体テストや統合テストが行いやすくなります。

### 設計のポイント

- **層の分け方**: プレゼンテーション層、インフラストラクチャ層、アプリケーション層、ドメイン層に分け、それぞれの役割と責務を明確にしました。
- **実装のアプローチ**: 拡張性と保守性を重視し、依存関係の逆転やDIを採用しました。

### ソフトウェアアーキテクチャに関しての理解

特定のアーキテクチャパターンに名前を与えることが重要なのではなく、その背後にある設計の概念を理解し、適切に適用することが重要だと考えています。アーキテクチャの名前に囚われることなく、層を責務ごとに明確に分離し、疎結合を実現することが一番の大事なポイントという理解をしております。

### 課題と認識

今回のケースでは、オニオンアーキテクチャを採用することはやや過剰かもしれませんが、より実務的にということを見越して、より堅牢な設計をと思い採用させていただきました。


## ディレクトリ構造

```plaintext
/project-root
├── Dockerfile
├── Go-Challenge Interface Specification Document.pdf
├── README.md
├── configs
│   └── config.go   
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── application
│   │   ├── dto
│   │   │   └── active_evse_location_dto.go  # ドメイン層のモデルとアプリケーション層のモデル間のデータ変換を担うDTO
│   │   ├── query
│   │   │   ├── active_evse_location_query.go    # アクティブなEVSEのロケーションを取得するクエリサービスのインターフェース
│   │   │   ├── active_evse_location_query_test.go   # クエリ機能のテストコード
│   │   │   └── mock
│   │   │       └── mock_query.go    # クエリサービスのモック実装
│   │   └── usecase
│   │       ├── active_evse_location_usecase.go    # アクティブなEVSEのロケーションを検索するユースケース
│   │       ├── active_evse_location_usecase_test.go   # ユースケースのテストコード
│   │       └── mock
│   │           └── mock_usecase.go    # ユースケースのモック実装
│   ├── domain
│   │   ├── evse.go   # EVSEのドメインエンティティ
│   │   ├── geo_location.go   # 地理的位置を表すエンティティ
│   │   └── location.go   # ロケーションのドメインエンティティ
│   ├── infrastructure
│   │   ├── database
│   │   │   └── database.go   # データベース設定と接続管理
│   │   ├── dto
│   │   │   ├── evse_dto.go   # EVSEデータのDTO定義
│   │   │   └── location_dto.go   # ロケーションデータのDTO定義
│   │   └── query
│   │       └── active_evse_location_query_gorm.go   # GORMを使用したクエリサービスの実装
│   └── presentation
│       ├── controllers
│       │   ├── active_evse_location_controller.go   # アクティブなEVSEのロケーションを検索するコントローラ
│       │   └── active_evse_location_controller_test.go   # コントローラのテストコード
│       ├── router
│       │   └── router.go   # アプリケーションのルーティング設定
│       └── validator
│           ├── location_request_validator.go   # リクエストのバリデーションルール定義
│           └── location_request_validator_test.go   # バリデータのテストコード
├── main.go    # アプリケーションのエントリーポイント
├── sample
│   ├── evses.csv    
│   └── locations.csv   
├── server.go   
└── tmp
    └── air.log   
```

## APIドキュメンテーション

### エンドポイント: `/api/locations`

- **メソッド**: `GET`
- **パラメータ**:

  | パラメータ  | 必須/任意 | 説明         |
    |-------------|-----------|--------------|
  | `latitude`  | 必須      | 緯度         |
  | `longitude` | 必須      | 経度         |
  | `radius`    | 任意      | 検索半径(km) |

- **レスポンス**:

  | ステータスコード | 説明                               |
    |-----------------|------------------------------------|
  | 200 OK          | ロケーションのリストを返します。    |
  | 400 Bad Request | 入力パラメータが無効の場合。       |
  | 500 Internal Server Error | サーバーエラーが発生した場合。 |

## クラスとメソッドの説明

### アプリケーション層

#### クエリサービス (query)

- **ActiveEVSELocationQueryService**
    - `FindLocationsWithActiveEVSE(ctx context.Context, latitude, longitude float64, radius int) -> ([]domain.AvailableEVSELocation, error)`: 与えられた緯度、経度、半径を基にアクティブなEVSEを持つロケーションを検索し、結果を返します。

#### ユースケース (usecase)

- **ActiveEVSELocationUseCase**
    - `FindLocationsWithActiveEVSE(ctx context.Context, latitude, longitude float64, radius int) -> ([]domain.Location, error)`: リポジトリを通じてクエリを実行し、結果をドメインモデルのリストに変換して返します。

### インフラストラクチャ層

#### DTO (gormモデル)

- **EVSE**
    - フィールド: `LocationID`, `UID`, `Status`.

- **Location**
    - フィールド: `ID`, `Name`, `Address`, `Latitude`, `Longitude`.

#### クエリサービス 実装 (query)

- **ActiveEVSELocationQueryServiceGorm**
    - `FindLocationsWithActiveEVSE(ctx context.Context, latitude, longitude float64, radius int) -> ([]domain.AvailableEVSELocation, error)`: 指定されたパラメータに基づきDBからアクティブなEVSEを持つロケーションを取得し、結果を返します。

### プレゼンテーション層

#### コントローラー (controllers)

- **ActiveEVSELocationController**
    - `FetchActiveEVSELocations(c *gin.Context)`: リクエストからパラメータを抽出し、ユースケースを呼び出して結果を返します。入力値が不正な場合や処理中にエラーが発生した場合は、適切なHTTPステータスコードとエラーメッセージを返します。

#### バリデーター (validators)

- **LatitudeValidation** と **LongitudeValidation**
    - 正規表現を用いて、入力された値が適切な形式であるかを確認します。

