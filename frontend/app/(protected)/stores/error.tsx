'use client'

export default function StoresError() {
  return (
    <div className='flex min-h-screen items-center justify-center bg-gradient-to-br from-zinc-50 to-zinc-100 dark:from-zinc-900 dark:to-black px-4 py-12'>
      <div className='w-full max-w-md'>
        <div className='bg-white dark:bg-zinc-900 shadow-2xl rounded-2xl p-8 border border-zinc-200 dark:border-zinc-800'>
          <div className='text-center'>
            <h1 className='text-3xl font-bold text-zinc-900 dark:text-zinc-50 mb-2'>
              エラーが発生しました
            </h1>
            <p className='text-zinc-600 dark:text-zinc-400'>時間を置いてアクセスしてください</p>
          </div>
        </div>
      </div>
    </div>
  )
}
