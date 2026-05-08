<script setup lang="ts">
import { authService } from "../service/authService";
import { theme } from "../styles/theme";
import { computed, ref } from "vue";
import { useAuth } from "../composables/useAuth";
import { useRouter } from "vue-router";

const name = ref<string>("");
const email = ref<string>("");
const password = ref<string>("");
const passwordConfirmation = ref<string>("");

const isSubmitted = ref<boolean>(false);

const isPasswordVisible = ref<boolean>(false);
const isPasswordConfirmationVisible = ref<boolean>(false);

const { setLocalStorage } = useAuth();

const router = useRouter();

const nameError = computed(() => {
  if (!name.value) {
    return "名前を入力してください";
  }

  return "";
});

const emailError = computed(() => {
  if (!email.value) {
    return "メールアドレスを入力してください";
  }

  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

  if (!emailRegex.test(email.value)) {
    return "メールアドレス形式が正しくありません";
  }

  return "";
});

const passwordError = computed(() => {
  if (!password.value) {
    return "パスワードを入力してください";
  }

  if (password.value.length < 8) {
    return "8文字以上で入力してください";
  }

  if (password.value !== passwordConfirmation.value) {
    return "パスワードが一致しません";
  }

  return "";
});

const isValid = computed(() => {
  return !nameError.value && !emailError.value && !passwordError.value;
});

const handleSubmit = async () => {
  isSubmitted.value = true;

  if (!isValid.value) {
    return;
  }

  const { signupService } = authService();

  try {
    const { user, token } = await signupService({
      name: name.value,
      email: email.value,
      password: password.value,
    });

    setLocalStorage({ user, token });

    router.push("/");
  } catch (error) {
    console.error(error);
  }
};
</script>

<template>
  <div class="container">
    <div class="card">
      <h2 class="title">新規登録</h2>

      <p class="description">アカウントを作成してください。</p>

      <form class="form" autocomplete="off" @submit.prevent="handleSubmit">
        <div class="form-group">
          <label>名前</label>

          <input
            v-model="name"
            name="signup_name_input"
            type="text"
            placeholder="田中 太郎"
            autocomplete="off"
          />

          <p v-if="isSubmitted && nameError" class="error">
            {{ nameError }}
          </p>
        </div>

        <div class="form-group">
          <label>メールアドレス</label>

          <input
            v-model="email"
            name="signup_email_input"
            type="email"
            placeholder="test@example.com"
            autocomplete="off"
          />

          <p v-if="isSubmitted && emailError" class="error">
            {{ emailError }}
          </p>
        </div>

        <div class="form-group">
          <label>パスワード</label>

          <div class="password-wrapper">
            <input
              v-model="password"
              name="signup_password_input"
              :type="isPasswordVisible ? 'text' : 'password'"
              placeholder="********"
              autocomplete="new-password"
            />

            <button
              type="button"
              class="toggle-button"
              @click="isPasswordVisible = !isPasswordVisible"
            >
              {{ isPasswordVisible ? "非表示" : "表示" }}
            </button>
          </div>

          <p v-if="isSubmitted && passwordError" class="error">
            {{ passwordError }}
          </p>
        </div>

        <div class="form-group">
          <label>パスワード(確認用)</label>

          <div class="password-wrapper">
            <input
              v-model="passwordConfirmation"
              name="signup_password_confirmation_input"
              :type="isPasswordConfirmationVisible ? 'text' : 'password'"
              placeholder="********"
              autocomplete="new-password"
            />

            <button
              type="button"
              class="toggle-button"
              @click="
                isPasswordConfirmationVisible = !isPasswordConfirmationVisible
              "
            >
              {{ isPasswordConfirmationVisible ? "非表示" : "表示" }}
            </button>
          </div>
        </div>

        <div class="button-group">
          <button class="signup-button" type="submit">新規登録</button>

          <button class="cancel-button" type="button">キャンセル</button>
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
  background-color: v-bind("theme.colors.background");
  padding: v-bind("theme.spacing.xl");
}

.card {
  width: 100%;
  max-width: 420px;
  background-color: white;
  border-radius: v-bind("theme.borderRadius.lg");
  padding: v-bind("theme.spacing.xxl");
  border: 1px solid v-bind("theme.colors.borderLight");
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.08);
}

.title {
  font-size: v-bind("theme.fontSize['2xl']");
  font-weight: bold;
  margin-bottom: v-bind("theme.spacing.sm");
  color: v-bind("theme.colors.textPrimary");
}

.description {
  color: v-bind("theme.colors.textSecondary");
  margin-bottom: v-bind("theme.spacing.xl");
  font-size: v-bind("theme.fontSize.base");
}

.form {
  display: flex;
  flex-direction: column;
  gap: v-bind("theme.spacing.lg");
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-group label {
  margin-bottom: v-bind("theme.spacing.sm");
  font-size: v-bind("theme.fontSize.base");
  font-weight: 600;
  color: v-bind("theme.colors.textPrimary");
}

.form-group input {
  width: 100%;
  height: 44px;
  border: 1px solid v-bind("theme.colors.border");
  border-radius: v-bind("theme.borderRadius.md");
  padding: 0 48px 0 v-bind("theme.spacing.md");
  font-size: v-bind("theme.fontSize.base");
  transition: all v-bind("theme.transition.base");
  background-color: white;
  box-sizing: border-box;
}

.form-group input:focus {
  outline: none;
  border-color: v-bind("theme.colors.primary");
  box-shadow: 0 0 0 3px rgba(0, 123, 255, 0.15);
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
  color: v-bind("theme.colors.primary");
  cursor: pointer;
  font-size: v-bind("theme.fontSize.sm");
}

.toggle-button:hover {
  opacity: 0.7;
}

.error {
  margin-top: 6px;
  color: v-bind("theme.colors.danger");
  font-size: v-bind("theme.fontSize.sm");
}

.button-group {
  display: flex;
  gap: v-bind("theme.spacing.md");
  margin-top: v-bind("theme.spacing.sm");
}

.signup-button,
.cancel-button {
  flex: 1;
  height: 44px;
  border: none;
  border-radius: v-bind("theme.borderRadius.md");
  font-size: v-bind("theme.fontSize.base");
  font-weight: bold;
  cursor: pointer;
  transition: all v-bind("theme.transition.base");
}

.signup-button {
  background-color: v-bind("theme.colors.primary");
  color: white;
}

.signup-button:hover {
  background-color: v-bind("theme.colors.primaryHover");
}

.cancel-button {
  background-color: v-bind("theme.colors.backgroundDark");
  color: v-bind("theme.colors.textPrimary");
}

.cancel-button:hover {
  background-color: v-bind("theme.colors.borderLight");
}
</style>
