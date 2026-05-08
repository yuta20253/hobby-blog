import axios from "axios";
import { useAuth } from "../composables/useAuth";

type Login = {
    email: string;
    password: string;
};

type User = {
    id: number;
    name: string;
    email: string;
};

type LoginResponse = {
    user: User;
    token: string;
};

export const authService = () => {
    const loginService = async ({ email, password }: Login): Promise<LoginResponse> => {
        try {
            const url = import.meta.env.VITE_API_URL + '/api/auth/login';
            const res = await axios.post<LoginResponse>(url, { email, password });
            return res.data;
        } catch (error) {
            console.error(error);
            throw error;
        }
    };

    const logoutService = () => {
        const { logout } = useAuth()
        logout();
    }

    return { loginService, logoutService };
};
