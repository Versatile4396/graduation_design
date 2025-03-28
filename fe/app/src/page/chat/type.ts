import type { userInfo } from '@/ajax/type/user'

const enum ReadStatus {
  Unread = 0,
  Read = 1
}
export interface ChatMessage {
  Content: string
  Read: ReadStatus
  StartTime: number
  EndTime: number
}

export interface ChatInstance {
  from: MsgFrom
  code: number
  content: string
}

export const enum MsgFrom {
  Me = 'me',
  You = 'you'
}

export interface IChatList {
  latestMsg: string
  time: number
  userInfo: userInfo
}
