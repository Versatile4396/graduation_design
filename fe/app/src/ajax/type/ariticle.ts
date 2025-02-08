export enum FormType {
  create,
  editor,
  copy,
}

export interface IArticle {
  title: string;
  content: string;
  cover: string;
  topic_id: number;
  tag_id: number;
  category_id: number;
}
