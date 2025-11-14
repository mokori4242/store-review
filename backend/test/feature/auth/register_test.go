package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"store-review/test"
	"testing"
)

func TestRegisterUser_Success(t *testing.T) {
	test.TPostgres(t)
	router := test.SetupRouter(t)

	jsonBody, _ := json.Marshal(map[string]any{
		"nickname": "Test User",
		"email":    "test@example.com",
		"password": "password123",
	})
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	goldenFilePath := "../../testdata/auth/register_test.golden"

	test.AssertResponse(t, res, http.StatusCreated, goldenFilePath)
}

func TestRegister_Validation(t *testing.T) {
	test.TPostgres(t)
	router := test.SetupRouter(t)

	tests := []struct {
		name           string
		body           map[string]any
		expectedStatus int
		description    string
	}{
		{
			name: "必須フィールド不足（メールなし）",
			body: map[string]any{
				"nickname": "Test User 3",
				"password": "password123",
			},
			expectedStatus: http.StatusBadRequest,
			description:    "ニックネーム必須",
		},
		{
			name: "必須フィールド不足（パスワードなし）",
			body: map[string]any{
				"nickname": "Test User 4",
				"email":    "test4@example.com",
			},
			expectedStatus: http.StatusBadRequest,
			description:    "パスワード必須",
		},
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

			res := w.Result()
			defer res.Body.Close()

			test.AssertResponseHeader(t, res, tt.expectedStatus)
		})
	}
}
