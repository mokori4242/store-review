// proxy.ts
import { NextResponse, type NextRequest } from 'next/server'

export function proxy(request: NextRequest): NextResponse {
  // 認証トークンの有無をチェック
  const authToken = request.cookies.get('accessToken')
  if (!authToken) {
    // 未認証の場合はログインページへリダイレクト
    return NextResponse.redirect(new URL('/login', request.url))
  }
  // 認証済ならそのまま次の処理へ
  return NextResponse.next()
}

export const config = {
  matcher: ['/stores/:path*']
}
