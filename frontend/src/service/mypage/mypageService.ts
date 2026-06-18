import axios from "axios";
import type { User } from "../../types/user";
import type { Post } from "../../types/post";

type MyPageResponse = {
  user: User;
  posts: Post[];
};

export const useFetchMyData = async (
  token: string
): Promise<{ user: User; posts: Post[] }> => {
  try {
    const url = import.meta.env.VITE_API_URL + "/api/mypage";
    const res = await axios.get<MyPageResponse>(url, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return {
      user: res.data.user,
      posts: res.data.posts ?? [],
    };
  } catch (error) {
    console.error(error);
    throw error;
  }
};
