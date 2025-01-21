<template>
  <Navigator v-if="navigatorStatus"></Navigator>
  <RouterView />
</template>
<script lang="ts" setup>
import { Navigator } from "@mono/components";
import { RouterView } from "vue-router";
import { computed, onMounted } from "vue";
import router from "./router";
import { userLocalInfo } from "@/utils/common";

const navigatorStatus = computed(() => {
  return router.currentRoute.value.name !== "login";
});

const initApp = () => {
  // 判断是否已经登录
  if (userLocalInfo.value?.user_id) {
    // 已经登录了的给浏览器query添加uid
    router.push({
      path: router.currentRoute.value.path,
      query: {
        uid: userLocalInfo.value?.user_id,
      },
    });
  }
};
initApp();
</script>
<style scoped></style>
