package store_test

import (
	"net/http"
	"net/http/httptest"
	"store-review/test"
	"testing"
)

func TestGetListStores(t *testing.T) {
	test.TPostgres(t)
	err := test.Seeding(t, "../../../internal/infrastructure/seed/store_seed.sql")
	err = test.Seeding(t, "../../../internal/infrastructure/seed/user_seed.sql")
	if err != nil {
		t.Fatalf("Failed to seed: %v", err)
	}
	router := test.SetupRouter(t)

	token := test.GetAccessToken(t, router)

	// ストア一覧を取得
	req, _ := http.NewRequest("GET", "/stores", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	goldenFilePath := "../../testdata/store/list_test.golden"

	test.AssertResponse(t, res, http.StatusOK, goldenFilePath)
}
