import { useQuery, type UseQueryResult } from '@tanstack/react-query'

import { getStoreList } from '@/app/_api/store-list/get-store-list'
import { type ResponseError } from '@/app/_types/response-error'

import { type StoreListResponse } from '../model/types'

/**
 * 店舗一覧取得カスタムフック
 * @returns useQueryの結果
 */
export const useStoreList = (): UseQueryResult<StoreListResponse, ResponseError> => {
  return useQuery({
    queryKey: ['storeList'],
    queryFn: () => getStoreList()
  })
}
