import { type AxiosRequestConfig } from 'axios'

export enum Method {
  POST = 'post',
  GET = 'get',
  DELETE = 'delete',
  PUT = 'put'
}
export enum ContentType {
  Application = 'application/json',
  Multipart = 'multipart/form-data'
}
export interface IRequestParams {
  method: string
  url: string
  data?: any
  config?: AxiosRequestConfig<any> | undefined
  userId?: number
  contentType?: string
}
