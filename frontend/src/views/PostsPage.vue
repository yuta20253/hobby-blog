<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useFetchPosts } from "../service/post/postService";
import type { Post } from "../types/post";
import { theme } from "../styles/theme";

const posts = ref<Post[]>([]);

onMounted(async () => {
  posts.value = await useFetchPosts();
});
</script>

<template>
  <div class="container">
    <h2 class="title">投稿一覧</h2>

    <div class="card">
      <RouterLink
        v-for="post in posts"
        :key="post.id"
        class="card-content"
        :to="`/posts/${post.id}`"
      >
        <div class="post-row">
          <p class="post-title">{{ post.title }}</p>
        </div>
      </RouterLink>
    </div>

    <p v-if="posts.length === 0" class="empty">投稿がありません</p>
  </div>
</template>

<style scoped>
.container {
  max-width: 720px;
  margin: 0 auto;
  padding: v-bind("theme.spacing.xl");
}

.title {
  font-size: v-bind("theme.fontSize['2xl']");
  font-weight: bold;
  margin-bottom: v-bind("theme.spacing.lg");
  color: v-bind("theme.colors.textPrimary");
}

.card {
  display: flex;
  flex-direction: column;
  gap: v-bind("theme.spacing.md");
}

.card-content {
  padding: v-bind("theme.spacing.md");
  border-radius: v-bind("theme.borderRadius.md");
  transition: all 0.2s ease;
  background-color: v-bind("theme.colors.background");
  border: 1px solid v-bind("theme.colors.border");
  text-decoration: none;
}

.card-content:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.1);
}

.post-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: v-bind("theme.spacing.md");
}

.post-title {
  font-size: v-bind("theme.fontSize.base");
  font-weight: 600;
  color: v-bind("theme.colors.textPrimary");
}

.empty {
  margin-top: v-bind("theme.spacing.lg");
  color: v-bind("theme.colors.textSecondary");
  text-align: center;
}
</style>
