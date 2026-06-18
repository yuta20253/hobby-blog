<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useFetchMyData } from "../service/mypage/mypageService";
import type { Post } from "../types/post";
import { useAuth } from "../composables/useAuth";
import type { User } from "../types/user";
import { theme } from "../styles/theme";

const user = ref<User | null>(null);
const posts = ref<Post[]>([]);
const isLoading = ref(true);
const isError = ref(false);
const BASE_URL = import.meta.env.VITE_API_URL;
const { getLocalStorage } = useAuth();

onMounted(async () => {
  try {
    const token = getLocalStorage();
    if (!token) {
      isError.value = true;
      return;
    }
    const { user: fetchedUser, posts: fetchedPosts } =
      await useFetchMyData(token);
    user.value = fetchedUser ?? null;
    posts.value = fetchedPosts ?? [];
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
    <h2 class="title">マイページ</h2>
    <div v-if="isLoading" class="loading">読み込み中...</div>
    <div v-else-if="isError" class="error">読み込みに失敗しました</div>
    <div v-else>
      <!-- ユーザー情報 -->
      <div v-if="user" class="user-section">
        <div class="user-card">
          <h3 class="user-name">{{ user.name }}</h3>
          <p class="user-email">{{ user.email }}</p>
        </div>
      </div>
      <div class="posts-section">
        <h3 class="section-title">あなたの投稿</h3>
        <div v-if="!posts || posts.length === 0" class="empty">投稿がありません</div>
        <div v-else class="posts-list">
          <div
            v-for="post in posts"
            :key="post.id"
            class="post-card"
          >
            <RouterLink
              :to="`/posts/${post.id}`"
              class="post-main"
            >
              <div class="post-header">
                <span class="post-category">{{ post.category?.name }}</span>
                <p class="post-title">{{ post.title }}</p>
              </div>
              <p class="post-preview">
                {{ post.content?.substring(0, 100) }}...
              </p>
            </RouterLink>
            <div v-if="post.media_files && post.media_files.length" class="media-list">
              <div
                v-for="media in post.media_files"
                :key="media.id"
                class="media-item"
              >
                <img
                  v-if="media.type === 'image'"
                  :src="BASE_URL + media.file_path"
                  :alt="media.file_name"
                  class="media-image"
                />
                <video
                  v-else-if="media.type === 'video'"
                  controls
                  class="media-video"
                >
                  <source :src="BASE_URL + media.file_path" type="video/mp4" />
                  お使いのブラウザは動画タグに対応していません。
                </video>
              </div>
            </div>
            <div class="post-actions">
              <RouterLink
                :to="`/posts/${post.id}/edit`"
                class="edit-button"
              >
                編集
              </RouterLink>
            </div>
          </div>
        </div>
      </div>
    </div>
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

/* 状態 */
.loading,
.error {
  text-align: center;
  margin-top: v-bind("theme.spacing.lg");
  color: v-bind("theme.colors.textSecondary");
}

.error {
  color: v-bind("theme.colors.danger");
}

/* ユーザー */
.user-section {
  margin-bottom: v-bind("theme.spacing.xl");
}

.user-card {
  background: white;
  border: 1px solid v-bind("theme.colors.borderLight");
  border-radius: v-bind("theme.borderRadius.lg");
  padding: v-bind("theme.spacing.lg");
}

.user-name {
  font-size: 18px;
  font-weight: 700;
  color: v-bind("theme.colors.textPrimary");
}

.user-email {
  font-size: v-bind("theme.fontSize.sm");
  color: v-bind("theme.colors.textSecondary");
  margin-top: 4px;
}

/* 投稿セクション */
.posts-section {
  margin-top: v-bind("theme.spacing.lg");
}

.section-title {
  font-size: 18px;
  font-weight: 700;
  margin-bottom: v-bind("theme.spacing.md");
  color: v-bind("theme.colors.textPrimary");
}

/* 投稿リスト */
.posts-list {
  display: flex;
  flex-direction: column;
  gap: v-bind("theme.spacing.md");
}

/* 投稿カード */
.post-card {
  display: block;
  padding: v-bind("theme.spacing.lg");
  border-radius: v-bind("theme.borderRadius.lg");
  background-color: white;
  border: 1px solid v-bind("theme.colors.borderLight");
  text-decoration: none;
  transition: all 0.2s ease;
}

.post-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 10px 24px rgba(0, 0, 0, 0.08);
}

/* ヘッダー */
.post-header {
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
}

/* タイトル */
.post-title {
  font-size: 16px;
  font-weight: 700;
  color: v-bind("theme.colors.textPrimary");
}

/* 本文プレビュー */
.post-preview {
  font-size: v-bind("theme.fontSize.sm");
  color: v-bind("theme.colors.textSecondary");
  line-height: 1.5;
}

/* メディア */
.media-list {
  display: flex;
  gap: 8px;
  margin-top: 8px;
  flex-wrap: wrap;
}

.media-item {
  max-width: 200px;
}

.media-image {
  width: 100%;
  height: auto;
}

.media-video {
  display: block;
  width: 100%;
  height: 100%;
  object-fit: cover;
}

/* 空状態 */
.empty {
  text-align: center;
  margin-top: v-bind("theme.spacing.lg");
  color: v-bind("theme.colors.textSecondary");
}

.post-main {
  display: block;
  text-decoration: none;
  color: inherit;
}

.post-actions {
  margin-top: 12px;
  display: flex;
  justify-content: flex-end;
}

.edit-button {
  font-size: 13px;
  padding: 6px 12px;
  border-radius: 6px;
  background-color: v-bind("theme.colors.primary");
  color: white;
  text-decoration: none;
  transition: all 0.2s ease;
}

.edit-button:hover {
  opacity: 0.85;
}
</style>
