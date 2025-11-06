import { useMutation, type UseMutationResult } from '@tanstack/react-query'

import { postRegister } from '@/app/_api/register/post-register'
import { type ResponseError } from '@/app/_types/response-error'

import { type RegisterRequestBody, type RegisterResponse } from '../model/types'

/**
 * ユーザー登録カスタムフックのオプション
 */
export interface UseRegisterOptions {
  onSuccess?: () => void
  onError?: (error: ResponseError) => void
}

/**
 * ユーザー登録カスタムフック
 * @param options - フックのオプション
 * @returns useMutationの結果
 */
export const useRegister = (
  options?: UseRegisterOptions
): UseMutationResult<RegisterResponse, ResponseError, RegisterRequestBody> => {
  return useMutation({
    mutationFn: (data: RegisterRequestBody) => postRegister(data),
    onSuccess: () => {
      options?.onSuccess?.()
    },
    onError: (error: ResponseError) => {
      options?.onError?.(error)
    }
  })
}
