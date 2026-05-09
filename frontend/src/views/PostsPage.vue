<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useFetchPosts } from "../service/post/postService";
import type { Post } from "../types/post";
import { theme } from "../styles/theme";

const posts = ref<Post[]>([]);
const isLoading = ref(true);
const isError = ref(false);

onMounted(async () => {
  try {
    posts.value = await useFetchPosts();
  } catch (e) {
    console.error(e);
    isError.value = true;
  } finally {
    isLoading.value = false;
  }
});
</script>

<template>
  <div class="container">
    <h2 class="title">投稿一覧</h2>
    <div v-if="isLoading" class="loading">読み込み中...</div>
    <div v-else-if="isError" class="error">投稿の取得に失敗しました</div>
    <div v-else class="card">
      <RouterLink
        v-for="post in posts"
        :key="post.id"
        class="card-content"
        :to="`/posts/${post.id}`"
      >
        <div class="post-row">
          <span class="post-category">
            {{ post.category?.name }}
          </span>

          <p class="post-title">
            {{ post.title }}
          </p>
        </div>

        <p class="post-user">👤 {{ post.user?.name }}</p>
      </RouterLink>
    </div>

    <p v-if="!isLoading && posts.length === 0" class="empty">
      投稿がありません
    </p>
  </div>
</template>

<style scoped>
.container {
  max-width: 720px;
  margin: 0 auto;
  padding: v-bind("theme.spacing.xl");
  background-color: v-bind("theme.colors.background");
  min-height: 100vh;
}

.title {
  font-size: v-bind("theme.fontSize['2xl']");
  font-weight: 700;
  margin-bottom: v-bind("theme.spacing.lg");
  color: v-bind("theme.colors.textPrimary");
}

/* 状態表示 */
.loading,
.error {
  text-align: center;
  margin-top: v-bind("theme.spacing.lg");
  color: v-bind("theme.colors.textSecondary");
}

.error {
  color: v-bind("theme.colors.danger");
}

/* 一覧 */
.card {
  display: flex;
  flex-direction: column;
  gap: v-bind("theme.spacing.md");
}

/* カード */
.card-content {
  display: block;
  padding: v-bind("theme.spacing.lg");
  border-radius: v-bind("theme.borderRadius.lg");
  background-color: white;
  border: 1px solid v-bind("theme.colors.borderLight");
  text-decoration: none;
  transition: all 0.2s ease;
}

/* hover */
.card-content:hover {
  transform: translateY(-3px);
  box-shadow: 0 10px 24px rgba(0, 0, 0, 0.08);
}

/* 上段 */
.post-row {
  display: flex;
  align-items: center;
  gap: v-bind("theme.spacing.sm");
  margin-bottom: v-bind("theme.spacing.sm");
}

/* カテゴリ */
.post-category {
  font-size: 11px;
  font-weight: 600;
  color: v-bind("theme.colors.primary");
  background-color: rgba(0, 123, 255, 0.1);
  padding: 4px 10px;
  border-radius: 999px;
  white-space: nowrap;
}

/* タイトル */
.post-title {
  font-size: 16px;
  font-weight: 700;
  color: v-bind("theme.colors.textPrimary");
  line-height: 1.5;
}

/* ユーザー */
.post-user {
  font-size: v-bind("theme.fontSize.sm");
  color: v-bind("theme.colors.textSecondary");
  margin-top: 4px;
}

/* 空状態 */
.empty {
  margin-top: v-bind("theme.spacing.lg");
  color: v-bind("theme.colors.textSecondary");
  text-align: center;
}
</style>
