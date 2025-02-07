<template>
  <div class="container">
    <div class="title">
      <h3>发布文章</h3>
    </div>
    <div class="form">
      <RxForm
        :schema="publicSchema"
        :initial-values="{ CategoryId: 1 }"
      ></RxForm>
    </div>
  </div>
</template>

<script lang="ts" setup>
import RxForm from "@/page/components/RxForm/index.vue";
import { publicSchema } from "@/page/components/RxForm/schema/public";
</script>
<style scoped lang="scss">
.container {
  z-index: 10;
  .title {
    border-bottom: 1px solid #ccc;
    margin-bottom: 8px;
  }
}
</style>
