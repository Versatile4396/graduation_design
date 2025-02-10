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
import { computed, watch } from "vue";
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
    const query = getUrlQuery();
    // router.push({ name: routerName.PERSONAL, query })
  }
)
initApp();
</script>
<style scoped>
.occpancy {
  margin-top: 16px;
}
</style>
