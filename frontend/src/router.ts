import { createRouter, createWebHistory } from "vue-router";
import Home from "./views/Home.vue";
import Root from "./views/Root.vue";
import Login from "./views/Login.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", name: "root", component: Root },
    { path: "/login", name: "login", component: Login },
    { path: "/home", name: "home", component: Home },
    { path: "/setup", component: () => import("./views/Setup.vue") },
  ],
});

export default router;
