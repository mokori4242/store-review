'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import { type ResponseError } from '@/app/_types/response-error'

import { LoginFormPresentational } from './login-form-presentational'
import { type LoginFormSchema, loginFormSchema } from '../model/login-form-schema'
import { useLogin } from '../api/use-login'

export const LoginFormContainer = () => {
  const router = useRouter()
  const [successMessage, setSuccessMessage] = useState('')
  const [generalError, setGeneralError] = useState<string | undefined>()

  const {
    register,
    handleSubmit,
    formState: { errors }
  } = useForm<LoginFormSchema>({
    resolver: zodResolver(loginFormSchema),
    mode: 'onBlur',
    defaultValues: {
      email: '',
      password: ''
    }
  })

  const loginMutation = useLogin({
    onSuccess: () => {
      // JWTはHTTPOnly Cookieに保存されるため、localStorageへの保存は不要
      setSuccessMessage('ログインが完了しました！')
      setGeneralError(undefined)
      router.push('/')
    },
    onError: (error: ResponseError) => {
      if (error) {
        setGeneralError(error.message)
      }
      setSuccessMessage('')
    }
  })

  const handleFormSubmit = (data: LoginFormSchema) => {
    setSuccessMessage('')
    setGeneralError(undefined)

    loginMutation.mutate({
      email: data.email,
      password: data.password
    })
  }

  return (
    <LoginFormPresentational
      register={register}
      errors={errors}
      isLoading={loginMutation.isPending}
      successMessage={successMessage}
      generalError={generalError}
      onSubmit={handleSubmit(handleFormSubmit)}
    />
  )
}
