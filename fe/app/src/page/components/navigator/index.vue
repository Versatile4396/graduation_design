<template>
  <div class="main-header-box">
    <div
      :class="{
        header: true,
        visible: navStatus,
      }"
    ></div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
const navStatus = ref(true);
var isScrolling = false;
const scrollCallback = () => {
  if (!isScrolling) {
    isScrolling = true;
    window.requestAnimationFrame(() => {
      const scrollTop =
        document.documentElement.scrollTop || document.body.scrollTop;
      if (scrollTop > 400 && navStatus.value) {
        navStatus.value = false;
      } else if (scrollTop <= 200 && !navStatus.value) {
        navStatus.value = true;
      }
      isScrolling = false;
    });
  }
};

window.addEventListener("scroll", scrollCallback);
// 导航头逻辑 滚动超过550px 隐藏
</script>
<style scoped lang="scss">
/* @import url(); 引入css类 */
.main-header-box {
  position: relative;
  height: 60px;
  max-width: 1440px;
  z-index: 1002;
  .header {
    height: 60px;
    position: fixed;
    background-color: #fff;
    top: 0;
    left: 0;
    right: 0;
    transition: transform 0.2s;
    display: block;
    transform: translate3d(0, -100%, 0);
    box-shadow: 0 2px 8px #d8dbdd54;
    color: #909090;
    z-index: 250;
  }
  .visible {
    transform: translateZ(0);
  }
}
</style>
