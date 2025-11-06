import { z } from 'zod'

// Zodスキーマ定義
export const registerFormSchema = z
  .object({
    nickname: z
      .string()
      .min(1, 'ニックネームは必須です')
      .max(40, 'ニックネームは40文字以内で入力してください'),
    email: z
      .string()
      .min(1, 'メールアドレスは必須です')
      .email('有効なメールアドレスを入力してください'),
    password: z
      .string()
      .min(1, 'パスワードは必須です')
      .min(8, 'パスワードは8文字以上で入力してください'),
    confirmPassword: z.string().min(1, 'パスワード確認は必須です')
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: 'パスワードが一致しません',
    path: ['confirmPassword']
  })

// Zodスキーマから型を推論
export type RegisterFormSchema = z.infer<typeof registerFormSchema>
