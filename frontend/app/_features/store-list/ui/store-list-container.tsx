'use client'

import { useStoreList } from '../api/use-store-list'
import { StoreListPresentational } from './store-list-presentational'

export const StoreListContainer = () => {
  const { data, isLoading, error } = useStoreList()

  return (
    <StoreListPresentational
      stores={data?.stores || []}
      isLoading={isLoading}
      error={error?.message}
    />
  )
}
