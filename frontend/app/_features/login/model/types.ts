/**
 * ログインAPIのリクエストボディ
 */
export interface LoginRequestBody {
  email: string
  password: string
}

/**
 * ログインAPIのレスポンス
 */
export interface LoginResponse {
  accessToken: string
}
