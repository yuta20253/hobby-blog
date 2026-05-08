import { ref } from "vue";

const isAuthenticated = ref(!!localStorage.getItem("token"));

export const useAuth = () => {
    const setLogin = (token: string) => {
        localStorage.setItem("token", token);
        isAuthenticated.value = true;
    };

    const logout = () => {
        localStorage.removeItem("token");
        isAuthenticated.value = false;
    }

    return {
        isAuthenticated,
        setLogin,
        logout,
    };
};

