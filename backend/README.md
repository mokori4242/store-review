# Go-Gin アプリケーション技術スタック概要

## プロジェクト概要
ユーザー管理機能を持つREST APIアプリケーション

## 技術スタック

### バックエンドフレームワーク
- **Go** (v1.24.6)
- **Gin** (v1.10.1) - 高性能なHTTPウェブフレームワーク

### データベース
- **SQLite3** - 軽量なリレーショナルデータベース
- **SQLC** - SQLファーストなGoコード生成ツール

### 主要ライブラリ
- `github.com/mattn/go-sqlite3` - SQLiteドライバー
- `golang.org/x/crypto/bcrypt` - パスワードハッシュ化
- `github.com/go-playground/validator/v10` - リクエスト検証

### プロジェクト構造
```
go-gin/
├── main.go              # メインアプリケーションファイル
├── main_test.go         # テストファイル
├── go.mod              # Go modules設定
├── sqlc.yaml           # SQLC設定
├── Makefile            # データベース管理用コマンド
├── sqlite.db           # SQLiteデータベースファイル
├── migration/          # データベーススキーママイグレーション
│   ├── 202511041100_create_users_table.up.sql
│   └── 202509181735_create_user_cars_table.sql
├── query/              # SQLクエリファイル
│   ├── users.sql
│   └── user_cars.sql
└── gen/                # SQLC生成されたGoコード
    ├── db.go
    ├── models.go
    ├── users.sql.go
    └── user_cars.sql.go
```

## 機能概要

### ユーザー管理API
- **POST /users** - ユーザー作成
- **GET /users/:id** - ユーザー取得
- **PUT /users/:id** - ユーザー更新
- **DELETE /users/:id** - ユーザー削除

### データモデル
#### Usersテーブル
- id (INTEGER, PRIMARY KEY, AUTOINCREMENT)
- name (TEXT, NOT NULL)
- email (TEXT, UNIQUE, NOT NULL)
- phone_number (TEXT, NULLABLE)
- password (TEXT, NOT NULL) - bcryptでハッシュ化
- created_at (DATETIME)
- updated_at (DATETIME)

### 主要機能
1. **バリデーション機能**
   - 名前: 最大40文字
   - メール: 有効なメール形式
   - 電話番号: 11桁の数字（カスタムバリデーター）
   - パスワード: 最小8文字

2. **セキュリティ機能**
   - パスワードのbcryptハッシュ化
   - リクエストバリデーション

3. **データベース操作**
   - SQLC による型安全なクエリ生成
   - マイグレーション管理

### テスト環境
- インメモリSQLiteを使用した単体テスト
- HTTPテスト用のモック環境

### 開発・運用ツール
- **Makefile** - データベース管理コマンド
  - `make init-db` - データベース初期化
  - `make create-tables` - テーブル作成
  - `make refresh` - 完全リフレッシュ
  - `make reset-db` - データベースリセット

## アーキテクチャの特徴
1. **SQLファースト設計** - SQLCを使用してSQLからGoコードを自動生成
2. **レイヤー分離** - データアクセス層とAPIハンドラー層の分離
3. **型安全** - コンパイル時の型チェック
4. **軽量** - 最小限の依存関係
5. **テスタブル** - インメモリデータベースによるテスト環境