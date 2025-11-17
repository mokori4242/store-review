'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import { type ResponseError } from '@/app/_types/response-error'

import { RegisterFormPresentational } from './register-form-presentational'
import { type RegisterFormSchema, registerFormSchema } from '../model/register-form-schema'
import { useRegister } from '../api/use-register'

export const RegisterFormContainer = () => {
  const router = useRouter()
  const [successMessage, setSuccessMessage] = useState('')
  const [generalError, setGeneralError] = useState<string | undefined>()

  const {
    register,
    handleSubmit,
    formState: { errors }
  } = useForm<RegisterFormSchema>({
    resolver: zodResolver(registerFormSchema),
    mode: 'onBlur',
    defaultValues: {
      nickname: '',
      email: '',
      password: '',
      confirmPassword: ''
    }
  })

  const registerMutation = useRegister({
    onSuccess: () => {
      setSuccessMessage('登録が完了しました！')
      setGeneralError(undefined)
      router.push('/login')
    },
    onError: (error: ResponseError) => {
      if (error) {
        setGeneralError(error.message)
      }
      setSuccessMessage('')
    }
  })

  const handleFormSubmit = (data: RegisterFormSchema) => {
    setSuccessMessage('')
    setGeneralError(undefined)

    registerMutation.mutate({
      nickname: data.nickname,
      email: data.email,
      password: data.password
    })
  }

  return (
    <RegisterFormPresentational
      register={register}
      errors={errors}
      isLoading={registerMutation.isPending}
      successMessage={successMessage}
      generalError={generalError}
      onSubmit={handleSubmit(handleFormSubmit)}
    />
  )
}
