/**
 * 店舗情報
 */
export interface Store {
  id: number
  name: string
  regularHolidays: string[]
  categoryNames: string[]
  paymentMethods: string[]
  webProfiles: string[]
}

/**
 * 店舗一覧APIのレスポンス
 */
export interface StoreListResponse {
  stores: Store[]
}
