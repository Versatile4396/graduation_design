<template>
  <div class="container-box">
    <div class="aritcle-detail-container-box">
      <div class="operator-box">
        <div
          class="operator-icon"
          v-for="icon in operatorIconConfig"
          :key="icon.iconName"
          @click="icon.handleClick"
        >
          <div class="icon">
            <span v-if="icon?.content" class="icon-right-top" :style="{ '--color': icon.color }">{{
              icon?.content
            }}</span>
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
      <div class="editor-content">
        <div class="gpt-wrapper" v-if="aInfo?.GptSummarize">
          <div class="gpt-title">智能总结</div>
          <div class="gpt-summarize">
            {{ aInfo.GptSummarize }}
          </div>
        </div>
        <Editor v-model="aInfo.content" :defaultConfig="editorConfig" :mode="mode" />
      </div>
    </div>
    <div class="comments-container" id="comments-container">
      <div class="title">评论（{{ commentListCmp.length }}）</div>
      <div class="comments-input">
        <div class="avatar-box">
          <el-avatar class="avatar" :src="userInfo.avatar" />
        </div>
        <div class="input-box">
          <CommentInput :is-show="true" @comment="commentArticle"></CommentInput>
        </div>
      </div>

      <div class="comments-list">
        <Comment
          v-for="comment of commentListCmp"
          :key="comment.comment_id"
          :commentInfo="comment"
          @secondary-review="secondaryReview(comment)"
          @send-comment="commentArticle"
        />
      </div>
    </div>
    <div class="author-info-box">
      <div class="author-title">
        <el-avatar :src="authorInfo.avatar"></el-avatar>
        <div class="name">{{ authorInfo.username }}</div>
      </div>
      <div class="author-create-info">
        <div class="count-info">
          <span>{{ authorInfo.article_count ?? 0 }}</span>
          <span>文章</span>
        </div>
        <div class="count-info">
          <span>{{ authorInfo.like_count ?? 0 }}</span>

          <span>点赞</span>
        </div>
        <div class="count-info">
          <span>{{ authorInfo.view_count ?? 0 }}</span>
          <span>浏览</span>
        </div>
      </div>
      <div class="chat-author">
        <el-button type="primary" @click="goToChat">
          <template #default><span class="chat-btn">私聊</span></template>
        </el-button>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import Comment from './component/comment.vue'
import Ajax from '@/ajax'
import { getUrlQuery } from '@/utils/common'
import { computed, onMounted, ref } from 'vue'
import { Editor } from '@wangeditor/editor-for-vue'
import SvgIcon from '@/assets/iconfont/SvgIcon.vue'
import { userInfoStore } from '@/store/user'
import { type IArticle, type IAuthor, type IComment, type IArticleBrief } from './component/types'
import CommentInput from './component/comment-input.vue'
import { scrollToAnchor } from '@/utils/utils'
import router, { routerName } from '@/router'
import { Message } from '@/utils/message'

const { userInfo } = userInfoStore()

const editorConfig = {
  readOnly: true
} as any
const mode = 'default'
const aInfo = ref<Partial<IArticle>>({})
const authorInfo = ref<Partial<IAuthor>>({})
const aBrief = ref<IArticleBrief>()

const operatorIconConfig = computed(() => [
  {
    iconName: 'icon-icon',
    size: '34px',
    handleClick: async () => {
      await articleLike()
    },
    color: isLiked.value ? 'rgb(79, 125, 239)' : 'rgb(195, 200, 208)',
    content: LikeCount.value
  },
  {
    iconName: 'icon-pinglun',
    handleClick: () => {
      scrollToAnchor('comments-container')
    },
    content: Number(commentList.value?.length || 0),
    color: 'rgb(195, 200, 208)'
  },
  {
    iconName: 'icon-htmal5icon24',
    handleClick: async () => {
      await articleCollect()
    },
    content: collectCount.value,
    color: isCollect.value ? 'rgb(79, 125, 239)' : 'rgb(195, 200, 208)'
  },
  {
    iconName: 'icon-zhuanfa',
    handleClick: () => {},
    color: 'rgb(195, 200, 208)'
  },
  {
    iconName: 'icon-Frame-5',
    handleClick: () => {},
    color: 'rgb(195, 200, 208)'
  }
])

const LikeCount = computed(() => {
  return Number(aBrief.value?.like_count)
})
const collectCount = computed(() => {
  return Number(aBrief.value?.collect_count)
})

const isLiked = ref(false)
const isCollect = ref(false)
const commentList = ref<IComment[]>([])
// 处理二级评论内容
const commentListCmp = computed(() => {
  const firstComments = commentList.value.filter((item) => {
    return !item.parent_comment_id
  })
  const secondComments = commentList.value.filter((item) => {
    return item.parent_comment_id
  })
  const result = firstComments.map((item) => {
    const secondComment = secondComments.filter((second) => {
      return second.parent_comment_id === item.comment_id
    })
    return {
      ...item,
      second_comments: secondComment
    }
  })
  return result
})
// 跳转私聊界面
const goToChat = () => {
  const { uid } = getUrlQuery()
  console.log(authorInfo.value.user_id, 'uiuser_iduser_idd')
  const toUid = authorInfo.value?.user_id
  if (!uid) {
    Message.info('请先登录')
    return
  }
  if (String(uid) === String(toUid)) {
    Message.info('不能和自己私聊')
    return
  }
  router.push({ name: routerName.CHAT, query: { uid, toUid } })
}

const articleLike = async () => {
  const { uid, aid } = getUrlQuery()
  isLiked.value = !isLiked.value
  await Ajax.post('/article/like', {
    user_id: Number(uid),
    article_id: Number(aid),
    is_like: isLiked.value
  })
  // 获取文章brief 内容
  await getArticleBrief()
}

const articleCollect = async () => {
  const { uid, aid } = getUrlQuery()
  isCollect.value = !isCollect.value
  await Ajax.post('/article/collect', {
    user_id: Number(uid),
    article_id: Number(aid),
    is_collection: isCollect.value
  })
  // 获取文章brief 内容
  await getArticleBrief()
}

const getArticleBrief = async () => {
  const { aid } = getUrlQuery()
  const res = await Ajax.get('/article/getById/' + String(aid))
  console.log(res, 'resres')
  aBrief.value = res.data.article_brief
}

const estimateReadTime = (content: string) => {
  const words = content.split(' ').length
  const minutes = Math.ceil(words / 200)
  return `阅读 ${minutes} min`
}
// 拿取文章内容
const getAContent = async () => {
  const { aid } = getUrlQuery()
  const res = await Ajax.get('/article/getById/' + String(aid))
  aInfo.value = res.data.article_info
  authorInfo.value = res.data.author_info
  aInfo.value.create_at = aInfo.value.create_at?.split('T')[0]
  aBrief.value = res.data.article_brief
}
// 查询是否点赞过
const getLikeAndCollectStatus = async () => {
  const { uid, aid } = getUrlQuery()
  const likeRes = await Ajax.post('/article/isLiked', {
    user_id: Number(uid),
    article_id: Number(aid)
  })

  if (likeRes?.data) {
    operatorIconConfig.value[0].color = '#3f7ef7'
    isLiked.value = likeRes?.data
  }
  const collectRes = await Ajax.post('/article/isCollected', {
    user_id: Number(uid),
    article_id: Number(aid)
  })

  if (collectRes?.data) {
    operatorIconConfig.value[0].color = '#3f7ef7'
    isCollect.value = collectRes?.data
  }
}
// 发送评论
const commentArticle = async (commentText: string, commentId: number | string) => {
  const { uid, aid } = getUrlQuery()
  await Ajax.post('/comment/create', {
    user_id: Number(uid),
    article_id: Number(aid),
    content: commentText,
    parent_comment_id: commentId ?? null
  })
  // 重新拉取评论
  await getCommentList()
}

// 二级评论
const secondaryReview = async (comment: IComment) => {
  commentList.value.forEach((item) => {
    item.second_comments_status = false
    if (item.comment_id === comment.comment_id) {
      item.second_comments_status = true
    }
  })
}

const getCommentList = async () => {
  const { aid } = getUrlQuery()
  const res = await Ajax.post('/comment/list', {
    article_id: Number(aid),
    pagination: {
      page: 1,
      pagesize: 100
    }
  })
  commentList.value = res.data || []
}
const getAuthorCount = async () => {
  const { data } = await Ajax.post('/user/count/info/' + authorInfo.value.user_id, {})
  authorInfo.value.view_count = data?.view_count
  authorInfo.value.like_count = data?.like_count
  authorInfo.value.article_count = data?.article_count
}
onMounted(async () => {
  await getAContent()
  await getLikeAndCollectStatus()
  await getCommentList()
  await getAuthorCount()
})
</script>
<style scoped lang="scss">
.container-box {
  margin: 0 auto;
  width: 820px;

  .operator-box {
    position: fixed;
    margin-left: -7rem;
    top: 140px;
    z-index: 2;
    display: flex;
    align-items: center;
    justify-self: center;
    flex-direction: column;

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
        position: relative;

        .icon-right-top {
          top: -2px;
          font-size: 12px;
          height: 16px;
          width: 20px;
          line-height: 18px;
          right: -5px;
          border-radius: 12px;
          text-align: center;
          position: absolute;
          color: #fff;
          background-color: var(--color);
        }
      }

      .icon:hover {
        background-color: #f2f3f5;
      }
    }
  }

  .author-info-box {
    background-color: #fff;
    position: fixed;
    margin-left: 53rem;
    top: 76px;
    width: 300px;
    padding: 16px;
    z-index: 2;
    border-radius: 4px;

    .author-title {
      display: flex;
      gap: 16px;
      align-items: center;

      .name {
        font-size: 16px;
        font-weight: 600;
        height: 2rem;
        line-height: 2rem;
        text-align: center;
      }
    }

    .author-create-info {
      display: grid;
      grid-template-columns: 1fr 1fr 1fr;
      gap: 16px;
      margin-top: 16px;

      .count-info {
        color: #8b919e;
        font-size: 12px;
        font-weight: 400;
        line-height: 22px;
        display: flex;
        align-items: center;
        justify-content: center;
        flex-direction: column;
      }

      .count-info span:first-of-type {
        font-size: 16px;
        font-weight: 600;
        line-height: 22px;
        color: #262932;
      }
    }

    .chat-author {
      display: grid;
      margin-top: 16px;
      width: 100%;
    }
  }

  .aritcle-detail-container-box {
    border-radius: 4px 4px 0 0;
    padding: 2.667rem 2.667rem 0 2.667rem;
    background-color: #fff;
    border-radius: 4px;

    .editor-content {
      .gpt-wrapper {
        background-color: #f7f8fa;
        padding: 12px;
        border-radius: 12px;

        .gpt-title {
          padding: 8px;
          line-height: 24px;
          border-bottom: 1px solid #e5e6eb;
          color: #b0b0b0;
          font-size: 16px;
          font-style: normal;
          font-weight: 500;
          line-height: 16px;
          background: radial-gradient(
            495.98% 195.09% at 144.79%,
            at 10.71%,
            #ff8a01 0,
            #b051b9 22.37%,
            #672bff 45.54%,
            #06f 99.99%
          );
          background: radial-gradient(
            495.98% 195.09% at 144.79% 10.71%,
            #ff8a01 0,
            #b051b9 22.37%,
            #672bff 45.54%,
            #06f 99.99%
          );
          background-clip: text;
          -webkit-background-clip: text;
          -webkit-text-fill-color: transparent;
          margin-left: 4px;
          margin-right: 8px;
        }

        .gpt-summarize {
          padding: 8px 0;
          text-indent: 2em;
          font-family: 'PingFang SC';
          font-size: 14px;
          font-weight: 400;
          line-height: 24px;
          text-align: left;
          color: #252933;
          max-height: 200px;
        }
      }
    }

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
    min-height: 200px;
    padding: 20px 32px;
    border-radius: 4px;

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
