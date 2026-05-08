<script setup lang="ts">
import Header from "./components/Header.vue";
import { theme } from "./styles/theme";
</script>

<template>
  <div class="app">
    <Header />
    <main class="main">
      <router-view />
    </main>
  </div>
</template>

<style scoped>
.app {
  min-height: 100vh;
  background-color: v-bind("theme.colors.background");
}

.main {
  padding: v-bind("theme.spacing.lg");
  max-width: 1200px;
  margin: 0 auto;
}
</style>
