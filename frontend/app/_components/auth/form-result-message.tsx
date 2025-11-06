import { clsx } from 'clsx'

interface Props {
  children: string
  variant: 'success' | 'error'
}

export const FormResultMessage = ({ children, variant }: Props) => {
  const divStyles = {
    success: 'bg-green-50 dark:bg-green-900/20 border-green-200 dark:border-green-800',
    error: 'bg-red-50 dark:bg-red-900/20 border-red-200 dark:border-red-800'
  }

  const textStyles = {
    success: 'text-green-800 dark:text-green-300',
    error: 'text-red-800 dark:text-red-300'
  }

  return (
    <div className={clsx('mb-6 p-4 rounded-lg border', divStyles[variant])}>
      <p className={clsx('text-sm', textStyles[variant])}>{children}</p>
    </div>
  )
}
