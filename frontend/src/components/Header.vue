<script setup lang="ts">
import { onMounted } from "vue";
import { theme } from "../styles/theme";
import { useAuth } from "../composables/useAuth";
import { authService } from "../service/auth/authService";
import { useRouter } from "vue-router";

const { isAuthenticated } = useAuth();
const { logoutService } = authService();
const router = useRouter();

onMounted(() => {
  isAuthenticated.value = !!localStorage.getItem("token");
});

const handleLogout = () => {
  logoutService();
  router.push("/login");
};
</script>

<template>
  <header class="header">
    <div class="logo">Hobby Blog</div>

    <nav class="nav">
      <RouterLink to="/" class="nav-link"> Home </RouterLink>

      <RouterLink to="/posts" class="nav-link"> Posts </RouterLink>

      <template v-if="isAuthenticated">
        <RouterLink to="/mypage" class="nav-link"> My Page </RouterLink>

        <button
          type="button"
          class="nav-link logout-button"
          @click="handleLogout"
        >
          Logout
        </button>
      </template>

      <template v-else>
        <RouterLink to="/login" class="nav-link"> Login </RouterLink>

        <RouterLink to="/signup" class="nav-link"> SignUp </RouterLink>
      </template>
    </nav>
  </header>
</template>

<style scoped>
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: v-bind("theme.spacing.lg");
  border-bottom: 1px solid v-bind("theme.colors.border");
  background-color: v-bind("theme.colors.background");
}

.logo {
  font-size: v-bind("theme.fontSize['2xl']");
  font-weight: bold;
  color: v-bind("theme.colors.textPrimary");
}

.nav {
  display: flex;
  align-items: center;
  gap: v-bind("theme.spacing.md");
}

.nav-link {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 40px;
  padding: 0 v-bind("theme.spacing.md");
  border: none;
  border-radius: v-bind("theme.borderRadius.base");
  background: transparent;
  text-decoration: none;
  color: v-bind("theme.colors.primary");
  font-size: v-bind("theme.fontSize.base");
  font-weight: 600;
  cursor: pointer;
  transition: all v-bind("theme.transition.base");
}

.nav-link:not(.logout-button):hover {
  background-color: v-bind("theme.colors.backgroundDark");
}

.logout-button:hover {
  background-color: v-bind("theme.colors.dangerHover");
}
</style>
