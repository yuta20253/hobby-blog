<script setup lang="ts">
import { onMounted, ref } from "vue";
import { theme } from "../styles/theme";
import { useFetchPost } from "../service/post/postService";
import type { PostDetail } from "../types/post";
import { useRoute } from "vue-router";
import { useRouter } from "vue-router";

const post = ref<PostDetail | null>(null);
const route = useRoute();
const router = useRouter();

const postId = route.params.id;

onMounted(async () => {
  post.value = await useFetchPost(Number(postId));
});

const handleBack = () => {
  router.push("/posts");
};
</script>

<template>
  <div class="container">
    <h2 class="title">投稿詳細</h2>

    <div v-if="post" class="card">
      <div class="category">
        {{ post.category.name }}
      </div>

      <div class="post-title">
        {{ post.title }}
      </div>

      <div class="meta">
        <span>👤 {{ post.user.name }}</span>
        <span>・</span>
        <span>{{ post.user.email }}</span>
      </div>

      <div class="content">
        {{ post.content }}
      </div>
    </div>

    <p v-else class="loading">読み込み中...</p>
    <div class="return-button">
      <button @click="handleBack">一覧へ戻る</button>
    </div>
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

/* カード */
.card {
  padding: v-bind("theme.spacing.xl");
  border-radius: v-bind("theme.borderRadius.lg");
  background-color: white;
  border: 1px solid v-bind("theme.colors.border");
  display: flex;
  flex-direction: column;
  gap: v-bind("theme.spacing.lg");
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.06);
  align-items: stretch;
}

/* カテゴリ（ちょいリッチに） */
.category {
  display: inline-block;
  font-size: v-bind("theme.fontSize.sm");
  color: v-bind("theme.colors.primary");
  background-color: rgba(0, 123, 255, 0.1);
  padding: 4px 10px;
  border-radius: 999px;
  width: fit-content;
  font-weight: 600;
}

/* タイトル（主役にする） */
.post-title {
  font-size: 22px;
  font-weight: 700;
  color: v-bind("theme.colors.textPrimary");
  line-height: 1.4;
}

/* メタ情報 */
.meta {
  font-size: v-bind("theme.fontSize.sm");
  color: v-bind("theme.colors.textSecondary");
  display: flex;
  gap: 8px;
  align-items: center;
}

/* 本文 */
.content {
  width: 100%;
  margin-top: v-bind("theme.spacing.sm");
  font-size: 15px;
  color: v-bind("theme.colors.textPrimary");
  line-height: 1.8;
  white-space: pre-wrap;
  border-top: 1px solid v-bind("theme.colors.borderLight");
  padding-top: v-bind("theme.spacing.md");
  text-align: left;
}

/* ローディング */
.loading {
  color: v-bind("theme.colors.textSecondary");
  text-align: center;
}

.return-button {
  margin-top: v-bind("theme.spacing.lg");
  display: flex;
  justify-content: flex-start;
}

.return-button button {
  padding: 8px 14px;
  border-radius: v-bind("theme.borderRadius.md");
  border: 1px solid v-bind("theme.colors.border");
  background-color: white;
  color: v-bind("theme.colors.textPrimary");
  font-size: v-bind("theme.fontSize.sm");
  cursor: pointer;
  transition: all 0.2s ease;
}

.return-button button:hover {
  background-color: v-bind("theme.colors.backgroundDark");
}
</style>
