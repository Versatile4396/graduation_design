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
        <inputChat @send-msg="handleSendMsg" @AudioOnline="handleSendAudio"></inputChat>
      </div>
    </div>
  </div>
  <video id="localVideo" ref="localVideo" class="local-video"></video>
  <video id="remoteVideo" ref="remoteVideo" class="remote-video"></video>
  <el-dialog
    v-model="chatModalStatus"
    :close-on-click-modal="false"
    :close-on-press-escape="false"
    :show-close="false"
    :width="350"
  >
    <template #footer>
      <div class="chat-footer">
        <svg-icon
          style="cursor: pointer"
          icon-name="icon-jujie"
          size="45px"
          @click="handleReject"
        ></svg-icon>
        <svg-icon
          style="cursor: pointer"
          icon-name="icon-jietong"
          size="45px"
          @click="handleConnect"
          v-if="!caller"
        ></svg-icon>
      </div>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { getUrlQuery } from '@/utils/common'
import inputChat from './component/input-chat.vue'
import chatWrapper from './component/chat-wrapper.vue'
import ChatList from './component/chat-list.vue'
import SvgIcon from '@/assets/iconfont/SvgIcon.vue'
import { computed, onMounted, onUnmounted, ref } from 'vue'
import { userInfoStore } from '@/store/user'
import { useWebSocket } from './utils'
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
  currentChat,
  localPeer,
  chatModalStatus,
  called,
  caller,
  getLocalStream
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

const handleSendMsg = (msg: any) => {
  sendMessage(msg)
}

// 在线聊天

// 点击 发起语音通话
const handleSendAudio = async () => {
  await getLocalStream()
  // 拉起视频窗口
  caller.value = true
  called.value = false
  chatModalStatus.value = true
  await navigator.mediaDevices.getUserMedia({ audio: true }).then((stream) => {
    stream.getTracks().forEach((track) => {
      localPeer.addTrack(track, stream)
    })
    localPeer.createOffer().then((offer) => {
      localPeer.setLocalDescription(offer)
      let data = {
        contentType: Constant.DIAL_AUDIO_ONLINE, // 消息内容类型
        content: JSON.stringify(offer),
        type: Constant.MESSAGE_TRANS_TYPE // 消息传输类型
      }
      sendMessage(data)
    })
  })
}

// 同意 接听语音通话
const handleConnect = () => {
  const data = {
    contentType: Constant.ACCEPT_AUDIO_ONLINE, // 消息内容类型
    type: Constant.MESSAGE_TRANS_TYPE // 消息传输类型
  }
  sendMessage(data)
}

const handleReject = (isReject: boolean = true, isAudio: boolean = true) => {
  chatModalStatus.value = false
  // 如果是拒接聊天
  const contentType = isReject ? Constant.REJECT_AUDIO_ONLINE : Constant.CANCELL_AUDIO_ONLINE
  if (isAudio) {
    let data = {
      contentType,
      type: Constant.MESSAGE_TRANS_TYPE,
      //@ts-ignore
      content: Constant.ConstantDetail[contentType]
    }
    sendMessage(data)
  }
}

onUnmounted(() => {
  socket.close()
})
</script>
<style scoped lang="scss">
.local-video {
  position: fixed;
  top: 50px;
  height: 600px;
  width: 300px;
  z-index: 9999;
}
.remote-video {
  position: fixed;
  right: 0px;
  height: 600px;
  width: 300px;
  z-index: 9999;
}
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
.chat-footer {
  margin: 0 auto;
  width: 200px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 68px;
}
</style>
