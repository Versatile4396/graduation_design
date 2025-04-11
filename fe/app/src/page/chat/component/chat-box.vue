<template>
  <div
    class="chat-box-container-box"
    :class="{
      'chat-box-container-box-left': !isFromMe,
      'chat-box-container-box-right': isFromMe
    }"
  >
    <div class="avatar">
      <el-avatar :size="35" :src="avatar" />
    </div>
    <div
      v-if="chat.contentType == 1"
      class="msg-content"
      v-html="chat.content"
      :class="{ 'msg-left': !isFromMe, 'msg-right': isFromMe }"
    ></div>
    <div v-else-if="chat.contentType == 2">
      <a :href="FileUrl" :download="chat.url">
        <svg-icon :icon-name="fileIcon" size="50px"></svg-icon>
      </a>
    </div>
    <div v-else-if="chat.contentType == 3">
      <el-image
        fit="cover"
        style="max-height: 100px; max-width: 100px"
        :src="FileUrl"
        :preview-src-list="[FileUrl]"
      />
    </div>
    <div v-else-if="chat.contentType === 4">
      <audio :src="FileUrl" controls preload="auto"></audio>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import { type ChatInstance } from '../type'
import { getUrlQuery } from '@/utils'
import { FileIconName } from './chat-box'
import SvgIcon from '@/assets/iconfont/SvgIcon.vue'
const BASE_FILE_URL = 'http://localhost:5555/api/chat/file/'
const props = defineProps<{
  chat: ChatInstance
  avatar: string
}>()

const isFromMe = computed(() => {
  return props.chat.fromUserId == Number(getUrlQuery().uid)
})
const fileIcon = computed(() => {
  if (props.chat?.fileSuffix) {
    return FileIconName[props.chat.fileSuffix]
  }
  const url = props.chat.url
  const suffix = url.substring(url.lastIndexOf('.') + 1)
  if (!FileIconName[suffix]) {
    return FileIconName.other
  }
  return FileIconName[suffix]
})
const FileUrl = computed(() => {
  if ([2, 3, 4].includes(props.chat.contentType)) {
    if (props.chat.url.startsWith('blob')) {
      return props.chat.url
    }
    return BASE_FILE_URL + props.chat.url || ''
  }
  return ''
})
</script>
<style scoped lang="scss">
.chat-box-container-box {
  display: flex;
  align-items: center;
  gap: 8px;
  .avatar {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
  }
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
