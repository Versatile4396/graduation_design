<template>
  <Navigator v-show="navigatorStatus"></Navigator>
  <div :class="{
    occpancy: navigatorStatus,
  }">
    <RouterView />
  </div>
</template>
<script lang="ts" setup>
import Navigator from "@/page/components/navigator/index.vue";
import { RouterView } from "vue-router";
import { computed, onMounted, watch } from "vue";
import router, { routerName } from "./router";
import { getUrlQuery, userLocalInfo } from "@/utils/common";
import { userInfoStore } from "./store/user";
const { setUserInfo } = userInfoStore()

const navigatorStatus = computed(() => {
  return ![routerName.LOGIN, routerName.CREATE_ARTICLE].includes(
    // @ts-ignore
    router.currentRoute.value.name
  );
});

const initApp = () => {
  // 监听路由变化
  router.beforeEach((to, _, next) => {
    // 将当前路由信息存储到 sessionStorage 中
    sessionStorage.setItem('currentRoute', to.fullPath);
    next();
  });
  // 判断是否已经登录
  if (userLocalInfo.value?.user_id) {
    setUserInfo(userLocalInfo.value);
    // 已经登录了的给浏览器query添加uid
    router.push({
      path: router.currentRoute.value.path,
      query: {
        uid: userLocalInfo.value?.user_id,
      },
    });
  }
};
watch(
  () => router.currentRoute.value.name,
  () => {
    // const query = getUrlQuery();
    // router.push({ name: routerName.Article , query })
  }
)
initApp();
onMounted(() => {
  // 从 sessionStorage 中获取之前保存的路由信息
  const currentRoute = sessionStorage.getItem('currentRoute');
  if (currentRoute) {
    // 导航到之前的路由
    router.push(currentRoute);
  }
})
</script>
<style scoped>
.occpancy {
  margin-top: 16px;
}
</style>
