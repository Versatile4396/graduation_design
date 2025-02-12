<template>
  <div class="main-personal-box">
    <div class="left-box">
      <div class="header-box">
        <div class="avatar"></div>
        <div class="right-box">
          <div class="briefly"></div>
          <div class="editor-content"></div>
        </div>
      </div>
      <div class="content-box">
        <div class="content-header">
          <el-tabs v-model="activeName" @tab-click="handleClick">
            <el-tab-pane label="动态" name="dynamic"></el-tab-pane>
            <el-tab-pane label="文章" name="article">
              <Article></Article>
            </el-tab-pane>
            <el-tab-pane label="沸点" name="popular"></el-tab-pane>
            <el-tab-pane label="收藏" name="collection"></el-tab-pane>
            <el-tab-pane label="关注" name="follow"></el-tab-pane>
            <el-tab-pane label="赞1" name="like"></el-tab-pane>
          </el-tabs>
        </div>
      </div>
    </div>
    <div class="right-box"></div>
  </div>
</template>

<script lang="ts" setup>
import { userInfoStore } from "@/store/user";
import { storeToRefs } from "pinia";
import Article from "./tabsview/article.vue";
import { ref } from "vue";
import type { TabsPaneContext } from 'element-plus'

const activeName = ref('article')

const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event)
}
</script>
<style scoped lang="scss">
.main-personal-box {
  width: 960px;
  height: 100px;
  margin: 0 auto;
  display: flex;
  flex-direction: row;

  .left-box {
    width: 704px;
    margin-right: 16px;

    .header-box {
      background-color: #fff;
      display: flex;
      padding: 2.5rem;

      .avatar {
        width: 90px;
        height: 90px;
        border-radius: 50%;
        background-color: #ccc;
        margin-right: 24px;
      }

      .right-box {
        display: flex;
        flex-direction: column;
        justify-content: space-between;

        .briefly {
          font-size: 18px;
          font-weight: 500;
          color: #1d2129;
          margin-bottom: 8px;
        }

        .editor-content {
          font-size: 14px;
          color: #909090;
        }
      }
    }

    .content-box {
      margin-top: 16px;
      background-color: #fff;

      .content-header {
        padding: 8px;

      }
    }
  }

  .right-box {
    width: 240px;
    height: 100px;
    background-color: #ccc;
  }

}
</style>
