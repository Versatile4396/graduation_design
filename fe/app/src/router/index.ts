import {
  createRouter,
  createWebHistory,
  type RouteRecordRaw,
} from "vue-router";

export const routes: RouteRecordRaw[] = [
  {
    path: "/",
    name: "home",
    meta: {
      title: "互助首页",
    },
    component: () => import("@/page/home/index.vue"),
  },
  {
    path: "/login",
    name: "login",
    meta: {
      title: "互助登录页",
    },
    component: () => import("@/page/login/index.vue"),
  },
  {
    path: "/recommend",
    name: "recommend",
    meta: {
      title: "推荐页",
    },
    component: () => import("@/page/recommend/index.vue"),
  },
  {
    path: "/course",
    name: "course",
    meta: {
      title: "课程",
    },
    component: () => import("@/page/course/index.vue"),
  },
  {
    path: "/popular",
    name: "popular",
    meta: {
      title: "互助登录页",
    },
    component: () => import("@/page/popular/index.vue"),
  },
  {
    path: "/personal",
    name: "personal",
    meta: {
      title: "个人中心",
    },
    component: () => import("@/page/personal/index.vue"),
  },
  {
    path: "/test",
    name: "test",
    meta: {
      title: "测试界面",
    },
    component: () => import("@/page/test/index.vue"),
  },
];
const BASE = "forum";
const router = createRouter({
  history: createWebHistory(BASE),
  routes,
});
export default router;
