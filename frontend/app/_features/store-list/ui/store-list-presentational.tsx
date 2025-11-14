import { type Store } from '../model/types'

interface Props {
  stores: Store[]
}

export const StoreListPresentational = ({ stores }: Props) => {
  if (stores.length === 0) {
    return (
      <div className='flex items-center justify-center min-h-screen'>
        <p className='text-gray-600'>店舗が見つかりませんでした</p>
      </div>
    )
  }

  return (
    <div className='container mx-auto px-4 py-8'>
      <h1 className='text-3xl font-bold mb-8'>店舗一覧</h1>
      <div className='grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6'>
        {stores.map((store) => (
          <div
            key={store.id}
            className='border border-gray-200 rounded-lg p-6 shadow-sm hover:shadow-md transition-shadow'
          >
            <h2 className='text-xl font-semibold mb-4'>{store.name}</h2>

            {store.categoryNames?.length > 0 && (
              <div className='mb-3'>
                <h3 className='text-sm font-medium text-white-700 mb-1'>カテゴリー</h3>
                <div className='flex flex-wrap gap-2'>
                  {store.categoryNames.map((category, index) => (
                    <span
                      key={index}
                      className='px-2 py-1 text-xs bg-blue-100 text-blue-800 rounded'
                    >
                      {category}
                    </span>
                  ))}
                </div>
              </div>
            )}

            {store.regularHolidays?.length > 0 && (
              <div className='mb-3'>
                <h3 className='text-sm font-medium text-white-700 mb-1'>定休日</h3>
                <p className='text-sm text-white-500'>{store.regularHolidays.join(', ')}</p>
              </div>
            )}

            {store.paymentMethods?.length > 0 && (
              <div className='mb-3'>
                <h3 className='text-sm font-medium text-white-700 mb-1'>支払い方法</h3>
                <p className='text-sm text-white-500'>{store.paymentMethods.join(', ')}</p>
              </div>
            )}

            {store.webProfiles?.length > 0 && (
              <div className='mb-3'>
                <h3 className='text-sm font-medium text-white-700 mb-1'>Webサイト</h3>
                <div className='flex flex-col gap-1'>
                  {store.webProfiles.map((url, index) => (
                    <a
                      key={index}
                      href={url}
                      target='_blank'
                      rel='noopener noreferrer'
                      className='text-sm text-blue-600 hover:underline'
                    >
                      {url}
                    </a>
                  ))}
                </div>
              </div>
            )}
          </div>
        ))}
      </div>
    </div>
  )
}
