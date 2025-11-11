package test

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"store-review/internal/handler/login"
	"store-review/internal/handler/middleware"
	"store-review/internal/handler/register"
	"store-review/internal/handler/store"
	"store-review/internal/infrastructure/gen"
	"store-review/internal/infrastructure/postgres/repository"
	"store-review/internal/usecase/auth"
	suc "store-review/internal/usecase/store"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

var testDB *sql.DB
var testQueries *db.Queries
var update = flag.Bool("update", false, "update golden files")

func TPostgres(t *testing.T) {
	t.Helper()

	ctx := context.Background()

	ctr, err := postgres.Run(ctx,
		"postgres:18-alpine",
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("tester"),
		postgres.WithPassword("password"),
		postgres.BasicWaitStrategies(),
	)
	testcontainers.CleanupContainer(t, ctr)
	if err != nil {
		t.Fatalf("Failed to start postgres container: %v", err)
	}

	// 接続文字列を取得
	connStr, err := ctr.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to get connection string: %v", err)
	}

	err = migrateDB(connStr)
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	testDB, err = sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// テスト時のみusersの作成・更新日の時間を固定
	// 増えてきたら関数に切り出すのが良さそう
	_, err = testDB.Exec("ALTER TABLE users ALTER COLUMN created_at SET DEFAULT '2025-11-11 12:00:00'::timestamp, ALTER COLUMN updated_at SET DEFAULT '2025-11-11 12:00:00'::timestamp;")
	if err != nil {
		t.Fatalf("Failed to lock time: %v", err)
	}

	// クエリオブジェクトを作成
	testQueries = db.New(testDB)
}

func Seeding(t *testing.T, seedFile string) error {
	t.Helper()

	data, err := os.ReadFile(seedFile)
	if err != nil {
		return fmt.Errorf("failed to get seed file: %s", err)
	}
	_, err = testDB.Exec(string(data))
	if err != nil {
		return fmt.Errorf("failed to seeding: %s", err)
	}

	return nil
}

func SetupRouter(t *testing.T) *gin.Engine {
	t.Helper()

	gin.SetMode(gin.TestMode)

	tJwtSecret := []byte("test")

	// リポジトリを作成
	userR := repository.NewUserRepository(testQueries)
	storeR := repository.NewStoreRepository(testQueries)

	// ユースケースを作成
	registerUC := auth.NewRegisterUseCase(userR)
	loginUC := auth.NewLoginUseCase(userR, tJwtSecret)
	sListUC := suc.NewListUseCase(storeR)

	// ハンドラーを作成
	registerH := register.NewHandler(registerUC)
	loginH := login.NewHandler(loginUC)
	storeH := store.NewHandler(sListUC)

	r := gin.Default()
	r.POST("/register", registerH.RegisterUser)
	r.POST("/login", loginH.Login)
	r.GET("/stores", middleware.JwtMiddleware(tJwtSecret), storeH.GetList)

	return r
}

func AssertResponse(t *testing.T, res *http.Response, code int, path string) {
	t.Helper()

	AssertResponseHeader(t, res, code)
	AssertResponseBodyWithFile(t, res, path)
}

func AssertResponseHeader(t *testing.T, res *http.Response, code int) {
	t.Helper()

	// ステータスコードのチェック
	if code != res.StatusCode {
		t.Errorf("expected status code is '%d',\n but actual given code is '%d'", code, res.StatusCode)
	}
	// Content-Typeのチェック
	if expected := "application/json; charset=utf-8"; res.Header.Get("Content-Type") != expected {
		t.Errorf("unexpected response Content-Type,\n expected: %#v,\n but given #%v", expected, res.Header.Get("Content-Type"))
	}
}

func AssertResponseBodyWithFile(t *testing.T, res *http.Response, path string) {
	t.Helper()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("unexpected error by io.ReadAll() '%#v'", err)
	}

	var actual bytes.Buffer
	err = json.Indent(&actual, body, "", "  ")
	if err != nil {
		t.Fatalf("unexpected error by json.Indent '%#v'", err)
	}

	// updateフラグが立っている場合はgoldenファイルを更新
	if *update {
		if err := UpdateGoldenFile(t, path, actual.Bytes()); err != nil {
			t.Fatalf("failed to update golden file: %v", err)
		}
		t.Logf("updated golden file: %s", path)
		return
	}

	rs := GetStringFromTestFile(t, path)
	assert.JSONEq(t, rs, actual.String())
}

// UpdateGoldenFile はgoldenファイルを新規作成または更新します
func UpdateGoldenFile(t *testing.T, path string, body []byte) error {
	t.Helper()

	// ディレクトリが存在しない場合は作成
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// goldenファイルに書き込み
	if err := os.WriteFile(path, body, 0644); err != nil {
		return fmt.Errorf("failed to write golden file: %w", err)
	}

	return nil
}

func GetStringFromTestFile(t *testing.T, path string) string {
	t.Helper()

	bt, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("unexpected error while opening file '%#v'", err)
	}
	return string(bt)
}

func GetCookie(t *testing.T, r *gin.Engine) *http.Cookie {
	t.Helper()

	// ログインしてトークンを取得
	loginBody := login.Request{
		Email:    "test1@example.com",
		Password: "password",
	}
	loginJSON, _ := json.Marshal(loginBody)
	loginReq, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(loginJSON))
	loginReq.Header.Set("Content-Type", "application/json")
	loginW := httptest.NewRecorder()
	r.ServeHTTP(loginW, loginReq)

	return loginW.Result().Cookies()[0]
}

func migrateDB(connStr string) error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get current file")
	}

	migrationPath := filepath.Join(filepath.Dir(filename), "../internal/infrastructure/migration")
	absMigrationPath, err := filepath.Abs(migrationPath)
	if err != nil {
		return fmt.Errorf("failed to get absolute migration path: %s", err)
	}

	m, err := migrate.New("file://"+absMigrationPath, connStr)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %s", err)
	}

	if err := m.Up(); err != nil && errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to apply migration: %s", err)
	}

	return nil
}
