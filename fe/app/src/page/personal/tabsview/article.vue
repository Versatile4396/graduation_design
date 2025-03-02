<template>
    <div class="tab-content-article">
        <section class="header">
        </section>
        <section class="content">
            <ArticlePreview @click="goToArticle(previewInfo?.aid!)" v-for="previewInfo in previewInfos"
                :previewInfo="previewInfo"></ArticlePreview>
        </section>
    </div>
</template>

<script lang='ts' setup>
import ArticlePreview from '@/page/components/ArticlePreview/index.vue'
import { ref } from 'vue';
import { goToArticle } from '@/utils/goto'

interface ArticleList {
    articles: any;
    article_briefs: any;
}

interface Props {
    articleList: ArticleList[];
}
const props = defineProps<Partial<Props>>()

const previewInfos = ref<PreviewInfo[]>([])
type PreviewInfo = InstanceType<typeof ArticlePreview>['$props']['previewInfo']

previewInfos.value = props.articleList?.map((item: any) => {
    return {
        aid: item.articles.article_id,
        likes: item.article_briefs.like_count,
        comments: item.article_briefs.comment_count,
        author: item.article_briefs.username,
        ...item.articles,
    }
}) as PreviewInfo[]

</script>
<style scoped lang="scss">
.tab-content-article {
    .header {
        margin-bottom: 12px;
    }
}
</style>