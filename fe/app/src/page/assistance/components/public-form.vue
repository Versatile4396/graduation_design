<template>
  <div class="public-container">
    <div class="title">
      <h3>发布文章</h3>
    </div>
    <div class="form">
      <RxForm :schema="publicSchema" :initial-values="{ category_id: 1 }" ref="form">
        <template #submit_wrapper>
          <div class="submit-wrapper">
            <el-button style="width: 88px" @click="handleCancel">取消</el-button>
            <el-button type="primary" @click="handleSubmit">确定发布</el-button>
          </div>
        </template>
      </RxForm>
    </div>
  </div>
</template>

<script lang="ts" setup>
import RxForm from '@/page/components/RxForm/index.vue'
import { publicSchema } from '@/page/components/RxForm/schema/assistance-public'
import { ref } from 'vue'
import type { Form } from '@formily/core'

const emits = defineEmits(['submit', 'cancel'])
const form = ref<Form>()

const handleSubmit = async () => {
  try {
    await form.value?.validate()
    emits('submit', form.value?.values)
  } catch (error) {}
}
const handleCancel = () => {
  emits('cancel')
}
</script>

<style scoped lang="scss">
.public-form-container-box {
}
</style>
