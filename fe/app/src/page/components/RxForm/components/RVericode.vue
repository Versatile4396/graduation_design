<template>
  <div class="vericode-container-box">
    <el-input
      v-model:model-value="rValue"
      :maxlength="6"
      placeholder="请输入验证码"
      @input="handleChange"
    />
    <el-button :disabled="buttonDisabled" type="primary" @click="getVeriCode">
      {{ buttonText }}</el-button
    >
  </div>
</template>

<script lang="ts" setup>
import Ajax from '@/ajax'
import { Message } from '@/utils'
import type { Form } from '@formily/core'
import { inject, ref } from 'vue'
import { formContextKey } from '@/page/components/RxForm/context/formcontext.ts'
interface Props {
  value: string
}
const form = inject<Form>(formContextKey)
const buttonDisabled = ref(false)
const buttonText = ref('获取验证码')
const loading = ref(false)
function validateEmail(email: string) {
  const regex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/
  return regex.test(email)
}
const getVeriCode = async () => {
  loading.value = true
  if (validateEmail(form?.values?.email || '') === false) {
    Message.err('请输入正确的邮箱')
    return
  }
  await Ajax.post('/user/sendEmailCode', {
    email: [form?.values?.email || '']
  })
  loading.value = false
  let count = 60
  const timer = setInterval(() => {
    buttonText.value = `${count}s`
    count--
    if (count === 0) {
      clearInterval(timer)
      buttonDisabled.value = false
      buttonText.value = '获取验证码'
    }
    buttonDisabled.value = true
  }, 1000)
}
const props = defineProps<Props>()
const rValue = ref(props.value)
const emits = defineEmits(['change'])
const handleChange = (val: string) => {
  emits('change', val)
}
</script>
<style scoped lang="scss">
.vericode-container-box {
  display: flex;
  gap: 12px;
  justify-content: center;
}
</style>
