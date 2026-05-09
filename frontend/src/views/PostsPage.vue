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
  <div>
    <h2>投稿一覧</h2>
    <div class="card">
      <div
        v-for="post in posts"
        :key="post.id"
        class="card-content"
      >
        <div>
          {{ post.title }}
          <button class="post-detail-button">詳細</button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.card {
  display: flex;
  flex-direction: column;
  gap: 16px;
  margin-top: 16px;
}

.card-content {
  padding: 16px;
  border-radius: v-bind('theme.borderRadius.md');
  transition: all 0.2s ease;
  background-color: v-bind('theme.colors.background');
  border: 1px solid v-bind('theme.colors.border');
  box-shadow: none;
}

.card-content:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.1);
}

.card-content > div {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.post-detail-button {
  padding: 6px 12px;
  border: none;
  border-radius: v-bind('theme.borderRadius.sm');
  background-color: v-bind('theme.colors.primary');
  color: white;
  font-size: v-bind('theme.fontSize.sm');
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.post-detail-button:hover {
  background-color: v-bind('theme.colors.primaryHover');
}
</style>
