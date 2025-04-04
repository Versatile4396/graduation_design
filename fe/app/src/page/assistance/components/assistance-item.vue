<template>
  <div class="assistance-item-container-box">
    <div class="detail-info">
      <div class="userinfo">
        <el-popover>
          <template #reference>
            <el-avatar :src="assistance.user_info.avatar" />
          </template>
        </el-popover>
        <div class="overview-wrapper">
          <div class="name">{{ assistance.user_info.nickname }}</div>
          <div class="overview">
            <span>{{ assistance.user_info.overview }}</span>
            <span style="margin: 0 4px; font-size: 12px">|</span>
            <span class="time">{{ createAt }}</span>
          </div>
        </div>
      </div>
      <div class="content" v-html="assistance.content"></div>
    </div>
    <div class="option-wrapper">
      <div class="operate">
        <div class="share"><svg-icon iconName="icon-zhuanfa" color="#4F4F4F"></svg-icon>分享</div>
        <div class="comment" @click="showComment">
          <svg-icon iconName="icon-pinglun" color="#4F4F4F" size="1.1rem" top="2px"></svg-icon
          ><span>评论</span>
        </div>
        <div class="other"><svg-icon iconName="icon-liulan" color="#4F4F4F"></svg-icon></div>
      </div>
      <div class="comment-wrapper" v-if="commentStatus">
        <CommentInput :is-show="true" @comment="commentAssistance"></CommentInput>
        <div class="comments-list">
          <Comment
            v-for="comment of commentListCmp"
            :key="comment.comment_id"
            :commentInfo="comment"
            @secondary-review="secondaryReview(comment)"
            @send-comment="commentAssistance"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { type Assistance } from '@/ajax/type'
import { computed, onMounted, ref } from 'vue'
import SvgIcon from '@/assets/iconfont/SvgIcon.vue'
import CommentInput from '@/page/article/component/comment-input.vue'
import Comment from '@/page/article/component/comment.vue'

import Ajax from '@/ajax'
import { getUrlQuery } from '@/utils/common'

interface Props {
  assistance: Assistance
}
const props = withDefaults(defineProps<Props>(), {})
const createAt = computed(() => {
  return props.assistance.create_at.slice(0, 10)
})
const commentStatus = ref(false)
const showComment = () => {
  commentStatus.value = !commentStatus.value
}
const commentAssistance = async (value: string, commentId: number) => {
  const user_id = Number(getUrlQuery().uid)
  await Ajax.post('/assistanceComment/create', {
    user_id,
    content: value,
    assistance_id: props.assistance.assistance_id,
    parent_comment_id: commentId || undefined
  })
  await getAssitanceComment()
}

const commentList = ref()
const commentListCmp = computed(() => {
  const firstComments = commentList.value?.filter((item: any) => {
    return !item.parent_comment_id
  })
  const secondComments = commentList.value?.filter((item: any) => {
    return item.parent_comment_id
  })
  const result = firstComments?.map((item: any) => {
    const secondComment = secondComments?.filter((second: any) => {
      return second.parent_comment_id === item.comment_id
    })
    return {
      ...item,
      second_comments: secondComment
    }
  })
  return result
})

const getAssitanceComment = async () => {
  const res = await Ajax.post('/assistanceComment/list', {
    assistance_id: props.assistance.assistance_id
  })
  commentList.value = res?.data
  const firstComments = commentList.value.filter((item: any) => {
    return !item.parent_comment_id
  })
  const secondComments = commentList.value.filter((item: any) => {
    return item.parent_comment_id
  })
  const result = firstComments.map((item: any) => {
    const secondComment = secondComments.filter((second: any) => {
      return second.parent_comment_id === item.comment_id
    })
    return {
      ...item,
      second_comments: secondComment
    }
  })
  console.log(result, firstComments, secondComments, 'commentListasdasd')
}

// 二级评论
const secondaryReview = async (comment: any) => {
  commentList.value.forEach((item: any) => {
    item.second_comments_status = false
    if (item.comment_id === comment.comment_id) {
      item.second_comments_status = true
    }
  })
}
onMounted(async () => {
  await getAssitanceComment()
})
</script>
<style scoped lang="scss">
.assistance-item-container-box {
  padding: 12px 12px 0 12px;
  border-radius: 4px;
  margin-bottom: 12px;
  background-color: #fff;
  .detail-info {
    display: flex;
    flex-direction: column;
    .userinfo {
      display: flex;
      align-items: center;
      .overview-wrapper {
        margin-left: 8px;
        font-weight: 500;
        font-size: 16px;
        color: #252933;
        .overview {
          font-size: 12px;
          color: #8a919f;
        }
      }
    }
  }
  .content {
    padding-left: 50px;
  }
  .option-wrapper {
    display: flex;
    flex-direction: column;

    .operate {
      border-top: 1px solid #e8e8e8;
      display: flex;
      align-items: center;
      padding: 0 24px;
      justify-content: space-between;
      font-size: 14px;
      width: 100%;
      div {
        user-select: none;
        width: 100%;
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
        padding-bottom: 4px;
        &:hover {
          background-color: #e8e8e8;
        }
      }
    }
  }
  .comment-wrapper {
    padding-bottom: 12px;
  }
}
</style>
