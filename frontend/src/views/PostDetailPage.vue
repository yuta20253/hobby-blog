<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { theme } from '../styles/theme';
import { useFetchPost } from '../service/post/postService';
import type { Post } from '../types/post';
import { useRoute } from 'vue-router';

const post = ref<Post | null>(null);
const route = useRoute();

const postId = route.params.id;

onMounted(async() => {
    post.value = await useFetchPost(Number(postId));

    console.log(post)
});
</script>

<template>
    <div>
        <h2>投稿詳細</h2>
    </div>
</template>