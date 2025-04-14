<template>
  <div class="follow-bar-container-box">
    <div class="user-info">
      <chatAvatar :user-info="userInfo"></chatAvatar>
      <span class="user-name">{{ userInfo.nickname }}</span>
    </div>
    <div class="opt">
      <el-popover v-if="isFollower">
        <template #reference>
          <el-button @click="handleCancleFollowClick(userInfo.user_id!)">
            <template #default><span class="chat-btn">已关注</span></template>
          </el-button>
        </template>
        <div style="text-align: center">点击取消关注</div>
      </el-popover>
      <el-button v-else type="primary" plain @click="handleFollowClick(userInfo.user_id!)">
        <template #default><span class="chat-btn">关注</span></template>
      </el-button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import Ajax from '@/ajax'
import type { UserInfo } from '@/ajax/type/user'
import { getUrlQuery } from '@/utils'
import chatAvatar from '@/components/chat-avatar.vue'

interface Props {
  userInfo: UserInfo
  isFollower?: boolean
}
const { userInfo, isFollower } = defineProps<Props>()

const emits = defineEmits(['getFollowList'])

const handleFollowClick = async (user_id: number) => {
  const { uid } = getUrlQuery()
  await Ajax.post('/follow/create', {
    follower_id: Number(uid),
    followed_id: user_id!
  })
  emits('getFollowList')
}
const handleCancleFollowClick = async (user_id: number) => {
  const { uid } = getUrlQuery()
  await Ajax.post('/follow/delete', {
    follower_id: Number(uid),
    followed_id: user_id!
  })
  emits('getFollowList')
}
</script>
<style scoped lang="scss">
.follow-bar-container-box {
  display: flex;
  justify-content: space-between;
  .user-info {
    display: flex;
    align-items: center;
    gap: 16px;
  }
}
</style>
