<template>
  <div
    class="chat-box-container-box"
    :class="{
      'chat-box-container-box-left': !isFromMe,
      'chat-box-container-box-right': isFromMe
    }"
  >
    <el-avatar :size="35" class="avatar" :src="avatar" />
    <div
      v-if="chat.contentType == 1"
      class="msg-content"
      v-html="htmlContent"
      :class="{ 'msg-left': !isFromMe, 'msg-right': isFromMe }"
    ></div>
    <div v-else-if="chat.contentType === 4">
      <audio :src="audioUrl" controls preload="auto"></audio>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import { type ChatInstance } from '../type'
import { getUrlQuery } from '@/utils'
const BASE_FILE_URL = 'http://localhost:5555/api/chat/file/'
const props = defineProps<{
  chat: ChatInstance
  avatar: string
}>()
const htmlContent = computed(() => {
  return props.chat.content
})
const audioUrl = computed(() => {
  if (props.chat.url.startsWith('blob')) {
    return props.chat.url
  }
  return BASE_FILE_URL + props.chat.url || ''
})

const isFromMe = computed(() => {
  return props.chat.fromUserId == Number(getUrlQuery().uid)
})
</script>
<style scoped lang="scss">
.chat-box-container-box {
  display: flex;
  align-items: center;
  gap: 8px;
}
.chat-box-container-box-left {
  flex-direction: row;
}
.chat-box-container-box-right {
  flex-direction: row-reverse;
}
.msg-content {
  color: #262932;
  padding: 12px;
  border-radius: 4px;
  max-width: 400px;
  word-break: break-all;
  overflow-wrap: break-word;
}
.msg-left {
  background-color: #f2f3f8;
}
.msg-right {
  background-color: #d1e0fd;
}
</style>
