package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"store-review/test"
	"testing"
)

func TestLogin_Success(t *testing.T) {
	test.TPostgres(t)
	err := test.Seeding(t, "../../../internal/infrastructure/seed/user_seed.sql")
	if err != nil {
		t.Fatalf("Failed to seed: %v", err)
	}
	router := test.SetupRouter(t)

	jsonBody, _ := json.Marshal(map[string]string{
		"email":    "test1@example.com",
		"password": "password",
	})
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	test.AssertResponseHeader(t, res, http.StatusOK)

	c := res.Cookies()[0]

	err = c.Valid()
	if err != nil {
		t.Errorf("Invalid cookie value: %v", c)
	}

	cs := c.String()[0:11]

	if cs != "accessToken" {
		t.Errorf("Invalid access token: %v", c)
	}
}

func TestLogin_Validation(t *testing.T) {
	test.TPostgres(t)
	router := test.SetupRouter(t)

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
			expectedStatus: http.StatusBadRequest,
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

			res := w.Result()
			defer res.Body.Close()

			test.AssertResponseHeader(t, res, tt.expectedStatus)
		})
	}
}
