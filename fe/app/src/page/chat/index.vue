<template>
  <div class="chat-container-box">
    <div class="side-bar">
      <div class="search-bar"></div>
      <div class="list-bar">
        <chat-list
          :key="currentChat.user_id"
          :currentChat="currentChat!"
          :chatList="chatList"
          v-if="chatList?.length"
          @chatChange="handleChatChange"
        ></chat-list>
      </div>
    </div>
    <div class="chat-container">
      <div class="chat-header">
        <span class="username">{{ currentChat?.nickname + '_' + currentChat.user_id }}</span>
        <div class="operator"></div>
      </div>
      <div class="chat-content" ref="chatWrapperDom">
        <chat-wrapper
          :key="currentChat.user_id"
          :message-list="messageList"
          :avatarMe="userInfo.avatar!"
          :avatarYou="currentChat?.avatar!"
        ></chat-wrapper>
      </div>
      <div class="chat-send-wrapper">
        <inputChat @send-msg="handleSendMsg" @send-audio="handleSendAudio"></inputChat>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { getUrlQuery } from '@/utils/common'
import inputChat from './component/input-chat.vue'
import chatWrapper from './component/chat-wrapper.vue'
import ChatList from './component/chat-list.vue'
import { computed, onMounted, onUnmounted, ref } from 'vue'
import { userInfoStore } from '@/store/user'
import { useWebSocket } from './utils'
import { nextTick } from 'vue'
import type { userInfo } from '@/ajax/type/user'
import router, { routerName } from '@/router'
import * as Constant from './utils/constant'

// 获取聊天列表用户信息？
const queryInfo = getUrlQuery()
const uid = queryInfo?.uid as string
const { userInfo } = userInfoStore()
const chatWrapperDom = ref<HTMLElement>()

// 建立ws链接 发起聊天
const {
  socket,
  chatList,
  sendMessage,
  getHistoryMessage,
  getChatList,
  historyMessage,
  currentChat
} = useWebSocket(uid as string, chatWrapperDom)

const messageList = computed(() => {
  return historyMessage.value
})
onMounted(async () => {
  await getChatList()
  await getHistoryMessage(uid, getUrlQuery().toUid as string)
})

const handleChatChange = async (chat: userInfo) => {
  router.push({ name: routerName.CHAT, query: { uid: uid, toUid: chat.user_id } })
  await getHistoryMessage(uid, String(chat.user_id))
}

const handleSendMsg = (msg: string) => {
  sendMessage({
    from: uid,
    content: msg,
    messageType: 1,
    contentType: Constant.TEXT, // 消息类型，1.文本 2.图片 3.文件 4.语音 5.视频 6.位置 7.自定义,
    to: currentChat.value.user_id
  })
}

const handleSendAudio = (audio: any) => {
  sendMessage(audio)
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
