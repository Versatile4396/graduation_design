<template>
  <div class="main-container">
    <div class="index-nav">
      <div v-for="(item) in sideBars" class="sidebar-item" :class="{
        'active-sidebar-item': activeSideItem === item.value
      }" @click="sidebarItemclick(item.value)">
        <div class="icon">
          <svg-icon :iconName="item.iconName" :color="activeSideItem === item.value ? '#4f7def' : ''"></svg-icon>
        </div>
        <div class="text">{{ item.label }}</div>
      </div>
    </div>
    <div class="right-container">
      <div class="article-content">
        <div class="filter-bar">
          <el-tabs v-model="filterKey">
            <el-tab-pane label="推荐" name="recommend">
              <ArticlePreview @click="goToArticle(pf?.aid!)" v-for="pf in previewInfos" :previewInfo="pf"
                :key="pf?.aid"></ArticlePreview>
            </el-tab-pane>
            <el-tab-pane label="最新" name="lastest">
              <ArticlePreview @click="goToArticle(pf?.aid!)" v-for="pf in previewInfos" :previewInfo="pf"
                :key="pf?.aid">
              </ArticlePreview>
            </el-tab-pane>
          </el-tabs>
        </div>
      </div>
      <div class="irrelevant-container">
        <div class="irelevant-content">
          <div class="hello-box">
            <div class="title">{{ timePeriod }}</div>
            <div class="desc">点亮在社区的每一天~</div>
          </div>
        </div>
        <div class="hot-article">
          热门文章
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import ArticlePreview from '@/page/components/ArticlePreview/index.vue'
import { useCategorieStore } from '@/store/article';
import { storeToRefs } from 'pinia';
import { computed, onMounted, ref, watch } from 'vue';
import { ElTabPane, ElTabs } from 'element-plus';
import Ajax from '@/ajax';
import { goToArticle } from '@/utils/goto'
import { getUrlQuery } from '@/utils/common'


const { Categories } = storeToRefs(useCategorieStore());
const iconConfig = [
  "icon-zonghechaxun",
  "icon-bijibendiannao",
  "icon-wheat__easyic",
  "icon-jixiezulin",
  "icon-dianzixinxi",
  "icon-waiyu",
  "icon-jichuguanli",
  "icon-okr",
]
const sideBars = computed(() => {
  const temp = Categories.value?.['article_categories'].map((item: { value: number, label: string }, index: number) => {
    return {
      ...item,
      iconName: iconConfig[index + 1]
    }
  })
  temp?.unshift({
    value: 0,
    label: '综合',
    iconName: iconConfig[0]
  })
  return temp
})
const activeSideItem = ref(0)
const sidebarItemclick = async (v: number) => {
  activeSideItem.value = v
  // 重新请求
  await getPreviewInfos()
}
const timePeriod = computed(() => {
  const now = new Date()
  const hour = now.getHours()
  if (hour >= 0 && hour < 6) {
    return '凌晨好！'
  } else if (hour >= 6 && hour < 12) {
    return '上午好！'
  } else if (hour >= 12 && hour < 18) {
    return '下午好！'
  } else {
    return '晚上好！'
  }
})
const filterKey = ref('recommend')
watch(filterKey, async (nv) => {
  if (nv === 'recommend') {
    await getPreviewInfos()
  } else {
    await getPreviewInfos(false)
  }
})
onMounted(async () => {
  const query = getUrlQuery()
  if (query.keyword) {
    // 删除keyword
    delete query.keyword
  }
  await getPreviewInfos()
})
type PreviewInfo = InstanceType<typeof ArticlePreview>['$props']['previewInfo']

const previewInfos = ref<Array<PreviewInfo>>([])

const getPreviewInfos = async (order = true) => {
  const { data } = await Ajax.post('/article/list', {
    order_by_time: order,
    category_id: activeSideItem.value === 0 ? null : activeSideItem.value,
  })
  previewInfos.value = data?.map((item: any) => {
    return {
      aid: item.articles.article_id,
      likes: item.article_briefs.like_count,
      comments: item.article_briefs.comment_count,
      author: item.article_briefs.username,
      ...item.articles,
    }
  })
}


</script>
<style scoped lang="scss">
/* @import url(); 引入css类 */
.main-container {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;

  .index-nav {
    position: fixed;
    background-color: #fff;
    min-width: 180px;
    margin-right: 20px;
    padding: 8px;
    border-radius: 4px;

    .sidebar-item {
      padding: 12px 17px;
      color: $normal-font;
      display: flex;
      align-items: center;
      border-radius: 4px;

      .icon {
        height: 16px;
        width: 16px;
        margin-right: 12px;
        position: relative;
        top: -4px;
        color: #333 !important;
      }

      .text {
        cursor: default;
        font-weight: 600;
        margin-left: 8px;
      }
    }

    .active-sidebar-item {
      background-color: $active-bgc !important;
      color: $active-font;
    }

    .sidebar-item:hover {
      background-color: $hover-bgc;
      color: $hover-font;
    }


  }

  .right-container {
    margin-left: 215px;
    display: flex;

    .article-content {
      border-radius: 4px;
      background-color: #fff;
      min-width: 720px;
      padding: 8px 0;
      margin-right: 16px;

      .filter-bar {
        padding: 4px;
      }
    }

    .irrelevant-container {
      display: flex;
      flex-direction: column;
      row-gap: 12px;
      border-radius: 4px;
      min-width: 260px;

      .irelevant-content {


        .hello-box {
          font-weight: 600;

          .desc {
            color: $normal-font;
            font-size: 12px;
            font-weight: 400;
            line-height: 24px;
            margin-top: 2px;
          }
        }
      }

      .hot-article {}

    }

    .irrelevant-container>div {
      border-radius: 4px;
      background-color: #fff;
      padding: 12px;
    }
  }
}
</style>
