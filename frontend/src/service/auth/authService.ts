import axios from "axios";
import { useAuth } from "../../composables/useAuth";
import type { User } from "../../types/user";

type Login = {
  email: string;
  password: string;
};

type SignUp = {
  name: string;
  email: string;
  password: string;
};

type AuthResponse = {
  user: User;
  token: string;
};

export const authService = () => {
  const loginService = async ({
    email,
    password,
  }: Login): Promise<AuthResponse> => {
    try {
      const url = import.meta.env.VITE_API_URL + "/api/auth/login";
      const res = await axios.post<AuthResponse>(url, { email, password });
      return res.data;
    } catch (error) {
      console.error(error);
      throw error;
    }
  };

  const logoutService = () => {
    const { removeLocalStgage } = useAuth();
    removeLocalStgage();
  };

  const signupService = async ({
    name,
    email,
    password,
  }: SignUp): Promise<AuthResponse> => {
    try {
      const url = import.meta.env.VITE_API_URL + "/api/auth/signup";
      const res = await axios.post<AuthResponse>(url, {
        name,
        email,
        password,
      });
      return res.data;
    } catch (error) {
      console.error(error);
      throw error;
    }
  };

  return { loginService, logoutService, signupService };
};
