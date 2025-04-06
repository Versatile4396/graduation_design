<template>
  <div class="backend-container-box">
    <div class="tab-list">
      <el-menu
        default-active="1"
        class="el-menu-vertical"
        active-text-color="#ffd04b"
        background-color="#545c64"
        text-color="#fff"
        @select="handleClick"
      >
        <template #title>互助帮学后台管理系统</template>
        <el-menu-item index="1">
          <span>用户管理</span>
        </el-menu-item>
        <el-menu-item index="2">
          <span>文章管理</span>
        </el-menu-item>
        <el-menu-item index="3">
          <span>互助帮学</span>
        </el-menu-item>
        <el-menu-item index="4">
          <span>公告配置</span>
        </el-menu-item>
      </el-menu>
    </div>
    <div class="content-wrapper">
      <component :is="currentComponent" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue'
import UserPage from './components/user.vue'
import ArticlePage from './components/article.vue'
import AnnouncementPage from './components/announcement.vue'
import AssistancePage from './components/assistance.vue'

const pageComponents = {
  '1': UserPage,
  '2': ArticlePage,
  '3': AssistancePage,
  '4': AnnouncementPage
}


//@ts-ignore
const currentComponent = computed(() => pageComponents[currentMenu.value as string])

const currentMenu = ref('1')
const handleClick = (v: string) => {
  currentMenu.value = v
}
</script>
<style scoped lang="scss">
.backend-container-box {
  width: 100%;
  height: 100%;
  display: flex;
  .tab-list {
    height: 100%;
    display: flex;
    .el-menu-vertical {
      width: 200px;
    }
  }
  .content-wrapper {
    width: 100%;
  }
}
</style>
