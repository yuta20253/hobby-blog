import { createRouter, createWebHistory } from "vue-router";
import HomePage from "../views/HomePage.vue";
import PostsPage from "../views/PostsPage.vue";
import LoginPage from "../views/LoginPage.vue";
import SignupPage from "../views/SignupPage.vue";
import MyPage from "../views/MyPage.vue";
import PostDetailPage from "../views/PostDetailPage.vue";

const routes = [
  { path: "/", name: "Home", component: HomePage },
  { path: "/posts", name: "Posts", component: PostsPage },
  { path: "/posts/:id", name: "Post", component: PostDetailPage },
  { path: "/login", name: "Login", component: LoginPage },
  { path: "/signup", name: "Signup", component: SignupPage },
  { path: "/mypage", name: "MyPage", component: MyPage },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
