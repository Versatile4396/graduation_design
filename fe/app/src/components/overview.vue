<template>
  <div class="overview-container-box">
    <div class="top">
      <div class="userinfo">
        <el-avatar :src="userInfo.avatar"></el-avatar>
        <div class="detail">
          <div class="name">{{ userInfo.nickname }}</div>
          <div class="overview">{{ userInfo.overview }}</div>
        </div>
      </div>
      <div class="operator">
        <el-button type="primary" @click="handleChatClick">
          <template #default><span class="chat-btn">私聊</span></template>
        </el-button>
      </div>
    </div>
    <div class="footer" v-if="showFooter"></div>
  </div>
</template>

<script lang="ts" setup>
import Ajax from '@/ajax'
import type { UserInfo } from '@/page/article/component/types'
import { getUrlQuery } from '@/utils/common'
import { goToChat } from '@/utils/goto'
interface Props {
  userInfo: UserInfo
  showFooter?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  showFooter: true
})

const handleChatClick = async () => {
  const uid = Number(getUrlQuery().uid)
  const toUid = props.userInfo.user_id
  await goToChat(uid, toUid)
}
</script>
<style scoped lang="scss">
.overview-container-box {
  border-radius: 4px;
  .top {
    display: flex;
    flex-direction: column;
    .userinfo {
      display: flex;
      .detail {
        display: flex;
        flex-direction: column;
        margin-left: 8px;
        gap: 4px;
        .name {
          font-size: 16px;
          font-weight: 500;
          color: #252933;
        }
        .overview {
          font-size: 12px;
          color: #8a919f;
        }
      }
    }
    .operator {
      margin-top: 12px;
      display: flex;
      justify-content: center;
      .chat-btn {
        width: 220px;
      }
    }
  }
}
</style>
