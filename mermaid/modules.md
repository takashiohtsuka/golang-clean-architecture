# モジュール役割図

## レイヤー構成と各モジュールの責務

```mermaid
graph TD
    subgraph CMD["cmd/app/main.go　エントリーポイント"]
        MAIN["・アプリ起動・DB接続取得・DI組み立て・サーバー起動"]
    end

    subgraph INFRA["infrastructure　フレームワーク・外部サービス接続"]
        DS["datastore・MySQL接続・gorm.DB生成"]
        RT["router・Echoルーティング定義・エンドポイントとControllerを紐付け"]
    end

    subgraph REG["registry　依存性注入コンテナ"]
        RG["registry・Repository/Usecase/Controllerを　コンストラクタインジェクションで組み立て・AppControllerを返す"]
    end

    subgraph ADAPTER["adapter　外部とドメインの橋渡し"]
        CTRL["controller・HTTPリクエスト受付・パラメータのバインド・バリデーション・Usecaseインターフェースを定義(Pattern B)・レスポンス返却"]
        REPO["repository・outputportインターフェースの具象実装・gorm経由でDB操作・WithTx()でトランザクション伝播"]
        MAPPER["mapper・domain/model ↔ domain/entityの変換・DB層とビジネス層のデータ構造を分離"]
    end

    subgraph USECASE["usecase　ビジネスロジック"]
        INTER["interactor・ユースケースの具象実装・ビジネスルールの実行・Repositoryを組み合わせて処理"]
        OUTPORT["outputport・Repositoryのインターフェース定義・interactorが使う側として定義(Pattern B)・DBへの依存を抽象化"]
    end

    subgraph DOMAIN["domain　ドメイン知識"]
        MODEL["model・GORMモデル・DBテーブルとマッピングされる構造体・gormタグでカラム定義"]
        ENTITY["entity・ドメインオブジェクト・ビジネスロジックが扱うデータ構造・DBに依存しない"]
        STRUCT["structure/concurrent・Limiter構造体・goroutineの同時実行数を制御するチャネル管理"]
    end

    MAIN --> DS
    MAIN --> RG
    MAIN --> RT
    RG --> CTRL
    RG --> INTER
    RG --> REPO
    RT --> CTRL
    CTRL -->|インターフェース経由| INTER
    INTER -->|インターフェース経由| REPO
    REPO --> MAPPER
    MAPPER --> MODEL
    MAPPER --> ENTITY
    INTER --> OUTPORT
    INTER --> STRUCT
```

---

## 各レイヤーの役割まとめ

```mermaid
graph LR
    subgraph 外側
        direction TB
        A["infrastructure────────────フレームワーク・DB接続Echo / GORM / MySQL"]
        B["registry────────────依存性注入コンストラクタインジェクション"]
    end

    subgraph 中間
        direction TB
        C["adapter/controller────────────Input PortHTTPリクエスト受付"]
        D["adapter/repository────────────Output Port実装DB操作の具象クラス"]
        E["adapter/mapper────────────データ変換model ↔ entity"]
    end

    subgraph 内側
        direction TB
        F["usecase/interactor────────────ビジネスロジックユースケース実装"]
        G["usecase/outputport────────────Output Port定義Repositoryインターフェース"]
    end

    subgraph 核心
        direction TB
        H["domain/entity────────────ドメインオブジェクトDB非依存"]
        I["domain/model────────────GORMモデルDBマッピング"]
        J["domain/structure────────────goroutine制御Limiter"]
    end

    外側 --> 中間 --> 内側 --> 核心
```

---

## インターフェース(Pattern B)と依存の向き

```mermaid
graph LR
    subgraph adapter/controller
        UI["<<interface>>UserUsecaseList() / Create()"]
        UC["userController(具象)"]
    end

    subgraph usecase/interactor
        US["UserUsecase(具象struct)"]
    end

    subgraph usecase/outputport
        RI["<<interface>>UserRepositoryFindAll() / Create() / WithTx()"]
    end

    subgraph adapter/repository
        RS["userRepository(具象)"]
    end

    UC -->|フィールドに持つ| UI
    US -.->|暗黙的に満たす| UI
    US -->|フィールドに持つ| RI
    RS -.->|暗黙的に満たす| RI

    style UI fill:#ffd700,color:#000
    style RI fill:#ffd700,color:#000
```

> **Point**: インターフェースは「使う側」のパッケージに定義する(Pattern B)。
> 具象実装はインターフェースを知らなくてよい。依存の向きが内側に向く。

---

## domain/model と domain/entity の違い

```mermaid
graph TD
    subgraph MySQL
        TBL["usersテーブルid / name / agecreated_at / updated_at / deleted_at"]
    end

    subgraph domain/model
        MDL["User structgorm.Model埋め込みgormタグでカラム定義GORMが直接読み書き"]
    end

    subgraph adapter/mapper
        MP["mapper.ToEntity()mapper.ToOrmModel()2つの構造体を相互変換"]
    end

    subgraph domain/entity
        ENT["User structビジネスロジックが扱うDBタグなしDBに依存しない"]
    end

    subgraph usecase/interactor
        UC["UserUsecaseentityを引数・戻り値に使うDBの都合を知らない"]
    end

    TBL <-->|GORM| MDL
    MDL <-->|変換| MP
    MP <-->|変換| ENT
    ENT <-->|使用| UC
```

---

## goroutine関連モジュール(fanIn / urlDownloadConcurrent)

```mermaid
graph TD
    subgraph domain/structure/concurrent
        LIM["Limiter structConcurrentCh chan struct\{\}同時実行数の上限を管理するバッファ付きチャネル"]
    end

    subgraph usecase/interactor
        FAN["FanInUsecasegoroutineを複数起動チャネルで結果をマージ(fan-in)context.Cancelで全goroutine停止"]
        DLC["URLDownloadConcurrentUsecaseLimiterで最大20本に制限sync.WaitGroupで完了を待機並列HTTPダウンロード"]
        DLS["URLDownloadSequentialUsecase1件ずつ順番にHTTPダウンロードgoroutineなし"]
    end

    LIM -->|フィールドに持つ| FAN
    LIM -->|フィールドに持つ| DLC

    subgraph 比較
        SEQ["シーケンシャル処理時間 = Σ(各ダウンロード時間)"]
        CON["コンカレント処理時間 ≒ max(各ダウンロード時間)"]
    end

    DLS --- SEQ
    DLC --- CON
```

---

## registry によるDI組み立てフロー

```mermaid
graph LR
    subgraph registry
        RG["NewAppController()"]
        RU["NewUserController()NewStaffController()NewRoleController()NewFanInController()..."]
    end

    subgraph adapter/repository
        NREPO["NewUserRepository(db)NewDBRepository(db)..."]
    end

    subgraph usecase/interactor
        NCASE["NewUserUsecase(  userRepo,  dbRepo)"]
    end

    subgraph adapter/controller
        NCTRL["NewUserController(  usecase)"]
    end

    RG --> RU
    RU -->|db渡し| NREPO
    NREPO -->|interface返却| RU
    RU -->|interface渡し| NCASE
    NCASE -->|struct返却| RU
    RU -->|struct渡し| NCTRL
    NCTRL -->|controller返却| RU
    RU -->|AppController返却| RG
```
