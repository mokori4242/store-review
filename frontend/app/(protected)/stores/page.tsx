import { getStoreList } from '@/app/_api/store-list/get-store-list'
import { StoreListPresentational } from '@/app/_features/store-list/ui/store-list-presentational'

export default async function StoresPage() {
  const stores = await getStoreList()

  return <StoreListPresentational stores={stores} />
}
