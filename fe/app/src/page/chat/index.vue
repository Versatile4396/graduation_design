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
        <span class="username">{{ curChatInfo?.userInfo.username }}</span>
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
import { objectEntries } from '@vueuse/core'
import inputChat from './component/input-chat.vue'
import chatWrapper from './component/chat-wrapper.vue'
import ChatList from './component/chat-list.vue'
import { nextTick, onMounted, onUnmounted, ref } from 'vue'
import Ajax from '@/ajax'
import type { ChatInstance, IChatList } from './type'
import { MsgFrom } from './type'
import { userInfoStore } from '@/store/user'
// 获取聊天列表用户信息？
const { uid, toUid } = getUrlQuery()
const query = {
  uid,
  toUid
}
const { userInfo } = userInfoStore()
const chatWrapperDom = ref<HTMLElement>()
const avatarMe = ref(userInfo.avatar!)
const queryStr = objectEntries(query)
  //@ts-ignore
  .map(([key, value]) => encodeURIComponent(key) + '=' + encodeURIComponent(value))
  .join('&')
// 建立ws链接 发起聊天
const socket = new WebSocket('ws://localhost:5555/api/chat/ws?' + queryStr)

const chatListInfo = ref<IChatList[]>([])
const curChatInfo = ref<IChatList>()
const chatInfo = ref<ChatInstance[]>([])
// 连接成功 获取历史消息
socket.onopen = () => {
  // 获取历史消息
  initChat()
}
socket.onmessage = (event) => {
  if (event.data) {
    const msg = JSON.parse(event.data)
    if (msg.code === 2) {
      // 历史聊天记录
      chatInfo.value.push(msg)
      nextTick(() => {
        chatWrapperDom.value?.scrollTo({
          top: chatWrapperDom.value?.scrollHeight,
          behavior: 'smooth'
        })
      })
    } else if (msg.code === 50001) {
      // 在线主动推送
      chatInfo.value.push(msg)
    }
  }
}

const handleSendMsg = (msg: string) => {
  // 判断socket是否在链接中
  if (socket.readyState === WebSocket.OPEN) {
    const formatMsg = {
      type: 1,
      content: msg
    }
    socket.send(JSON.stringify(formatMsg))
    chatInfo.value.push({
      from: MsgFrom.Me,
      content: msg,
      code: 1
    })
    if (chatWrapperDom.value) {
      nextTick(() => {
        chatWrapperDom.value?.scrollTo({
          top: chatWrapperDom.value?.scrollHeight,
          behavior: 'smooth'
        })
      })
    }
  }
}

const initChat = () => {
  socket.send(
    JSON.stringify({
      type: 2
    })
  )
}

const getChatHistory = async () => {
  // 获取历史消息
  const { data } = await Ajax.get(`/chat/history?uid=${uid}`)
  chatListInfo.value = data
  curChatInfo.value = data[0]
}

const handleChatChange = (chat: IChatList) => {
  // 点击切换聊天界面 需要重新获取历史消息 建立ws链接
  curChatInfo.value = chat
  chatInfo.value = []
  // 关闭链接
  socket.close()
}

onMounted(async () => {
  await getChatHistory()
})

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
