package feature_test

import (
	"log"
	"os"
	"store-review/test"
	"testing"
)

func TestMain(m *testing.M) {
	// テスト用DBを初期化
	if err := test.InitTestDB(); err != nil {
		log.Fatalf("failed to init test DB: %v", err)
	}
	defer test.CloseTestDB()

	// テスト用テーブルをセットアップ
	if err := test.SetupTestDatabase(); err != nil {
		log.Fatalf("failed to setup test database: %v", err)
	}

	// テスト実行
	code := m.Run()

	// テスト終了後にテーブルをクリーンアップ
	test.CleanupTestDatabase()

	os.Exit(code)
}
