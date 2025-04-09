<template>
  <div class="chat-container-box">
    <div class="side-bar">
      <div class="search-bar"></div>
      <div class="list-bar">
        <chat-list
          :currentChat="curChatInfo!"
          :chatList="chatListInfo"
          v-if="chatListInfo.length"
          @chatChange="handleChatChange"
        ></chat-list>
      </div>
    </div>
    <div class="chat-container">
      <div class="chat-header">
        <span class="username">{{ curChatInfo?.userInfo.nickname }}</span>
        <div class="operator"></div>
      </div>
      <div class="chat-content" ref="chatWrapperDom">
        <chat-wrapper :chatInfo :avatarMe :avatarYou="curChatInfo?.userInfo.avatar!"></chat-wrapper>
      </div>
      <div class="chat-send-wrapper">
        <inputChat @send-msg="handleSendMsg"></inputChat>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { getUrlQuery } from '@/utils/common'
import inputChat from './component/input-chat.vue'
import chatWrapper from './component/chat-wrapper.vue'
import ChatList from './component/chat-list.vue'
import { onUnmounted, ref } from 'vue'
import type { ChatInstance, IChatList } from './type'
import { userInfoStore } from '@/store/user'
import { useWebSocket } from './utils'
// 获取聊天列表用户信息？
const { uid, toUid } = getUrlQuery()
const { userInfo } = userInfoStore()
const chatWrapperDom = ref<HTMLElement>()
const avatarMe = ref(userInfo.avatar!)

// 建立ws链接 发起聊天
const { socket, sendMessage } = useWebSocket(uid as string)

const chatListInfo = ref<IChatList[]>([])
const curChatInfo = ref<IChatList>()
const chatInfo = ref<ChatInstance[]>([])

const handleChatChange = () => {}
const handleSendMsg = (msg: string) => {
  sendMessage({
    from: uid,
    content: msg,
    messageType: 1,
    to: toUid
  })
}
onUnmounted(() => {
  socket.close()
})
</script>
<style scoped lang="scss">
.chat-container-box {
  display: flex;
  margin: 0 auto;
  width: 1200px;
  height: 708px;
  border-radius: 4px;
  background-color: #fff;

  .side-bar {
    width: 340px;
    height: 100%;
    border-right: 1px solid #e4e6ea;

    .search-bar {
      height: 68px;
      border-bottom: 1px solid #e4e6ea;
    }

    .list-bar {
      display: flex;
      flex-direction: column;
    }
  }

  .chat-container {
    flex: 1;
    display: flex;
    flex-direction: column;

    .chat-header {
      height: 68px;
      border-bottom: 1px solid #e4e6ea;
      padding: 0 24px;
      line-height: 68px;
      .username {
        font-weight: 500;
        font-size: 18px;
        line-height: 24px;
        color: #262932;
        text-align: center;
        flex-shrink: 1;
        overflow: hidden;
        text-overflow: ellipsis;
        cursor: pointer;
        user-select: none;
      }
    }

    .chat-content {
      overflow-y: auto;
      padding: 24px;
      flex: 1;
    }

    .chat-send-wrapper {
      border-top: 1px solid #e4e6ea;
      height: 180px;
    }
  }
}
</style>
