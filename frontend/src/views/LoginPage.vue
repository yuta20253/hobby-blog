<script setup lang="ts">
    import { ref } from "vue";
    import { authService } from "../service/authService";
    import { theme } from "../styles/theme";
    import { useRouter } from "vue-router";
    import { useAuth } from "../composables/useAuth";


    const email = ref<string>("");
    const password = ref<string>("");
    const isPasswordVisible = ref<boolean>(false);
    const router = useRouter();

    const { setLocalStorage } = useAuth()

    const handleLogin = async () => {
    const { loginService } = authService()
    try {
        const {user, token} = await loginService({
          email: email.value,
          password: password.value,
        });

        setLocalStorage({user,token});

        router.push("/");

    } catch (error) {
        console.error(error);
    }
    };
</script>

<template>
  <div class="container">
    <div class="card">
      <h2 class="title">ログイン</h2>

      <p class="description">
        アカウントにログインしてください。
      </p>

      <form class="form" @submit.prevent="handleLogin" autocomplete="off">
        <div class="form-group">
          <label>メールアドレス</label>

          <input
            name="login_email_input"
            type="email"
            v-model="email"
            placeholder="test@example.com"
            autocomplete="off"
          />
        </div>

        <div class="form-group">
          <label>パスワード</label>

          <div class="password-wrapper">
            <input
                name="login_password_input"
                :type="isPasswordVisible ? 'text' : 'password'"
                v-model="password"
                placeholder="********"
                autocomplete="new-password"
            />

            <button
                type="button"
                class="toggle-button"
                @click="isPasswordVisible = !isPasswordVisible"
            >
                {{ isPasswordVisible ? '非表示' : '表示' }}
            </button>
          </div>
        </div>

        <div class="button-group">
          <button class="login-button" type="submit">
            ログイン
          </button>

          <button class="cancel-button" type="button">
            キャンセル
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped>
.container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: v-bind('theme.colors.background');
  padding: v-bind('theme.spacing.xl');
}

.card {
  width: 100%;
  max-width: 420px;
  background-color: white;
  border-radius: v-bind('theme.borderRadius.lg');
  padding: v-bind('theme.spacing.xxl');
  border: 1px solid v-bind('theme.colors.borderLight');
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.08);
}

.title {
  font-size: v-bind("theme.fontSize['2xl']");
  font-weight: bold;
  margin-bottom: v-bind('theme.spacing.sm');
  color: v-bind('theme.colors.textPrimary');
}

.description {
  color: v-bind('theme.colors.textSecondary');
  margin-bottom: v-bind('theme.spacing.xl');
  font-size: v-bind('theme.fontSize.base');
}

.form {
  display: flex;
  flex-direction: column;
  gap: v-bind('theme.spacing.lg');
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-group label {
  margin-bottom: v-bind('theme.spacing.sm');
  font-size: v-bind('theme.fontSize.base');
  font-weight: 600;
  color: v-bind('theme.colors.textPrimary');
}

.form-group input {
  height: 44px;
  border: 1px solid v-bind('theme.colors.border');
  border-radius: v-bind('theme.borderRadius.md');
  padding: 0 v-bind('theme.spacing.md');
  font-size: v-bind('theme.fontSize.base');
  transition: all v-bind('theme.transition.base');
  background-color: white;
}

.form-group input:focus {
  outline: none;
  border-color: v-bind('theme.colors.primary');
  box-shadow: 0 0 0 3px rgba(0, 123, 255, 0.15);
}

.button-group {
  display: flex;
  gap: v-bind('theme.spacing.md');
  margin-top: v-bind('theme.spacing.sm');
}

.login-button,
.cancel-button {
  flex: 1;
  height: 44px;
  border: none;
  border-radius: v-bind('theme.borderRadius.md');
  font-size: v-bind('theme.fontSize.base');
  font-weight: bold;
  cursor: pointer;
  transition: all v-bind('theme.transition.base');
}

.login-button {
  background-color: v-bind('theme.colors.primary');
  color: white;
}

.login-button:hover {
  background-color: v-bind('theme.colors.primaryHover');
}

.cancel-button {
  background-color: v-bind('theme.colors.backgroundDark');
  color: v-bind('theme.colors.textPrimary');
}

.cancel-button:hover {
  background-color: v-bind('theme.colors.borderLight');
}

.password-wrapper {
  position: relative;
  display: flex;
  align-items: center;
}

.password-wrapper input {
  width: 100%;
}

.toggle-button {
  position: absolute;
  right: 12px;
  background: none;
  border: none;
  color: v-bind('theme.colors.primary');
  cursor: pointer;
  font-size: v-bind('theme.fontSize.sm');
}
</style>
