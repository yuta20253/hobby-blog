export type Post = {
  id: number;
  user_id: number;
  category_id: number;
  title: string;
  content: string;
  status: string;
  user: {
    id: number;
    name: string;
    email: string;
  };
  category: {
    id: number;
    name: string;
  };
};
