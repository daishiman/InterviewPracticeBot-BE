# Interview Practice Bot Backend

本プロジェクトは、面接対策のためのボットのバックエンドシステムを構築するものです。

## プロジェクト構成

- [プロジェクトディレクトリ構成](#プロジェクトディレクトリ構成)
- [環境設定](#環境設定)
  - [Goのインストール](#goのインストール)
  - [依存関係のインストール](#依存関係のインストール)
  - [PlanetScaleのセットアップ](#planetscaleのセットアップ)
- [Docker設定](#docker設定)
- [データベースの設定](#データベースの設定)
- [品質改善ツール](#品質改善ツール)
- [GitHub Actionsの設定](#github-actionsの設定)

### プロジェクトディレクトリ構成

プロジェクトのディレクトリ構造とそれぞれのディレクトリでの主なソースコードの内容は以下の通りです。

```
project-root/
│
├── .github/                  # GitHub Actions の設定ファイル
│   ├── workflows/            # CI/CD パイプラインの定義
│       ├── test_workflow.yml # テストワークフローの定義
│
│
├── cmd/                      # メインアプリケーション
│   ├── server/               # サーバーのエントリーポイント
│
├── internal/                 # 内部パッケージ
│   ├── config/               # アプリケーションの設定や環境変数を管理
│   ├── domain/               # ドメインロジック
│   ├── infrastructure/       # インフラストラクチャ
│   ├── interfaces/           # インタフェース
│   ├── middleware/           # ミドルウェアのロジック (例: 認証ミドルウェア、ロギングミドルウェアなど)
│   └── usecase/              # ユースケース
│
├── migrations/               # データベースマイグレーションファイル
│
├── pkg/                      # 外部で利用可能なパッケージ
│
├── Dockerfile                # Docker のビルドファイル
├── docker-compose.yml        # Docker Compose の設定ファイル
├── go.mod                    # Go のモジュール依存関係
└── go.sum                    # Go のモジュール依存関係のチェックサム
```

### 環境設定

#### Goのインストール

Goは、[公式サイト](https://golang.org/dl/)からダウンロードできます。macOSでのインストール例:

```
arch -arm64 brew install go
```

#### 依存関係のインストール

プロジェクトルートで以下のコマンドを実行して、Go モジュールを初期化し、依存関係をインストールします。

```
go mod init InterviewPracticeBot-BE
go get -u github.com/gorilla/mux \
        github.com/joho/godotenv \
        github.com/dgrijalva/jwt-go \
        github.com/labstack/echo/v4 \
        github.com/rs/cors \
        gorm.io/gorm \
        gorm.io/plugin/dbresolver \
        github.com/smartystreets/goconvey \
        github.com/go-sql-driver/mysql
```

#### PlanetScaleのセットアップ

1. [PlanetScale](https://planetscale.com/)にアクセスしてアカウントを作成します。
2. PlanetScale CLIをインストールします。

```
arch -arm64 brew install planetscale/tap/pscale
```

3. CLIでログインし、新しいデータベースとデータベースのブランチを作成します。

```
pscale auth login
pscale database create <database-name>
pscale branch create <database-name> main
```

4. PlanetScaleデータベースへの接続情報を取得します。

```
pscale connect interviewpracticebot main
```

### Docker設定

プロジェクトルートに`Dockerfile`と`docker-compose.yml`を作成し、それぞれのファイルに対して記述します。

### データベースの設定

PlanetScaleのマイグレーションツールとプロセスを使用してマイグレーションを管理します。

```bash
pscale branch create <database-name> <branch-name>
pscale migrate apply <database-name> <branch-name>
pscale branch merge <database-name> <branch-name>
```

### 品質改善ツール

#### 1. golangci-lint (lintツール):

- golangci-lintは、コードを解析し、プログラムのエラーやバグ、スタイルの問題、または潜在的な問題を特定するためのツールです。

##### 導入方法:

```bash
go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.0
```

##### 使い方:

- プロジェクトディレクトリで以下のコマンドを実行してコードを解析します。

```bash
golangci-lint run
```

#### 2. gofmtまたはgoimports (フォーマットツール):

- gofmtは、コードの構文を整理し、一貫したスタイルとフォーマットでコードを自動的に再フォーマットします。

##### 使い方 (gofmtの場合):

- プロジェクトディレクトリで以下のコマンドを実行してコードを整形します。

```bash
gofmt -w .
```

### GitHub Actionsの設定

`.github/workflows/test_workflow.yml`にCI/CDの設定を記述します。
