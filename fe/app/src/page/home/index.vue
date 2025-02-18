<template>
  <div class="main-container">
    <div class="index-nav">
      <div v-for="(item) in sideBars" class="sidebar-item" :class="{
        'active-sidebar-item': activeSideItem === item.value
      }" @click="sidebarItemclick(item.value)">
        <div class="icon">
          <svg-icon :iconName="item.iconName"></svg-icon>
        </div>
        <div class="text">{{ item.label }}</div>
      </div>
    </div>
    <div class="right-container">
      <div class="article-content">
        <div class="filter-bar">
          <el-tabs v-model="filterKey">
            <el-tab-pane label="推荐" name="recommend">
              <ArticlePreview v-for="pf in previewInfos" :previewInfo="pf"></ArticlePreview>
            </el-tab-pane>
            <el-tab-pane label="最新" name="lastest"></el-tab-pane>
          </el-tabs>
        </div>
      </div>
      <div class="irrelevant-container">
        <div class="irelevant-content">reaasdasd</div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import ArticlePreview from '@/page/components/ArticlePreview/index.vue'
import { useCategorieStore } from '@/store/article';
import { storeToRefs } from 'pinia';
import { computed, ref, watch } from 'vue';
import { ElTabPane, ElTabs } from 'element-plus';

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
const sidebarItemclick = (v: number) => {
  activeSideItem.value = v
}
const filterKey = ref('recommend')
watch(filterKey, (nv) => {
  console.log(nv)
})
type PreviewInfo = InstanceType<typeof ArticlePreview>['$props']['previewInfo']

const previewInfos = ref<Array<PreviewInfo>>([{
  aid: '1',
  title: '123',
  status: 1,
  cover: '123',
  abstract: '123',
  tags: ['123'],
  publishTime: '123',
},{
  aid: '1',
  title: '123',
  status: 1,
  cover: '123',
  abstract: '123',
  tags: ['123'],
  publishTime: '123',
},{
  aid: '1',
  title: '123',
  status: 1,
  cover: '123',
  abstract: '123',
  tags: ['123'],
  publishTime: '123',
}])

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
      padding: 8px;
      margin-right: 16px;

      .filter-bar {
        padding: 4px;
      }
    }

    .irrelevant-container {
      .irelevant-content {
        border-radius: 4px;
        background-color: #fff;
        padding: 8px;
      }

      min-width: 260px;
    }
  }
}
</style>
