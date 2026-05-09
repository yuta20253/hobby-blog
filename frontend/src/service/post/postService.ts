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
    const token = localStorage.getItem("token");
    const url = import.meta.env.VITE_API_URL + "/api/posts";

    const res = await axios.get<PostsResponse>(url, {
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    });

    return res.data.posts;
  } catch (error) {
    console.error(error);
    throw error;
  }
};

export const useFetchPost = async (id: number): Promise<PostDetail> => {
  try {
    const token = localStorage.getItem("token");
    const url = import.meta.env.VITE_API_URL + `/api/posts/${id}`;

    const res = await axios.get<PostDetailResponse>(url, {
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    });

    return res.data.post;
  } catch (error) {
    console.error(error);
    throw error;
  }
};
