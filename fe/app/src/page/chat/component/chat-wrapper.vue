<template>
  <div class="chat-wrapper-container-box">
    <chat-box
      v-for="(chat, index) in chatInfo"
      :chat="chat"
      :avatar="uid !== chat.toUserId ? avatarMe : avatarYou"
      :key="index"
    ></chat-box>
  </div>
</template>

<script lang="ts" setup>
import { getUrlQuery } from '@/utils'
import { type ChatInstance } from '../type'
import chatBox from './chat-box.vue'

interface Props {
  chatInfo: ChatInstance[]
  avatarMe: string
  avatarYou: string
}
const uid = Number(getUrlQuery().uid)
const props = defineProps<Props>()
</script>
<style scoped lang="scss">
.chat-wrapper-container-box {
  display: flex;
  flex-direction: column;
  row-gap: 12px;
}
</style>
