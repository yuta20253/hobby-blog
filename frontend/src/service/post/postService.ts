import axios from "axios";
import type { Post } from "../../types/post";

type PostsResponse = {
  posts: Post[];
};

export const fetchPosts = async (): Promise<Post[]> => {
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
