import type { Category } from "./category";
import type { Media } from "./media";
import type { User } from "./user";

export type Post = {
  id: number;
  user_id: number;
  category_id: number;
  title: string;
  content: string;
  status: string;
  user: User;
  category: Category;
  media_files: Media[];
};
