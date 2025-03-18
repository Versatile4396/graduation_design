<template>
  <div class="index-container-box">
    <div class="filter-box">
      <!-- <div class="sort-box">
                <span>综合排序</span>
                <span>最新优先</span>
                <span>最热优先</span>
            </div> -->
      <div class="time-filter-box"></div>
    </div>
    <div class="article-list-box">
      <ArticlePreview
        @click="goToArticle(pf?.aid!)"
        v-for="pf in previewInfos"
        :previewInfo="pf"
        :key="pf?.aid"
      >
      </ArticlePreview>
    </div>
  </div>
</template>

<script lang="ts" setup>
import ArticlePreview from '@/page/components/ArticlePreview/index.vue'
import { computed, onMounted, ref } from 'vue'
import Ajax from '@/ajax'
import { goToArticle } from '@/utils/goto'
import { getUrlQuery } from '@/utils/common'

type PreviewInfo = InstanceType<typeof ArticlePreview>['$props']['previewInfo']

const previewInfos = ref<PreviewInfo[]>([])

const getSearchArticleList = async () => {
  const { keyword } = getUrlQuery()
  const res = await Ajax.post('/article/searchLike', {
    keyword: keyword
  })
  previewInfos.value = res?.data?.map((item: any) => {
    return {
      aid: item.articles.article_id,
      likes: item.article_briefs.like_count,
      comments: item.article_briefs.comment_count,
      author: item.article_briefs.username,
      ...item.articles
    }
  })
}

onMounted(async () => {
  await getSearchArticleList()
})
</script>
<style scoped lang="scss">
.index-container-box {
  width: 700px;
  margin: 0 auto;
  margin-top: 16px;
  border-radius: 4px;
  background-color: #fff;
  padding: 16px 0;

  .filter-box {
    display: flex;
    justify-content: space-between;
    height: 24px;
    padding: 0 12px;
    border-bottom: $fill-color-1 1px solid;
    line-height: 24px;

    .sort-box {
      display: grid;
      grid-template-columns: 1fr 1fr 1fr;
    }
  }

  .article-list-box {
    padding: 0 12px;
  }
}
</style>
