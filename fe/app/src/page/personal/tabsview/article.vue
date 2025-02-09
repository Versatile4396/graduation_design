<template>
    <div class="tab-content-article">
        <section class="header">
            <el-button>最新</el-button>
            <el-button>热门</el-button>
        </section>
        <section class="content">
            <ArticlePreview :previewInfo="previewInfo"></ArticlePreview>
        </section>
    </div>
</template>

<script lang='ts' setup>
import ArticlePreview from '@/page/components/ArticlePreview/index.vue'
import { ArticleStatus } from '@/page/components/ArticlePreview/type';

type PreviewInfo = InstanceType<typeof ArticlePreview>['$props']['previewInfo']

const previewInfo: PreviewInfo = {
    status: ArticleStatus.Published,
    title: '这是一篇文章,这是一篇文章这是一篇文章这是一篇文章这是一篇文章这是一篇文章这是一篇文章这是一篇文章这是一篇文章这是一篇文章这是一篇文章这是一篇文章这是一篇文章',
    abstract: '这是一篇文章的摘要,这是一篇文章这是一篇文章这是一篇文章这是一篇文章这是一篇文章这是一篇文章这是一篇文章这是一篇文章这是一篇文章这是一篇文章这是一篇文章这是一篇文章',
    cover: 'https://p9-xtjj-sign.byteimg.com/tos-cn-i-73owjymdk6/2e1da62fee344b9c992f0cf39b4aa9eb~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAgVmVyU2F0MWxl:q75.awebp?rk3s=f64ab15b&x-expires=1739703589&x-signature=7ysp5c8CXjl2WKBZqAQ%2FUtfOric%3D',
    publishTime: '2021-01-01',
    tags: ['标签1', '标签2'],
    author: '作者',
    pageviews: 100,
    likes: 100,
    comments: 100,
}


</script>
<style scoped lang="scss">
.tab-content-article {
    .header {
        margin-bottom: 12px;
    }
}
</style>