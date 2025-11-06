import { clsx } from 'clsx'
import {
  type FieldErrors,
  type FieldValues,
  type Path,
  type UseFormRegister
} from 'react-hook-form'

interface Props<T extends FieldValues> {
  variant: Path<T>
  register: UseFormRegister<T>
  errors: FieldErrors<T>
  children: string
  type?: 'text' | 'email' | 'password'
  placeholder?: string
}

export const LabelInput = <T extends FieldValues>({
  variant,
  register,
  errors,
  children,
  type = 'text',
  placeholder
}: Props<T>) => {
  const error = errors[variant]

  const inputStyles = error
    ? 'border-red-500 focus:ring-red-500'
    : 'border-zinc-300 dark:border-zinc-700 focus:ring-blue-500'

  return (
    <div>
      <label className='block text-sm font-medium text-zinc-900 dark:text-zinc-100 mb-2'>
        {children}
      </label>
      <input
        type={type}
        {...register(variant)}
        className={clsx(
          'w-full px-4 py-3 rounded-lg border bg-white dark:bg-zinc-800 text-zinc-900 dark:text-zinc-100 focus:ring-2 focus:outline-none transition-colors',
          inputStyles
        )}
        placeholder={placeholder}
      />
      {error && (
        <p className='mt-1 text-sm text-red-600 dark:text-red-400'>{error.message as string}</p>
      )}
    </div>
  )
}
