package test

import (
	"database/sql"
	"go-gin/internal/handler/register"
	"go-gin/internal/infrastructure/gen"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var testDB *sql.DB
var testQueries *db.Queries

// InitTestDB テスト用DBを初期化
func InitTestDB() error {
	var err error
	testDB, err = sql.Open("postgres", os.Getenv("DATABASE_CONNECTION"))
	if err != nil {
		return err
	}
	return nil
}

// SetupTestDatabase テスト用テーブルをセットアップ
func SetupTestDatabase() error {
	// テーブルが存在する場合はドロップ
	if _, err := testDB.Exec("DROP TABLE IF EXISTS users"); err != nil {
		return err
	}

	// マイグレーションファイルのパスを取得
	migrationPath := filepath.Join("..", "..", "internal", "infrastructure", "migration", "202511041100_create_users_table.up.sql")
	sqlBytes, err := os.ReadFile(migrationPath)
	if err != nil {
		return err
	}

	// マイグレーションを実行
	if _, err := testDB.Exec(string(sqlBytes)); err != nil {
		return err
	}

	// クエリオブジェクトを作成
	testQueries = db.New(testDB)

	return nil
}

// CleanupTestDatabase テスト用テーブルをクリーンアップ
func CleanupTestDatabase() {
	if testDB != nil {
		if _, err := testDB.Exec("DROP TABLE IF EXISTS users"); err != nil {
			log.Printf("failed to cleanup table: %v", err)
		}
	}
}

// CloseTestDB テスト用DBを閉じる
func CloseTestDB() {
	if testDB != nil {
		if err := testDB.Close(); err != nil {
			log.Printf("failed to close testDB: %v", err)
		}
	}
}

// TruncateUsers usersテーブルをトランケート
func TruncateUsers() error {
	if testDB == nil {
		return nil
	}
	_, err := testDB.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")
	return err
}

// GetTestDB テスト用DBを取得
func GetTestDB() *sql.DB {
	return testDB
}

// GetTestQueries テスト用クエリオブジェクトを取得
func GetTestQueries() *db.Queries {
	return testQueries
}

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	// registerハンドラーを作成
	re := register.NewHandler(testQueries)

	r := gin.Default()
	r.POST("/register", re.RegisterUser)

	return r
}
