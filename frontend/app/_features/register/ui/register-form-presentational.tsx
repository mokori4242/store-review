import { type UseFormRegister, type FieldErrors } from 'react-hook-form'

import { MessagePresentational } from '@/app/_features/register/ui/message-presentational'
import { LabelInputPresentational } from '@/app/_features/register/ui/label-input-presentational'

import { type RegisterFormSchema } from '../model/register-form-schema'

interface RegisterFormPresentationalProps {
  register: UseFormRegister<RegisterFormSchema>
  errors: FieldErrors<RegisterFormSchema>
  isLoading: boolean
  successMessage: string
  generalError?: string
  onSubmit: (e: React.FormEvent<HTMLFormElement>) => void
}

export const RegisterFormPresentational = ({
  register,
  errors,
  isLoading,
  successMessage,
  generalError,
  onSubmit: handleSubmit
}: RegisterFormPresentationalProps) => {
  return (
    <div className='flex min-h-screen items-center justify-center bg-gradient-to-br from-zinc-50 to-zinc-100 dark:from-zinc-900 dark:to-black px-4 py-12'>
      <div className='w-full max-w-md'>
        <div className='bg-white dark:bg-zinc-900 shadow-2xl rounded-2xl p-8 border border-zinc-200 dark:border-zinc-800'>
          <div className='mb-8 text-center'>
            <h1 className='text-3xl font-bold text-zinc-900 dark:text-zinc-50 mb-2'>新規登録</h1>
            <p className='text-zinc-600 dark:text-zinc-400'>アカウントを作成して始めましょう</p>
          </div>

          {successMessage && (
            <MessagePresentational variant='success'>{successMessage}</MessagePresentational>
          )}

          {generalError && (
            <MessagePresentational variant='error'>{generalError}</MessagePresentational>
          )}

          <form onSubmit={handleSubmit} className='space-y-6'>
            <LabelInputPresentational variant='nickname' register={register} errors={errors}>
              ニックネーム
            </LabelInputPresentational>
            <LabelInputPresentational variant='email' register={register} errors={errors}>
              メールアドレス
            </LabelInputPresentational>
            <LabelInputPresentational variant='password' register={register} errors={errors}>
              パスワード
            </LabelInputPresentational>
            <LabelInputPresentational variant='confirmPassword' register={register} errors={errors}>
              パスワード確認
            </LabelInputPresentational>
            <button
              type='submit'
              disabled={isLoading}
              className={`w-full py-3 px-4 rounded-lg font-medium text-white transition-colors ${
                isLoading
                  ? 'bg-zinc-400 cursor-not-allowed'
                  : 'bg-zinc-900 dark:bg-zinc-100 dark:text-zinc-900 hover:bg-zinc-800 dark:hover:bg-zinc-200'
              }`}
            >
              {isLoading ? '登録中...' : '登録する'}
            </button>
          </form>

          <div className='mt-6 text-center'>
            <p className='text-sm text-zinc-600 dark:text-zinc-400'>
              既にアカウントをお持ちですか？{' '}
              <a
                href='/login'
                className='text-zinc-900 dark:text-zinc-100 font-medium hover:underline'
              >
                ログイン
              </a>
            </p>
          </div>
        </div>
      </div>
    </div>
  )
}
