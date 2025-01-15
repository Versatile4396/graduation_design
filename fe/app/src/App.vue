<template>
  <Navigator v-if="navigatorStatus"></Navigator>
  <RouterView />
</template>
<script lang="ts" setup>
import { Navigator } from "@mono/components";
import { RouterView } from "vue-router";
import { computed } from "vue";
import "@/assets/style/global.css";

const navigatorStatus = computed(() => {
  return !location.href.includes("/login");
});
</script>
<style scoped></style>
