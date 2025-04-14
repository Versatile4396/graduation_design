<template>
  <div class="assistance-item-container-box">
    <div class="detail-info">
      <div class="abstract-detail-info">
        <div class="userinfo">
          <chatAvatar :user-info="assistance.user_info"></chatAvatar>
          <div class="overview-wrapper">
            <div class="name">{{ assistance.user_info.nickname }}</div>
            <div class="overview">
              <span>{{ assistance.user_info.overview }}</span>
              <span style="margin: 0 4px; font-size: 12px">|</span>
              <span class="time">{{ createAt }}</span>
            </div>
          </div>
        </div>
        <div class="status">
          <el-tag>{{ assistance.status ? '已发布' : '未发布' }}</el-tag>
        </div>
      </div>
      <div class="assistance-content">
        <div class="content" v-html="assistance.content"></div>
        <div class="image-list">
          <el-image
            v-for="image of imageList"
            :key="image"
            fit="cover"
            style="max-height: 120px; max-width: 140px; border-radius: 8px"
            :src="image"
            :preview-src-list="[image]"
          />
        </div>
      </div>
    </div>
    <div class="option-wrapper" v-if="commentVisible">
      <div class="operate">
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
import chatAvatar from '@/components/chat-avatar.vue'
import Ajax from '@/ajax'
import { getUrlQuery } from '@/utils/common'

interface Props {
  assistance: Assistance
  commentVisible?: boolean
}
const props = withDefaults(defineProps<Props>(), {
  commentVisible: true
})
const createAt = computed(() => {
  return props.assistance.create_at.slice(0, 10)
})
const imageList = computed(() => {
  function isJsonString(str: string) {
    try {
      JSON.parse(str)
      return true
    } catch (e) {
      return false
    }
  }
  if (!isJsonString(props.assistance?.cover)) {
    return []
  }
  return JSON.parse(props.assistance?.cover)
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
    .abstract-detail-info {
      display: flex;
      justify-content: space-between;
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
        .popover-wrapper {
          display: flex;
        }
      }
    }
  }
  .assistance-content {
    display: flex;
    flex-direction: column;
    border-bottom: 1px solid #e8e8e8;
    .image-list {
      display: flex;
      gap: 12px;
      margin: 12px 0;
    }
  }
  .option-wrapper {
    display: flex;
    flex-direction: column;

    .operate {
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
