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
import { publicSchema } from '@/page/components/RxForm/schema/public'
import type { Form } from '@formily/core'
import { ref } from 'vue'

const form = ref<Form>()

const emits = defineEmits(['submit', 'cancel'])

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
<style lang="scss" scoped>
.public-container {
  z-index: 10;

  .title {
    border-bottom: 1px solid #ccc;
    margin-bottom: 8px;
  }

  .form {
    .submit-wrapper {
      display: flex;
      flex-direction: row-reverse;
      gap: 16px;
    }
  }
}
</style>
