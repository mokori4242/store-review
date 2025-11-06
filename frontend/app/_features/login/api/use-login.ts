import { useMutation, type UseMutationResult } from '@tanstack/react-query'

import { postLogin } from '@/app/_api/login/post-login'
import { type ResponseError } from '@/app/_types/response-error'

import { type LoginRequestBody, type LoginResponse } from '../model/types'

/**
 * ログインカスタムフックのオプション
 */
export interface UseLoginOptions {
  onSuccess?: (data: LoginResponse) => void
  onError?: (error: ResponseError) => void
}

/**
 * ログインカスタムフック
 * @param options - フックのオプション
 * @returns useMutationの結果
 */
export const useLogin = (
  options?: UseLoginOptions
): UseMutationResult<LoginResponse, ResponseError, LoginRequestBody> => {
  return useMutation({
    mutationFn: (data: LoginRequestBody) => postLogin(data),
    onSuccess: (data: LoginResponse) => {
      options?.onSuccess?.(data)
    },
    onError: (error: ResponseError) => {
      options?.onError?.(error)
    }
  })
}
