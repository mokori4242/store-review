/**
 * ユーザー登録APIのリクエストボディ
 */
export interface RegisterRequestBody {
  nickname: string
  email: string
  password: string
}

/**
 * ユーザー登録APIのレスポンス
 */
export interface RegisterResponse {
  status: number
  id: number
  nickname: string
  email: string
  createdAt: string
  updatedAt: string
  detail: string
}
