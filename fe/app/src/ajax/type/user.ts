interface user {
  refresh_token: string
  access_token: string
  user_name: string
  email: string
  gender: number
  user_id: string
  avatar: string
  nickname: string
  overview: string
  create_at: string
  update_at: string
}

export interface UserInfo {
  avatar: string
  email: string
  gender: number
  user_id: number
  username: string
  nickname: string
  overview: string
}
export interface userInfo extends Partial<user> {}
