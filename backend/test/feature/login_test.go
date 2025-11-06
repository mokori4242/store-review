package feature_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"store-review/internal/handler/login"
	"store-review/test"
	"testing"

	"github.com/gin-gonic/gin"
)

// registerTestUser テスト用のユーザーを登録するヘルパー関数
func registerTestUser(router *gin.Engine, email, password, nickname string) error {
	body := map[string]any{
		"email":    email,
		"password": password,
		"nickname": nickname,
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		return nil
	}
	return nil
}

func TestLogin(t *testing.T) {
	router := test.SetupRouter()

	tests := []struct {
		name           string
		setupUser      bool
		userEmail      string
		userPassword   string
		userNickname   string
		loginBody      map[string]any
		expectedStatus int
	}{
		{
			name:         "正常なログイン",
			setupUser:    true,
			userEmail:    "test@example.com",
			userPassword: "password123",
			userNickname: "Test User",
			loginBody: map[string]any{
				"email":    "test@example.com",
				"password": "password123",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:      "存在しないメールアドレス",
			setupUser: false,
			loginBody: map[string]any{
				"email":    "notexist@example.com",
				"password": "password123",
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:         "間違ったパスワード",
			setupUser:    true,
			userEmail:    "test2@example.com",
			userPassword: "correctpassword",
			userNickname: "Test User 2",
			loginBody: map[string]any{
				"email":    "test2@example.com",
				"password": "wrongpassword",
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:      "必須フィールド不足（メールなし）",
			setupUser: false,
			loginBody: map[string]any{
				"password": "password123",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:      "必須フィールド不足（パスワードなし）",
			setupUser: false,
			loginBody: map[string]any{
				"email": "test@example.com",
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// テストケース実行前にusersテーブルをクリーンアップ
			if err := test.TruncateUsers(); err != nil {
				t.Fatalf("failed to truncate table: %v", err)
			}

			// テストユーザーを事前登録
			if tt.setupUser {
				if err := registerTestUser(router, tt.userEmail, tt.userPassword, tt.userNickname); err != nil {
					t.Fatalf("failed to register test user: %v", err)
				}
			}

			// ログインリクエストを送信
			jsonBody, _ := json.Marshal(tt.loginBody)
			req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
				t.Logf("Response body: %s", w.Body.String())
			}

			// 成功時のレスポンス検証
			if tt.expectedStatus == http.StatusOK {
				var response login.Response
				err := json.Unmarshal(w.Body.Bytes(), &response)
				if err != nil {
					t.Fatalf("failed to unmarshal response: %v", err)
				}

				if response.AccessToken == "" {
					t.Error("expected valid access token, got empty string")
				}
			}
		})
	}
}

func TestLogin_Validation(t *testing.T) {
	router := test.SetupRouter()

	tests := []struct {
		name           string
		body           map[string]any
		expectedStatus int
		description    string
	}{
		{
			name: "無効なメールアドレス",
			body: map[string]any{
				"email":    "invalid-email",
				"password": "password123",
			},
			expectedStatus: http.StatusBadRequest,
			description:    "無効なメールアドレス形式は拒否される",
		},
		{
			name: "空のメールアドレス",
			body: map[string]any{
				"email":    "",
				"password": "password123",
			},
			expectedStatus: http.StatusBadRequest,
			description:    "空のメールアドレスは拒否される",
		},
		{
			name: "空のパスワード",
			body: map[string]any{
				"email":    "test@example.com",
				"password": "",
			},
			expectedStatus: http.StatusBadRequest,
			description:    "空のパスワードは拒否される",
		},
		{
			name: "存在しないユーザー",
			body: map[string]any{
				"email":    "test@example.com",
				"password": "password123",
			},
			expectedStatus: http.StatusUnauthorized,
			description:    "存在しないユーザーは拒否される",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
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
