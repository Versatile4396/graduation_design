<template>
  <div class="assistance-container">
    <div class="left-box">
      <div class="submit-wrapper">
        <div class="input-wrapper">
          <InputChat ref="inputChat" :enter-lock="true" footer-msg=""></InputChat>
        </div>
        <div class="submit-botton">
          <el-popover
            v-model:visible="popoverStatus"
            trigger="click"
            placement="bottom"
            :width="580"
            popper-style="box-shadow: rgb(14 18 22 / 35%) 0px 10px 38px -10px, rgb(14 18 22 / 20%) 0px 10px 20px -15px; padding: 20px;"
          >
            <template #reference>
              <el-button style="width: 100px" type="primary">发布</el-button>
            </template>
            <SubmitForm @submit="submitHandle" @cancel="popoverStatus = false"></SubmitForm>
          </el-popover>
        </div>
      </div>
    </div>
    <div class="right-box">
      <div class="personal-info">
        <div class="top">
          <div class="avatar">
            <el-avatar :src="userInfo.avatar"></el-avatar>
            <span class="name">
              {{ userInfo.user_name }}
            </span>
          </div>
          <div class="name"></div>
        </div>
        <div class="bottom"></div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { userInfoStore } from '@/store/user'
import { storeToRefs } from 'pinia'
import InputChat from '../chat/component/input-chat.vue'
import { ref } from 'vue'
import SubmitForm from './components/public-form.vue'
import Ajax from '@/ajax'
import { getUrlQuery } from '@/utils/common'
import { computed } from 'vue'
import { Message } from '@/utils/message'

const { userInfo } = storeToRefs(userInfoStore())
const popoverStatus = ref(false)
interface Iparams {
  category_id: number
  title: string
  content: string
}
const inputChat = ref()
const assistanceContent = computed(() => {
  return inputChat?.value?.value || ''
})
const submitHandle = async (value: Iparams) => {
  if (assistanceContent.value?.trim() == '') {
    Message.warn('请输入内容')
    return
  }
  const params = {
    ...value,
    user_id: Number(getUrlQuery().uid),
    content: assistanceContent.value
  }
  await Ajax.post('/assistance/create', params)
  popoverStatus.value = false
}
</script>
<style scoped lang="scss">
.assistance-container {
  width: 1000px;
  margin: 0 auto;
  display: flex;
  .left-box {
    width: 800px;
    margin-right: 16px;
    background-color: #fff;
    padding: 16px;
    border-radius: 4px;
    .submit-wrapper {
      height: 218px;
      .input-wrapper {
        background-color: #f8f7f7;
        border-radius: 4px;
        height: 140px;
      }
      .submit-botton {
        height: 78px;
        padding-top: 12px;
        display: flex;
        flex-direction: row-reverse;
      }
    }
  }
  .right-box {
    width: 200px;
    .personal-info {
      border-radius: 4px;
      padding: 16px;
      background-color: #fff;
      display: flex;
      flex-direction: column;
      .top {
        display: flex;
        align-items: center;
        margin-bottom: 16px;
        padding-bottom: 16px;
        border-bottom: 1px solid #e0e0e0;
        .avatar {
          display: flex;
          align-items: center;
          gap: 8px;
          margin-right: 16px;
        }
      }
    }
  }
}
</style>
