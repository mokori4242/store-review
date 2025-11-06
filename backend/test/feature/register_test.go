package feature_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"store-review/internal/handler/register"
	"store-review/test"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	router := test.SetupRouter()

	tests := []struct {
		name           string
		body           map[string]any
		expectedStatus int
	}{
		{
			name: "正常なユーザー作成",
			body: map[string]any{
				"nickname": "Test User",
				"email":    "test@example.com",
				"password": "password123",
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "必須フィールド不足（メールなし）",
			body: map[string]any{
				"nickname": "Test User 3",
				"password": "password123",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "必須フィールド不足（パスワードなし）",
			body: map[string]any{
				"nickname": "Test User 4",
				"email":    "test4@example.com",
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

			jsonBody, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
				t.Logf("Response body: %s", w.Body.String())
			}

			if tt.expectedStatus == http.StatusCreated {
				var response register.Response
				err := json.Unmarshal(w.Body.Bytes(), &response)
				if err != nil {
					t.Fatalf("failed to unmarshal response: %v", err)
				}

				if response.ID == 0 {
					t.Error("expected valid user ID, got 0")
				}
				if response.Nickname != tt.body["nickname"].(string) {
					t.Errorf("expected name %s, got %s", tt.body["nickname"].(string), response.Nickname)
				}
				if response.Email != tt.body["email"].(string) {
					t.Errorf("expected email %s, got %s", tt.body["email"].(string), response.Email)
				}
			}
		})
	}
}

func TestRegister_Validation(t *testing.T) {
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
				"nickname": "Invalid Email User",
				"email":    "invalid-email",
				"password": "password123",
			},
			expectedStatus: http.StatusBadRequest,
			description:    "無効なメールアドレスは拒否される",
		},
		{
			name: "短すぎるパスワード",
			body: map[string]any{
				"nickname": "Short Password User",
				"email":    "shortpass@example.com",
				"password": "short",
			},
			expectedStatus: http.StatusBadRequest,
			description:    "8文字未満のパスワードは拒否される",
		},
		{
			name: "長すぎる名前",
			body: map[string]any{
				"nickname": "Very Long Name That Exceeds The Maximum Length Of Forty Characters",
				"email":    "longname@example.com",
				"password": "password123",
			},
			expectedStatus: http.StatusBadRequest,
			description:    "40文字を超える名前は拒否される",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
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
