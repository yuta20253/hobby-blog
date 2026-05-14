<script setup lang="ts">
import { ref, onMounted, computed } from "vue";
import { useFetchMyData } from "../service/mypage/mypageService";
import type { Post } from "../types/post";
import { useAuth } from "../composables/useAuth";
import type { User } from "../types/user";
import { theme } from "../styles/theme";

const user = ref<User | null>(null);
const posts = ref<Post[]>([]);
const isLoading = ref(true);
const isLoadingMore = ref(false);
const isError = ref(false);
const { getLocalStorage } = useAuth();
const limit = 10;
const currentOffset = ref(0);
const hasMore = computed(
  () => posts.value.length % limit === 0 && posts.value.length > 0
);

const loadInitialData = async () => {
  try {
    const token = getLocalStorage();
    if (!token) {
      isError.value = true;
      return;
    }
    const data = await useFetchMyData(token, limit, 0);
    user.value = data.user;
    posts.value = data.posts;
    currentOffset.value = 0;
  } catch (e) {
    console.error(e);
    isError.value = true;
  } finally {
    isLoading.value = false;
  }
};

const loadMorePosts = async () => {
  try {
    const token = getLocalStorage();
    if (!token) {
      isError.value = true;
      return;
    }
    isLoadingMore.value = true;
    const newOffset = currentOffset.value + limit;
    const data = await useFetchMyData(token, limit, newOffset);
    posts.value.push(...data.posts);
    currentOffset.value = newOffset;
  } catch (e) {
    console.error(e);
    isError.value = true;
  } finally {
    isLoadingMore.value = false;
  }
};

onMounted(() => {
  loadInitialData();
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

      <!-- 投稿一覧 -->
      <div class="posts-section">
        <h3 class="section-title">あなたの投稿</h3>
        <div v-if="posts.length === 0" class="empty">投稿がありません</div>
        <div v-else class="posts-list">
          <RouterLink
            v-for="post in posts"
            :key="post.id"
            class="post-card"
            :to="`/posts/${post.id}`"
          >
            <div class="post-header">
              <span class="post-category">{{ post.category?.name }}</span>
              <p class="post-title">{{ post.title }}</p>
            </div>
            <p class="post-preview">{{ post.content?.substring(0, 100) }}...</p>
          </RouterLink>
        </div>

        <!-- 続きを読むボタン -->
        <div v-if="hasMore" class="load-more-container">
          <button
            class="load-more-btn"
            :disabled="isLoadingMore"
            @click="loadMorePosts"
          >
            {{ isLoadingMore ? "読み込み中..." : "続きを読む" }}
          </button>
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

.loading,
.error {
  text-align: center;
  margin-top: v-bind("theme.spacing.lg");
  color: v-bind("theme.colors.textSecondary");
}

.error {
  color: v-bind("theme.colors.danger");
}

.user-section {
  margin-bottom: v-bind("theme.spacing.xl");
}

.user-card {
  padding: v-bind("theme.spacing.lg");
  background-color: white;
  border: 1px solid v-bind("theme.colors.borderLight");
  border-radius: v-bind("theme.borderRadius.lg");
}

.user-name {
  font-size: v-bind("theme.fontSize.lg");
  font-weight: 700;
  color: v-bind("theme.colors.textPrimary");
  margin-bottom: v-bind("theme.spacing.sm");
  margin: 0;
}

.user-email {
  font-size: v-bind("theme.fontSize.sm");
  color: v-bind("theme.colors.textSecondary");
  margin: 0;
}

.posts-section {
  margin-top: v-bind("theme.spacing.lg");
}

.section-title {
  font-size: v-bind("theme.fontSize.lg");
  font-weight: 700;
  color: v-bind("theme.colors.textPrimary");
  margin-bottom: v-bind("theme.spacing.md");
  margin-top: 0;
}

.posts-list {
  display: flex;
  flex-direction: column;
  gap: v-bind("theme.spacing.md");
}

.post-card {
  display: block;
  padding: v-bind("theme.spacing.lg");
  background-color: white;
  border: 1px solid v-bind("theme.colors.borderLight");
  border-radius: v-bind("theme.borderRadius.lg");
  text-decoration: none;
  transition: all 0.2s ease;
}

.post-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 10px 24px rgba(0, 0, 0, 0.08);
}

.post-header {
  display: flex;
  align-items: center;
  gap: v-bind("theme.spacing.sm");
  margin-bottom: v-bind("theme.spacing.sm");
}

.post-category {
  font-size: 11px;
  font-weight: 600;
  color: v-bind("theme.colors.primary");
  background-color: rgba(0, 123, 255, 0.1);
  padding: 4px 10px;
  border-radius: 999px;
  white-space: nowrap;
}

.post-title {
  font-size: 16px;
  font-weight: 700;
  color: v-bind("theme.colors.textPrimary");
  line-height: 1.5;
  flex: 1;
  margin: 0;
}

.post-preview {
  font-size: v-bind("theme.fontSize.sm");
  color: v-bind("theme.colors.textSecondary");
  line-height: 1.6;
  margin: 0;
}

.empty {
  text-align: center;
  padding: v-bind("theme.spacing.lg");
  color: v-bind("theme.colors.textSecondary");
}

.load-more-container {
  display: flex;
  justify-content: center;
  margin-top: v-bind("theme.spacing.xl");
}

.load-more-btn {
  padding: v-bind("theme.spacing.md") v-bind("theme.spacing.lg");
  background-color: v-bind("theme.colors.primary");
  color: white;
  border: none;
  border-radius: v-bind("theme.borderRadius.md");
  font-size: v-bind("theme.fontSize.base");
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
}

.load-more-btn:hover:not(:disabled) {
  background-color: #0056b3;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 123, 255, 0.3);
}

.load-more-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}
</style>
