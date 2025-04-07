<template>
  <div class="user-form">
    <RxForm
      ref="form"
      :schema="registerSchema"
      :initial-values="{ ...row }"
      submit-text="确定更改"
      cancel-text="取消更改"
      :submit="submitHandle"
      @cancel="cancelHandle"
    ></RxForm>
  </div>
</template>

<script lang="ts" setup>
import RxForm from '@/page/components/RxForm/index.vue'
import { registerSchema } from '@/page/components/RxForm/schema/login'
import { ref } from 'vue'
import type { Form } from '@formily/core'

interface Props {
  row?: any
}
const props = defineProps<Props>()
const emits = defineEmits(['update', 'cancel'])
const form = ref<Form>()
const submitHandle = async () => {
  try {
    await form.value?.validate()
    emits('update', form.value?.values)
  } catch (error) {}
}
const cancelHandle = () => {
  emits('cancel')
}
</script>

<style scoped lang="scss">
.user-form {
  width: 396px;
}
</style>
