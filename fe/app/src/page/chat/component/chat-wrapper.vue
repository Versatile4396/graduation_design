<template>
  <div class="chat-wrapper-container-box">
    <chat-box
      v-for="(chat, index) in chatInfo"
      :chat="chat"
      :avatar="chat.from === MsgFrom.Me ? avatarMe : avatarYou"
      :key="index"
    ></chat-box>
  </div>
</template>

<script lang="ts" setup>
import { MsgFrom, type ChatInstance } from '../type'
import chatBox from './chat-box.vue'

interface Props {
  chatInfo: ChatInstance[]
  avatarMe: string
  avatarYou: string
}
const props = defineProps<Props>()
console.log(props)
</script>
<style scoped lang="scss">
.chat-wrapper-container-box {
  display: flex;
  flex-direction: column;
  row-gap: 12px;
}
</style>
