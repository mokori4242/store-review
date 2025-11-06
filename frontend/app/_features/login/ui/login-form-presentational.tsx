import { type UseFormRegister, type FieldErrors } from 'react-hook-form'

import { LabelInput } from '@/app/_components/auth/label-input'
import { FormResultMessage } from '@/app/_components/auth/form-result-message'

import { type LoginFormSchema } from '../model/login-form-schema'

interface LoginFormPresentationalProps {
  register: UseFormRegister<LoginFormSchema>
  errors: FieldErrors<LoginFormSchema>
  isLoading: boolean
  successMessage: string
  generalError?: string
  onSubmit: (e: React.FormEvent<HTMLFormElement>) => void
}

export const LoginFormPresentational = ({
  register,
  errors,
  isLoading,
  successMessage,
  generalError,
  onSubmit: handleSubmit
}: LoginFormPresentationalProps) => {
  return (
    <div className='flex min-h-screen items-center justify-center bg-gradient-to-br from-zinc-50 to-zinc-100 dark:from-zinc-900 dark:to-black px-4 py-12'>
      <div className='w-full max-w-md'>
        <div className='bg-white dark:bg-zinc-900 shadow-2xl rounded-2xl p-8 border border-zinc-200 dark:border-zinc-800'>
          <div className='mb-8 text-center'>
            <h1 className='text-3xl font-bold text-zinc-900 dark:text-zinc-50 mb-2'>ログイン</h1>
            <p className='text-zinc-600 dark:text-zinc-400'>アカウントにログインします</p>
          </div>

          {successMessage && (
            <FormResultMessage variant='success'>{successMessage}</FormResultMessage>
          )}

          {generalError && <FormResultMessage variant='error'>{generalError}</FormResultMessage>}

          <form onSubmit={handleSubmit} className='space-y-6'>
            <LabelInput
              variant='email'
              register={register}
              errors={errors}
              type='email'
              placeholder='example@email.com'
            >
              メールアドレス
            </LabelInput>
            <LabelInput
              variant='password'
              register={register}
              errors={errors}
              type='password'
              placeholder='8文字以上'
            >
              パスワード
            </LabelInput>
            <button
              type='submit'
              disabled={isLoading}
              className={`w-full py-3 px-4 rounded-lg font-medium text-white transition-colors ${
                isLoading
                  ? 'bg-zinc-400 cursor-not-allowed'
                  : 'bg-zinc-900 dark:bg-zinc-100 dark:text-zinc-900 hover:bg-zinc-800 dark:hover:bg-zinc-200'
              }`}
            >
              {isLoading ? 'ログイン中...' : 'ログイン'}
            </button>
          </form>

          <div className='mt-6 text-center'>
            <p className='text-sm text-zinc-600 dark:text-zinc-400'>
              アカウントをお持ちでないですか？{' '}
              <a
                href='/register'
                className='text-zinc-900 dark:text-zinc-100 font-medium hover:underline'
              >
                新規登録
              </a>
            </p>
          </div>
        </div>
      </div>
    </div>
  )
}
