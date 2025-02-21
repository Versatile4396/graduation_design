<template>
    <div class="container-box">
        <div class="aritcle-detail-container-box">
            <div class="operator-box">
                <div class="operator-icon" v-for="icon in operatorIconConfig" :key="icon.iconName"
                    @click="icon.handleClick">
                    <div class="icon">
                        <svg-icon :iconName="icon.iconName" size="30px" :color="icon.color ?? ''"></svg-icon>
                    </div>
                </div>
            </div>
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
        <div class="comments-container">
            <div class="title">评论（{{ commentList.length }}）</div>
            <div class="comments-input">
                <div class="avatar-box">
                    <el-avatar class="avatar" :src="userInfo.avatar" />
                </div>
                <div class="input-box">
                    <el-input v-model="commentText" resize="none" :rows="4" type="textarea" placeholder="友善评论，平等表达">
                    </el-input>
                </div>
            </div>
            <div class="submit-btn">
                <el-button width="120px" type="primary" @click="commentArticle">发送</el-button>
            </div>
            <div class="comments-list"></div>
        </div>
        <div class="right-box"></div>
    </div>
</template>

<script lang='ts' setup>
import Ajax from '@/ajax';
import { getUrlQuery } from '@/utils/common';
import { computed, onMounted, ref } from 'vue';
import { Editor } from "@wangeditor/editor-for-vue";
import SvgIcon from '@/assets/iconfont/SvgIcon.vue';
import { Message } from '@/utils/message';
import { userInfoStore } from "@/store/user";
const { userInfo } = userInfoStore();

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
const aBrief = ref()
const operatorIconConfig = computed(() => [
    {
        iconName: "icon-icon",
        size: "34px",
        handleClick: () => {
            articleLike();
        },
        color: isLiked.value ? '#3f7ef7' : ""
    }, {
        iconName: "icon-pinglun",
        handleClick: () => {
        }
    }, {
        iconName: "icon-zhuanfa",
        handleClick: () => {
        }
    }, {
        iconName: "icon-Frame-5",
        handleClick: () => {
        }
    }
])
const isLiked = ref(false)
const commentText = ref('')
const commentList = ref([])

const articleLike = () => {
    const { uid, aid } = getUrlQuery();
    isLiked.value = !isLiked.value
    Ajax.post('/article/like', {
        user_id: Number(uid),
        article_id: Number(aid),
        is_like: isLiked.value
    }).then((res) => {
        if (res.code === 0 && isLiked.value) {
            Message.info('点赞成功')
        }
    })
}

const estimateReadTime = (content: string) => {
    const words = content.split(' ').length;
    const minutes = Math.ceil(words / 200);
    return `阅读 ${minutes} min`;
}
// 拿取文章内容
const getAContent = async () => {
    const { aid } = getUrlQuery();
    const res = await Ajax.get("/article/getById/" + String(aid));
    aInfo.value = res.data.article_info;
    authorInfo.value = res.data.author_info;
    aInfo.value.create_at = aInfo.value.create_at?.split('T')[0];
    aBrief.value = res.data.article_brief;
    console.log(aBrief.value, 'aBrief')
}
// 查询是否点赞过
const getLikeStatus = async () => {
    const { uid, aid } = getUrlQuery();
    const res = await Ajax.post('/article/isLiked', {
        user_id: Number(uid),
        article_id: Number(aid),
    })

    if (res?.data) {
        operatorIconConfig.value[0].color = '#3f7ef7'
        isLiked.value = res?.data
        console.log(operatorIconConfig, 'operatorIconConfig')
    }
}
// 发送评论
const commentArticle = async () => {
    const { uid, aid } = getUrlQuery();
    await Ajax.post('/comment/create', {
        user_id: Number(uid),
        article_id: Number(aid),
        content: commentText.value
    })
    commentText.value = ''
    // 重新拉取评论
    await getCommentList()
}
const getCommentList = async () => {
    const { aid } = getUrlQuery();
    const res = await Ajax.post('/comment/list', {
        article_id: Number(aid)
    })
    commentList.value = res.data || []
}
onMounted(async () => {
    await getAContent()
    await getLikeStatus()
    await getCommentList()
})
</script>
<style scoped lang="scss">
.container-box {
    margin: 0 auto;
    width: 820px;

    .operator-box {
        display: flex;
        align-items: center;
        justify-self: center;
        flex-direction: column;
        width: 60px;
        position: fixed;
        margin-left: -7rem;
        top: 140px;
        z-index: 2;
        height: 100px;

        .operator-icon {
            .icon {
                cursor: pointer;
                margin: 0 auto;
                display: flex;
                justify-content: center;
                align-items: center;
                width: 50px;
                height: 50px;
                border-radius: 50%;
                background-color: #fff;
                margin-bottom: 1.667rem;
            }

            .icon:first-of-type::before {
                content: "";
            }

            .icon:hover {
                background-color: #f2f3f5;
            }
        }
    }

    .aritcle-detail-container-box {
        border-radius: 4px 4px 0 0;
        padding: 2.667rem 2.667rem 0 2.667rem;
        background-color: #fff;
        border-radius: 4px;
        display: none;

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

    .comments-container {
        background-color: #fff;
        margin-top: 20px;
        padding: 20px 32px;

        .title {
            font-size: 18px;
            font-weight: 600;
            line-height: 30px;
        }

        .comments-input {
            display: flex;
            margin-top: 20px;
            gap: 12px;

            .avatar-box {
                width: 40px;
                height: 40px;
                border-radius: 50%;
                background-color: #f2f3f5;
            }

            .input-box {
                width: 100%;
                height: 40px;
                background-color: #f2f3f5;
                border-radius: 4px;
                height: 92px;
            }
        }

        .submit-btn {
            margin-top: 12px;
            display: flex;
            flex-direction: row-reverse;
        }
    }
}
</style>