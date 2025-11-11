import { type StoreListResponse } from '@/app/_features/store-list/model/types'
import { ResponseError } from '@/app/_types/response-error'

/**
 * 店舗一覧取得API
 * @returns 店舗一覧レスポンス
 * @throws {ResponseError} 店舗一覧の取得に失敗した場合
 */
export const getStoreList = async (): Promise<StoreListResponse> => {
  const apiUrl = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'

  try {
    const response = await fetch(`${apiUrl}/stores`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      },
      credentials: 'include' // HTTPOnly Cookieを送信するために必要
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
