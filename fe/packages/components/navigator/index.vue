<template>
  <div class="main-header-box">
    <div class="header"></div>
  </div>
</template>

<script lang="ts" setup></script>
<style scoped lang="scss">
/* @import url(); 引入css类 */
.main-header-box {
  position: relative;
  height: 60px;
  max-width: 1440px;
  z-index: 1002;
  background-color: rgb(255, 255, 255);
  .header {
    position: fixed;
  }
}
</style>
