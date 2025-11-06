import { type RegisterRequestBody, type RegisterResponse } from '@/app/_features/register'
import { ResponseError } from '@/app/_types/response-error'

/**
 * ユーザー登録API
 * @param body - ユーザー登録情報
 * @returns ユーザー登録レスポンス
 * @throws {ResponseError} 登録に失敗した場合
 */
export const postRegister = async (body: RegisterRequestBody): Promise<RegisterResponse> => {
  const apiUrl = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'

  try {
    const response = await fetch(`${apiUrl}/register`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        nickname: body.nickname,
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
