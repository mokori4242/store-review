export class ResponseError extends Error {
  public response: Response
  /* eslint-disable @typescript-eslint/no-explicit-any */
  public customMessage: any // エラーメッセージがオブジェクトの場合などを考慮した、エラーメッセージ格納するフィールド
  constructor(message: string | object, response: Response) {
    super(isString(message) ? message : JSON.stringify(message))
    this.customMessage = message

    this.response = response
  }
}

/**
 * 値が文字列型かどうかをチェックする型ガード関数。
 * @param value チェックする値
 * @returns 値が文字列の場合はtrue、そうでない場合はfalse
 */
export const isString = (value: unknown): value is string => {
  return typeof value === 'string'
}
