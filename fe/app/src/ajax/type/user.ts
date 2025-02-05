interface user {
  refresh_token: string;
  access_token: string;
  username: string;
  email: string;
  gender: number;
  user_id: string;
  avatar: string;
}
export interface userInfo extends Partial<user> {}
