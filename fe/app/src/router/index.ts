import {
  createRouter,
  createWebHistory,
  type RouteRecordRaw,
} from "vue-router";

export enum routerName {
  HOME = "home",
  LOGIN = "login",
  RECOMMEND = "recommend",
  COURSE = "course",
  POPULAR = "popular",
  PERSONAL = "personal",
  CREATE_ARTICLE = "editor",
  TEST = "test",
}

export const routes: RouteRecordRaw[] = [
  {
    path: "/",
    name: routerName.HOME,
    meta: {
      title: "互助首页",
    },
    component: () => import("@/page/home/index.vue"),
  },
  {
    path: "/login",
    name: routerName.LOGIN,
    meta: {
      title: "互助登录页",
    },
    component: () => import("@/page/login/index.vue"),
  },
  {
    path: "/recommend",
    name: routerName.RECOMMEND,
    meta: {
      title: "推荐页",
    },
    component: () => import("@/page/recommend/index.vue"),
  },
  {
    path: "/course",
    name: routerName.COURSE,
    meta: {
      title: "课程",
    },
    component: () => import("@/page/course/index.vue"),
  },
  {
    path: "/popular",
    name: routerName.POPULAR,
    meta: {
      title: "互助登录页",
    },
    component: () => import("@/page/popular/index.vue"),
  },
  {
    path: "/personal",
    name: routerName.PERSONAL,
    meta: {
      title: "个人中心",
    },
    component: () => import("@/page/personal/index.vue"),
  },
  {
    path: "/create-article",
    name: routerName.CREATE_ARTICLE,
    meta: {
      title: "创建文章",
    },
    component: () => import("@/page/create/index.vue"),
  },
  {
    path: "/article/:id",
    name: "article",
    meta: {
      title: "文章详情",
    },
    component: () => import("@/page/article/index.vue"),
  },
  {
    path: "/test",
    name: routerName.TEST,
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
