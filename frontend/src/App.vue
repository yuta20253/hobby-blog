<script setup lang="ts">
import { ref, onMounted } from 'vue';
import axios from 'axios';

const status = ref("");

onMounted(async() => {
  try {
    const res = await axios.get('http://localhost:18080/health');

    status.value = res.data.status;
  } catch (error) {
    status.value = "error";
  }
});


</script>

<template>
  <div>
    <h1>Hobby Blog</h1>
    <p>API status: {{ status }}</p>
  </div>
</template>
