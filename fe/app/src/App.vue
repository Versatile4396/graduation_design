<template>
  <Navigator v-if="navigatorStatus"></Navigator>
  <RouterView />
</template>
<script lang="ts" setup>
import { Navigator } from "@mono/components";
import { RouterView } from "vue-router";
import { computed } from "vue";

const navigatorStatus = computed(() => {
  return !location.href.includes("/login");
});
</script>
<style scoped></style>
