import { clsx } from 'clsx'
import  { type FieldErrors, type UseFormRegister } from 'react-hook-form'

import  { type RegisterFormSchema } from '@/app/_features/register/model/register-form-schema'

interface Props {
  variant: 'nickname' | 'email' | 'password' | 'confirmPassword'
  register: UseFormRegister<RegisterFormSchema>
  errors: FieldErrors<RegisterFormSchema>
  children: string
}

export const LabelInputPresentational = ({ variant, register, errors, children }: Props) => {
  const inputTypes = {
    nickname: 'text',
    email: 'email',
    password: 'password',
    confirmPassword: 'password'
  }

  const placeholders = {
    nickname: '山田太郎',
    email: 'example@email.com',
    password: '8文字以上',
    confirmPassword: 'パスワードを再入力'
  }

  const inputStyles = errors[variant]
    ? 'border-red-500 focus:ring-red-500'
    : 'border-zinc-300 dark:border-zinc-700 focus:ring-blue-500'

  return (
    <div>
      <label className='block text-sm font-medium text-zinc-900 dark:text-zinc-100 mb-2'>
        {children}
      </label>
      <input
        type={inputTypes[variant]}
        {...register(variant)}
        className={clsx(
          'w-full px-4 py-3 rounded-lg border bg-white dark:bg-zinc-800 text-zinc-900 dark:text-zinc-100 focus:ring-2 focus:outline-none transition-colors',
          inputStyles
        )}
        placeholder={placeholders[variant]}
      />
      {errors[variant] && (
        <p className='mt-1 text-sm text-red-600 dark:text-red-400'>{errors[variant].message}</p>
      )}
    </div>
  )
}
