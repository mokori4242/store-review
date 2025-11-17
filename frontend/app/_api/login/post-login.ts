import { type LoginRequestBody, type LoginResponse } from '@/app/_features/login'
import { ResponseError } from '@/app/_types/response-error'

/**
 * ログインAPI
 * @param body - ログイン情報
 * @returns ログインレスポンス
 * @throws {ResponseError} ログインに失敗した場合
 */
export const postLogin = async (body: LoginRequestBody): Promise<LoginResponse> => {
  const apiUrl = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'

  try {
    const response = await fetch(`${apiUrl}/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      credentials: 'include', // Cookieを送受信するために必要
      body: JSON.stringify({
        email: body.email,
        password: body.password
      })
    })

    if (!response.ok) {
      throw new ResponseError(response.statusText, response)
    }

    return await response.json()
  } catch (error) {
    if (error instanceof ResponseError) {
      throw error
    }
    throw new Error('予期せぬエラー', { cause: error })
  }
}
