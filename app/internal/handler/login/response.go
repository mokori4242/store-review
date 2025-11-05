package login

type Response struct {
	AccessToken string `json:"accesstoken"`
}

// 後方互換性のため内部エイリアスを維持
type loginResponse = Response
