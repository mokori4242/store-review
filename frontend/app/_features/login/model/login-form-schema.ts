import { z } from 'zod'

// Zodスキーマ定義
export const loginFormSchema = z.object({
  email: z.email('有効なメールアドレスを入力してください').min(1, 'メールアドレスは必須です'),
  password: z
    .string()
    .min(1, 'パスワードは必須です')
    .min(8, 'パスワードは8文字以上で入力してください')
})

// Zodスキーマから型を推論
export type LoginFormSchema = z.infer<typeof loginFormSchema>
