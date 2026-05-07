<script setup lang="ts">
    import { computed } from 'vue';
    import { useRouter } from 'vue-router';
    import { theme } from '../styles/theme';

    const router = useRouter();

    const token = localStorage.getItem('token');

    const isAuthenticated = computed(() => !!token);

    const handleLogout = () => {
        localStorage.removeItem("token");
        router.push("/login");
    }

</script>

<template>
    <header class="header">
        <div class="logo">Hobby Blog</div>
        <nav class="nav">
            <RouterLink to="/">Home</RouterLink>
            <RouterLink to="/posts">Posts</RouterLink>

            <template v-if="isAuthenticated">
                <RouterLink to="/mypage">My Page</RouterLink>
                <button @click="handleLogout">Logout</button>
            </template>
            <template v-else>
                <RouterLink to="/login">Login</RouterLink>
                <RouterLink to="/signup">SignUp</RouterLink>
            </template>
        </nav>
    </header>
</template>

<style scoped>
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: v-bind('theme.spacing.lg');
  border-bottom: 1px solid v-bind('theme.colors.border');
  background-color: v-bind('theme.colors.background');
}

.logo {
  font-size: v-bind("theme.fontSize['2xl']");
  font-weight: bold;
  color: v-bind('theme.colors.textPrimary');
}

.nav {
  display: flex;
  gap: v-bind('theme.spacing.md');
  align-items: center;
}

.nav a {
  text-decoration: none;
  color: v-bind('theme.colors.primary');
  padding: v-bind('theme.spacing.sm') v-bind('theme.spacing.md');
  border-radius: v-bind('theme.borderRadius.base');
  transition: background-color v-bind('theme.transition.base');
}

.nav a:hover {
  background-color: v-bind('theme.colors.backgroundDark');
}

.nav button {
  background-color: v-bind('theme.colors.danger');
  color: white;
  border: none;
  padding: v-bind('theme.spacing.sm') v-bind('theme.spacing.md');
  border-radius: v-bind('theme.borderRadius.base');
  cursor: pointer;
  transition: background-color v-bind('theme.transition.base');
}

.nav button:hover {
  background-color: v-bind('theme.colors.dangerHover');
}
</style>
