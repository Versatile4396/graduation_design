<template>
  <div class="main-header-box">
    <div
      :class="{
        header: true,
        visible: navStatus,
      }"
    >
      <div class="header-content">
        <div class="logo" @click="handleLogoClick">
          <img src="../../../assets/image/study.svg" alt="" class="logo-img" />
          <span class="logo-text">跨知学域</span>
        </div>
        <div class="nav-content">
          <div class="nav-link">
            <div v-for="(node, index) in naviConfig">
              <div
                class="nav-node"
                :class="{
                  'active-node': index === activeNodeIndex,
                }"
                @click="handleNavClick(node.path)"
              >
                {{ node.label }}
              </div>
            </div>
          </div>
          <div class="avator"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import router from "@/router";
import { getUrlQuery } from "@/utils/common";
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
const handleLogoClick = () => {
  const query = getUrlQuery();
  router.push({ path: "/", query });
};
const handleNavClick = (path: string) => {
  const query = getUrlQuery();
  router.push({ path, query });
};
interface naviNode {
  path: string;
  name?: string;
  label: string;
}
const naviConfig: naviNode[] = [
  {
    path: "/",
    name: "home",
    label: "首页",
  },
  {
    path: "/recommend",
    name: "recommend",
    label: "推荐",
  },
  {
    path: "/course",
    name: "course",
    label: "课程",
  },

  {
    path: "/popular",
    label: "沸点",
  },
  {
    path: "/personal",
    name: "personal",
    label: "个人中心",
  },
];
const activeNodeIndex = ref(0);
//
router.afterEach((to) => {
  console.log(to);
  activeNodeIndex.value = naviConfig.findIndex((item) => item.path === to.path);
  console.log(activeNodeIndex.value);
});
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
  .header-content {
    margin: 0 auto;
    max-width: 1440px;
    height: 100%;
    display: flex;
    align-items: center;
    .logo {
      display: flex;
      align-items: center;
      font-size: 24px;
      margin-right: 24px;
      .logo-img {
        height: 50px;
        width: 50px;
      }
      .logo-text {
        font-family: cursive;
        cursor: pointer;
      }
    }
    .nav-content {
      display: flex;
      justify-content: space-between;
      align-items: center;
      .nav-link {
        display: flex;
        .nav-node {
          color: rgb(82, 87, 102);
          min-width: 52px;
          cursor: pointer;
        }
        .active-node {
          color: rgb(63, 126, 247);
        }
      }
    }
  }
}
</style>
