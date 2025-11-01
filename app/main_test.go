package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	db "go-gin/gen"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var testDB *sql.DB
var testQueries *db.Queries

func TestMain(m *testing.M) {
	// テスト用のインメモリSQLiteデータベースを作成
	var err error
	testDB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer testDB.Close()

	// テーブル作成（実際のスキーマに合わせて調整が必要）
	createTableSQL := `
	CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		phone_number TEXT,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = testDB.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}

	testQueries = db.New(testDB)

	// テスト実行
	code := m.Run()

	os.Exit(code)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	// グローバル変数qをテスト用のqueriesに設定
	q = testQueries

	// バリデーターをセットアップ
	setupValidator()

	r := gin.New()
	r.POST("/users", createUser)
	r.GET("/users/:id", getUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	return r
}

func TestCreateUser(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name           string
		body           map[string]any
		expectedStatus int
	}{
		{
			name: "正常なユーザー作成",
			body: map[string]any{
				"name":         "Test User",
				"email":        "test@example.com",
				"phone_number": "09012345678",
				"password":     "password123",
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "電話番号なしのユーザー作成",
			body: map[string]any{
				"name":     "Test User 2",
				"email":    "test2@example.com",
				"password": "password123",
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "必須フィールド不足（メールなし）",
			body: map[string]any{
				"name":     "Test User 3",
				"password": "password123",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "必須フィールド不足（パスワードなし）",
			body: map[string]any{
				"name":  "Test User 4",
				"email": "test4@example.com",
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedStatus == http.StatusCreated {
				var response response
				err := json.Unmarshal(w.Body.Bytes(), &response)
				if err != nil {
					t.Fatalf("failed to unmarshal response: %v", err)
				}

				if response.ID == 0 {
					t.Error("expected valid user ID, got 0")
				}
				if response.Name != tt.body["name"].(string) {
					t.Errorf("expected name %s, got %s", tt.body["name"].(string), response.Name)
				}
				if response.Email != tt.body["email"].(string) {
					t.Errorf("expected email %s, got %s", tt.body["email"].(string), response.Email)
				}
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	router := setupRouter()

	// テスト用ユーザーを作成
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testpass"), bcrypt.DefaultCost)
	params := db.CreateUserParams{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: string(hashedPassword),
	}
	user, err := testQueries.CreateUser(context.Background(), params)
	if err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	tests := []struct {
		name           string
		userID         string
		expectedStatus int
	}{
		{
			name:           "存在するユーザーの取得",
			userID:         fmt.Sprintf("%d", user.ID),
			expectedStatus: http.StatusOK,
		},
		{
			name:           "存在しないユーザーの取得",
			userID:         "99999",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "無効なユーザーID",
			userID:         "invalid",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/users/"+tt.userID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedStatus == http.StatusOK {
				var response response
				err := json.Unmarshal(w.Body.Bytes(), &response)
				if err != nil {
					t.Fatalf("failed to unmarshal response: %v", err)
				}

				if response.ID != user.ID {
					t.Errorf("expected user ID %d, got %d", user.ID, response.ID)
				}
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	router := setupRouter()

	// テスト用ユーザーを作成
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testpass"), bcrypt.DefaultCost)
	params := db.CreateUserParams{
		Name:     "Original User",
		Email:    "original@example.com",
		Password: string(hashedPassword),
	}
	user, err := testQueries.CreateUser(context.Background(), params)
	if err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	tests := []struct {
		name           string
		userID         string
		body           map[string]any
		expectedStatus int
	}{
		{
			name:   "正常なユーザー更新",
			userID: fmt.Sprintf("%d", user.ID),
			body: map[string]any{
				"name":         "Updated User",
				"email":        "updated@example.com",
				"phone_number": "09011112222",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:   "電話番号なしでの更新",
			userID: fmt.Sprintf("%d", user.ID),
			body: map[string]any{
				"name":  "Updated User 2",
				"email": "updated2@example.com",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:   "存在しないユーザーの更新",
			userID: "99999",
			body: map[string]any{
				"name":  "Non-existent User",
				"email": "nonexistent@example.com",
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:   "無効なユーザーID",
			userID: "invalid",
			body: map[string]any{
				"name":  "Invalid User",
				"email": "invalid@example.com",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:   "名前のみ更新",
			userID: fmt.Sprintf("%d", user.ID),
			body: map[string]any{
				"name": "Only Name Updated",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:   "メールのみ更新",
			userID: fmt.Sprintf("%d", user.ID),
			body: map[string]any{
				"email": "onlyemail@example.com",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:   "電話番号のみ更新",
			userID: fmt.Sprintf("%d", user.ID),
			body: map[string]any{
				"phone_number": "09088887777",
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("PUT", "/users/"+tt.userID, bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedStatus == http.StatusOK {
				var response response
				err := json.Unmarshal(w.Body.Bytes(), &response)
				if err != nil {
					t.Fatalf("failed to unmarshal response: %v", err)
				}

				if name, ok := tt.body["name"].(string); ok {
					if response.Name != name {
						t.Errorf("expected name %s, got %s", name, response.Name)
					}
				}
				if email, ok := tt.body["email"].(string); ok {
					if response.Email != email {
						t.Errorf("expected email %s, got %s", email, response.Email)
					}
				}
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	router := setupRouter()

	// テスト用ユーザーを作成
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testpass"), bcrypt.DefaultCost)
	params := db.CreateUserParams{
		Name:     "User to Delete",
		Email:    "delete@example.com",
		Password: string(hashedPassword),
	}
	user, err := testQueries.CreateUser(context.Background(), params)
	if err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	tests := []struct {
		name           string
		userID         string
		expectedStatus int
	}{
		{
			name:           "存在するユーザーの削除",
			userID:         fmt.Sprintf("%d", user.ID),
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "存在しないユーザーの削除",
			userID:         "99999",
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "無効なユーザーID",
			userID:         "invalid",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("DELETE", "/users/"+tt.userID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestValidation(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name           string
		body           map[string]any
		expectedStatus int
		description    string
	}{
		{
			name: "有効な11桁の電話番号",
			body: map[string]any{
				"name":         "Valid Phone User",
				"email":        "validphone@example.com",
				"phone_number": "09012345678",
				"password":     "password123",
			},
			expectedStatus: http.StatusCreated,
			description:    "11桁の電話番号は受け入れられる",
		},
		{
			name: "無効な電話番号（10桁）",
			body: map[string]any{
				"name":         "Invalid Phone User",
				"email":        "invalidphone@example.com",
				"phone_number": "0901234567",
				"password":     "password123",
			},
			expectedStatus: http.StatusBadRequest,
			description:    "10桁の電話番号は拒否される",
		},
		{
			name: "無効な電話番号（12桁）",
			body: map[string]any{
				"name":         "Invalid Phone User",
				"email":        "invalidphone2@example.com",
				"phone_number": "090123456789",
				"password":     "password123",
			},
			expectedStatus: http.StatusBadRequest,
			description:    "12桁の電話番号は拒否される",
		},
		{
			name: "無効な電話番号（文字が含まれる）",
			body: map[string]any{
				"name":         "Invalid Phone User",
				"email":        "invalidphone3@example.com",
				"phone_number": "090-1234-5678",
				"password":     "password123",
			},
			expectedStatus: http.StatusBadRequest,
			description:    "ハイフンが含まれる電話番号は拒否される",
		},
		{
			name: "無効なメールアドレス",
			body: map[string]any{
				"name":     "Invalid Email User",
				"email":    "invalid-email",
				"password": "password123",
			},
			expectedStatus: http.StatusBadRequest,
			description:    "無効なメールアドレスは拒否される",
		},
		{
			name: "短すぎるパスワード",
			body: map[string]any{
				"name":     "Short Password User",
				"email":    "shortpass@example.com",
				"password": "short",
			},
			expectedStatus: http.StatusBadRequest,
			description:    "8文字未満のパスワードは拒否される",
		},
		{
			name: "長すぎる名前",
			body: map[string]any{
				"name":     "Very Long Name That Exceeds The Maximum Length Of Forty Characters",
				"email":    "longname@example.com",
				"password": "password123",
			},
			expectedStatus: http.StatusBadRequest,
			description:    "40文字を超える名前は拒否される",
		},
		{
			name: "電話番号なし（空文字列）",
			body: map[string]any{
				"name":         "No Phone User",
				"email":        "nophone@example.com",
				"phone_number": "",
				"password":     "password123",
			},
			expectedStatus: http.StatusCreated,
			description:    "電話番号が空文字列の場合は受け入れられる",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("%s: expected status %d, got %d", tt.description, tt.expectedStatus, w.Code)
				t.Logf("Response body: %s", w.Body.String())
			}
		})
	}
}

func TestUpdateValidation(t *testing.T) {
	router := setupRouter()

	// テスト用ユーザーを作成
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testpass"), bcrypt.DefaultCost)
	params := db.CreateUserParams{
		Name:     "Original User",
		Email:    "original@example.com",
		Password: string(hashedPassword),
	}
	user, err := testQueries.CreateUser(context.Background(), params)
	if err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	tests := []struct {
		name           string
		body           map[string]any
		expectedStatus int
		description    string
	}{
		{
			name: "有効な更新（11桁電話番号）",
			body: map[string]any{
				"name":         "Updated User",
				"email":        "updated@example.com",
				"phone_number": "09087654321",
			},
			expectedStatus: http.StatusOK,
			description:    "有効なデータでの更新は成功する",
		},
		{
			name: "無効な電話番号での更新",
			body: map[string]any{
				"name":         "Updated User",
				"email":        "updated2@example.com",
				"phone_number": "090-8765-4321",
			},
			expectedStatus: http.StatusBadRequest,
			description:    "無効な電話番号での更新は拒否される",
		},
		{
			name: "電話番号なしでの更新",
			body: map[string]any{
				"name":  "Updated User Without Phone",
				"email": "updated3@example.com",
			},
			expectedStatus: http.StatusOK,
			description:    "電話番号なしでの更新は成功する",
		},
		{
			name: "無効なメールでの更新",
			body: map[string]any{
				"name":  "Updated User",
				"email": "invalid-email-format",
			},
			expectedStatus: http.StatusBadRequest,
			description:    "無効なメールアドレスでの更新は拒否される",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("PUT", fmt.Sprintf("/users/%d", user.ID), bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("%s: expected status %d, got %d", tt.description, tt.expectedStatus, w.Code)
				t.Logf("Response body: %s", w.Body.String())
			}
		})
	}
}

// ベンチマークテストの例
func BenchmarkCreateUser(b *testing.B) {
	router := setupRouter()

	body := map[string]any{
		"name":     "Benchmark User",
		"email":    "benchmark@example.com",
		"password": "password123",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		body["email"] = fmt.Sprintf("benchmark%d@example.com", i) // 重複を避けるため
		jsonBody, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}
}
