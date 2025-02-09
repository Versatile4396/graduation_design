<template>
    <div class="tab-content-article">
        <section class="header">
            <el-button>最新</el-button>
            <el-button>热门</el-button>
        </section>
        <section class="content">

            <ArticlePreview v-for="previewInfo in previewInfos" :previewInfo="previewInfo"></ArticlePreview>
        </section>
    </div>
</template>

<script lang='ts' setup>
import Ajax from '@/ajax';
import ArticlePreview from '@/page/components/ArticlePreview/index.vue'
import { ArticleStatus } from '@/page/components/ArticlePreview/type';
import { getUrlQuery } from '@/utils/common';
import { onMounted, ref } from 'vue';

type PreviewInfo = InstanceType<typeof ArticlePreview>['$props']['previewInfo']

const previewInfos = ref<PreviewInfo[]>([])
onMounted(async () => {
    const uid = getUrlQuery().uid;
    const res = await Ajax.post("/article/list", { user_id: uid })
    previewInfos.value = res.data.map((item: any) => {
        return {
            title: item.title,
            articleId: item.article_id,
            status: item.status === 1 ? ArticleStatus.Published : ArticleStatus.Draft,
            cover: item.cover,
            abstract: item.abstract,
            tags: item.tag_id == 1 ? ['标签1'] : [],
            publishTime: "string",
            author: "作者",
            pageviews: 100,
            likes: 100,
            comments: 100,
            collects: 100,
        }
    })
})

</script>
<style scoped lang="scss">
.tab-content-article {
    .header {
        margin-bottom: 12px;
    }
}
</style>