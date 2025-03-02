<template>
  <div class="main-personal-box">
    <div class="left-box">
      <div class="header-box">
        <div class="avatar">
          <el-avatar :src="userInfo.avatar" :size="90"></el-avatar>
        </div>
        <div class="right-box">
          <div class="briefly">{{ userInfo.username }}</div>
          <div class="editor-content">
            <el-button type="primary" @click="setUserInfo">设置</el-button>
          </div>
        </div>
      </div>
      <div class="content-box">
        <div class="content-header">
          <el-tabs v-model="activeName" @tab-click="handleClick">
            <el-tab-pane label="个人文章" name="article">
              <Article :article-list="articleList" :key="articleList.join('')"></Article>
            </el-tab-pane>
            <el-tab-pane label="收藏" name="collection">
              <Article :article-list="articleList" :key="articleList.join('')"></Article>
            </el-tab-pane>
            <el-tab-pane label="赞" name="like">
              <Article :article-list="articleList" :key="articleList.join('')"></Article>
            </el-tab-pane>
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
import { onMounted, ref } from "vue";
import type { TabsPaneContext } from 'element-plus'
import { getUrlQuery } from "@/utils/common";
import Ajax from "@/ajax";

const { userInfo } = userInfoStore();

const activeName = ref('article')
const articleList = ref([])

const handleClick = async (tab: TabsPaneContext, event: Event) => {
  switch (tab.index) {
    case '0':
      await getArticlePersonalList();
      break;
    case '1':
      await getCollectionList();
      break;
    case '2':
      break;
    default:
      break;
  }
}
const getArticlePersonalList = async () => {
  const uid = getUrlQuery().uid;
  const res = await Ajax.post("/article/list", { user_id: uid })
  articleList.value = res.data
}

const getCollectionList = async () => {
  const uid = getUrlQuery().uid;
  const res = await Ajax.post("/article/list", { user_id: uid })
  articleList.value = res.data
}

const setUserInfo = () => {

}
onMounted(async () => {
  await getArticlePersonalList();
})

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
  }

}
</style>
