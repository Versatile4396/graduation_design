<template>
  <div class="chat-list-container-box">
    <div
      class="chat-lite-item"
      :class="{ 'active-chat-item': chat.userInfo.user_id === currentChat.userInfo.user_id }"
      v-for="(chat, index) in props.chatList"
      :key="index"
      @click="handleChatChange(chat)"
    >
      <div class="avatar">
        <el-avatar :size="40" :src="chat.userInfo.avatar"></el-avatar>
      </div>
      <div class="chat-info">
        <div class="username">{{ chat.userInfo.username }}</div>
        <div class="chat-msg" v-html="chat.latestMsg"></div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import type { IChatList } from '../type'

interface Props {
  chatList: IChatList[]
  currentChat: IChatList
}
const props = defineProps<Props>()
const emits = defineEmits(['chatChange'])
const handleChatChange = (chat: IChatList) => {
  emits('chatChange', chat)
}
</script>
<style scoped lang="scss">
.chat-list-container-box {
  display: flex;
  align-items: center;
  flex-direction: column;
  position: relative;
  .chat-lite-item {
    display: flex;
    align-items: center;
    width: 100%;
    padding: 12px;
    .avatar {
      display: flex;
      width: 40px;
      align-items: center;
    }
    .chat-info {
      display: flex;
      flex-direction: column;
      width: 100%;
      margin-left: 12px;
      .username {
        font-size: 14px;
        font-weight: 600;
        margin-bottom: 4px;
      }
      .chat-msg {
        color: #888;
        font-size: 12px;
      }
    }
  }
  .chat-lite-item:hover {
    background-color: #f2f3f8;
    cursor: pointer;
  }
  .active-chat-item {
    background-color: #ebf2fe;
  }
}
</style>
