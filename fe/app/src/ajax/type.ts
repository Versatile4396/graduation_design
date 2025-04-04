import type { UserInfo } from '@/page/article/component/types'
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

export interface Assistance {
  abstract: string
  assistance_id: number
  category_id: 2
  content: string
  cover: string
  create_at: string
  tag_id: number
  title: string
  topic_id: number
  update_at: string
  user_id: number
  user_info: UserInfo
}
