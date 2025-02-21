export interface IArticle {
  article_id: number;
  content: string;
  title: string;
  user_id: string;
  category_id: number;
  topic_id: number;
  tag_id: number;
  cover: string;
  abstract: string;
  create_at: string;
  update_at: string;
  view_count: number;
}
export interface IAuthor {
  avatar: string;
  email: string;
  gender: number;
  user_id: number;
  username: string;
}

export interface IComment {
  article_id: number;
  comment_id: number;
  content: string;
  parent_comment_id: number;
  create_time: string;
  update_time: string;
  user_id: Number;
  user_info: UserInfo;
  second_comments?: IComment[];
  second_comments_status?: boolean;
}

export interface UserInfo {
  avatar: string;
  email: string;
  gender: number;
  user_id: number;
  username: string;
}
