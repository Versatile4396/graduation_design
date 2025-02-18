<template>
    <div class="aritcle-detail-container-box">
        <div class="article-title">
            {{ aInfo?.title }}
        </div>
        <div class="author-info-block">
            <span class="username">{{ authorInfo.username }}</span>
            <div class="base-info-article">
                <span class="create-time">{{ aInfo.create_at }}</span>
                <span class="page-view">
                    <svg-icon iconName="icon-liulan"></svg-icon>
                    {{ aInfo.view_count }}
                </span>
                <span class="reading-time">
                    <svg-icon iconName="icon-shijian" size="20px"></svg-icon>
                    {{ estimateReadTime(aInfo.content || '') }}
                </span>
                <span class="other-info"></span>
            </div>
        </div>
        <div class="content">
            <Editor v-model="aInfo.content" :defaultConfig="editorConfig" :mode="mode" />
        </div>
    </div>
</template>

<script lang='ts' setup>
import Ajax from '@/ajax';
import { getUrlQuery } from '@/utils/common';
import { onMounted, ref } from 'vue';
import { Editor } from "@wangeditor/editor-for-vue";
interface IArticle {
    article_id: number;
    content: string;
    title: string;
    user_id: string;
    category_id: number;
    topic_id: number;
    tag_id: number;
    cover: string;
    abstract: string;
    create_at: string;
    update_at: string;
    view_count: number;
}
interface IAuthor {
    avatar: string;
    email: string;
    gender: number;
    user_id: number;
    username: string;
}
const editorConfig = {
    readOnly: true,
} as any;
const mode = "default";
const aInfo = ref<Partial<IArticle>>({});
const authorInfo = ref<Partial<IAuthor>>({});
const estimateReadTime = (content: string) => {
    const words = content.split(' ').length;
    const minutes = Math.ceil(words / 200);
    return `${minutes} min`;
}
// 拿取文章内容
const getAContent = async () => {
    const { aid } = getUrlQuery();
    const res = await Ajax.get("/article/getById/" + String(aid));
    aInfo.value = res.data.article_info;
    authorInfo.value = res.data.author_info;
    aInfo.value.create_at = aInfo.value.create_at?.split('T')[0];
}
onMounted(async () => {
    await getAContent()
})
</script>
<style scoped lang="scss">
.aritcle-detail-container-box {
    margin: 0 auto;
    width: 820px;
    border-radius: 4px 4px 0 0;
    padding: 2.667rem 2.667rem 0 2.667rem;
    background-color: #fff;

    .article-title {
        margin: 0 0 1.3rem;
        font-size: 2.667rem;
        font-weight: 600;
        line-height: 1.31;
        color: $title-font;
    }

    .author-info-block {
        font-size: 14px;
        line-height: 1.667rem;
        display: flex;
        align-items: center;
        margin-bottom: 1.667rem;
        gap: 16px;

        .username {
            display: inline-block;
            vertical-align: top;
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
            color: #515767;
        }

        .base-info-article {
            display: flex;
            color: #8a919f;
            gap: 12px;

            .page-view {
                display: flex;
                gap: 4px;
                align-items: center;
            }

            .reading-time {
                display: flex;
                gap: 4px;
                align-items: center;
            }
        }
    }
}
</style>