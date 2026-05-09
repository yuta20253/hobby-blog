import axios from "axios";
import type { Post, PostDetail } from "../../types/post";

type PostsResponse = {
  posts: Post[];
};

type PostDetailResponse = {
  post: PostDetail;
};

export const useFetchPosts = async (): Promise<Post[]> => {
  try {
    const url = import.meta.env.VITE_API_URL + "/api/posts";

    const res = await axios.get<PostsResponse>(url);

    return res.data.posts;
  } catch (error) {
    console.error(error);
    throw error;
  }
};

export const useFetchPost = async (id: number): Promise<PostDetail> => {
  try {
    const url = import.meta.env.VITE_API_URL + `/api/posts/${id}`;

    const res = await axios.get<PostDetailResponse>(url);

    return res.data.post;
  } catch (error) {
    console.error(error);
    throw error;
  }
};
