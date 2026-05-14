import { ref } from "vue";
import type { User } from "../types/user";

type Props = {
  user: User;
  token: string;
};

const isAuthenticated = ref(!!localStorage.getItem("token"));

export const useAuth = () => {
  const setLocalStorage = ({ user, token }: Props) => {
    localStorage.setItem("user", JSON.stringify(user));
    localStorage.setItem("token", token);
    isAuthenticated.value = true;
  };

  const removeLocalStgage = () => {
    localStorage.removeItem("user");
    localStorage.removeItem("token");
    isAuthenticated.value = false;
  };

  const getLocalStorage = () => {
    const token = localStorage.getItem("token");
    return token;
  };

  return {
    isAuthenticated,
    setLocalStorage,
    removeLocalStgage,
    getLocalStorage,
  };
};
