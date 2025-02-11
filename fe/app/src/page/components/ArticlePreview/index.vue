<template>
    <div class="article-container-box">
        <div class="detail-info">
            <div class="title-row">
                <div class="title-text">
                    <el-tag class="status" v-if="previewInfo?.status" size="small" :type="statusMap?.type">
                        {{ statusMap?.text }}
                    </el-tag>
                    {{ previewInfo?.title }}
                </div>
            </div>
            <div class="content-info">
                <div class="abstract">{{ previewInfo?.abstract }}</div>
                <div class="base-info">
                    <div class="like-info">
                        <span>{{ publicText }}</span>
                        <span class="split">|</span>
                        <span class="icon-text">
                            <el-icon class="icon-style">
                                <View />
                            </el-icon>
                            <span>{{ previewInfo?.pageviews }}</span>
                        </span>
                        <span class="icon-text">
                            <el-icon class="icon-style">
                                <Like></Like>
                            </el-icon>
                            <span>{{ previewInfo?.likes }}</span>
                        </span>
                        <span class="icon-text">
                            <el-icon class="icon-style">
                                <Comment />
                            </el-icon>
                            <span>{{ previewInfo?.comments }}</span>
                        </span>
                    </div>
                    <div class="tags">
                        <el-tag v-for="item in previewInfo?.tags" type="info" size="small">{{ item }}</el-tag>
                    </div>
                </div>
            </div>
        </div>
        <img v-if="previewInfo?.cover" class="cover" :src="previewInfo?.cover" alt="">
    </div>
</template>

<script lang='ts' setup>
import { computed } from 'vue';
import { ArticleStatus } from './type';
import { View, Comment } from '@element-plus/icons-vue'
import Like from '@/assets/iconfont/like.vue'


interface Props {
    previewInfo: Partial<{
        status: ArticleStatus;
        title: string;
        abstract: string;
        cover: string;
        publishTime: string;
        tags: string[];
        author: string;
        pageviews: number;
        likes: number;
        comments: number;
        collects: number;
    }>
}
const statusMap = computed((): any => {
    if (props.previewInfo?.status === ArticleStatus.Published) {
        return {
            text: '已发布',
            type: 'success'
        }
    } else if (props.previewInfo?.status === ArticleStatus.Draft) {
        return {
            text: '草稿',
            type: 'info'
        }
    } else if (props.previewInfo?.status === ArticleStatus.Deleted) {
        return {
            text: '已删除',
            type: 'danger'
        }
    } else {
        return {
            text: '未知',
            type: 'info'
        }
    }
})
const publicText = computed(() => {
    if (props.previewInfo?.author) {
        return props.previewInfo?.author
    } else {
        return props.previewInfo?.publishTime
    }
})

const props = defineProps<Partial<Props>>()


</script>
<style scoped lang="scss">
.article-container-box {
    height: 74px;
    padding: 12px 12px 0;
    cursor: pointer;
    display: flex;

    .detail-info {
        flex: 1 1 auto;
        display: flex;
        flex-direction: column;
        width: 534px;

        .title-row {
            display: flex;
            margin-bottom: 2px;
            align-items: center;
            width: 100%;

            .title-text {
                font-weight: 600;
                font-size: 16px;
                line-height: 24px;
                width: 100%;
                white-space: nowrap;
                overflow: clip;
                text-overflow: ellipsis;
                -webkit-box-orient: vertical;

                .status {
                    position: relative;
                    transform: translateY(-8%);
                    margin-right: 4px;
                }
            }
        }

        .content-info {
            display: flex;
            flex-direction: column;
            width: 100%;

            .abstract {
                font-size: 14px;
                line-height: 22px;
                color: #606370;
                width: 100%;
                white-space: nowrap;
                overflow: hidden;
                text-overflow: ellipsis;
                -webkit-box-orient: vertical;
            }

            .base-info {
                display: flex;
                justify-content: space-between;

                .like-info {
                    display: flex;
                    align-items: center;
                    margin-top: 4px;
                    font-size: 14px;
                    color: #606370;

                    .split {
                        padding: 0 12px;
                    }

                    .icon-text {
                        display: flex;
                        align-items: center;

                        .icon-style {
                            margin-right: 4px;
                        }

                        margin-right: 8px;
                    }

                    .icon-text:hover {
                        color: rgb(88, 136, 234);
                    }
                }

                .tags {
                    display: flex;
                    gap: 8px;
                }
            }

        }
    }

    .cover {
        flex: 0 0 auto;
        width: 108px;
        height: 72px;
        margin-left: 24px;
        border-radius: 4px;
        object-fit: cover;
    }
}

.article-container-box:hover {
    background-color: #f7f8fa;
}
</style>