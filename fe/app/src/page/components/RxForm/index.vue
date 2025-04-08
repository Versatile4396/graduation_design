<template>
  <FormProvider v-if="form" :form="form">
    <component :is="fieldSchema" :schema="compShcema"></component>
    <div v-if="$slots['submit_wrapper']">
      <slot name="submit_wrapper"></slot>
    </div>
    <div v-else class="default_submit_wrapper">
      <ElButton class="submit_btn" :block="true" type="primary" @click="onFormSubmit">{{
        submitText
      }}</ElButton>
      <ElButton class="cancel_btn" :block="true" type="default" @click="onFormSubmitCancel">{{
        cancelText
      }}</ElButton>
    </div>
  </FormProvider>
</template>
<script lang="ts">
export default {
  name: 'Form'
}
</script>
<script lang="ts" setup>
import { createForm, setValidateLanguage } from '@formily/core'
import { createSchemaField, FormProvider } from '@formily/vue'
import { Input, Switch, Radio, FormItem, Password } from '@formily/element-plus'
import { RRadio, RUpload, RSelect, RVericode } from './components/index'
import { ElButton } from 'element-plus'
import { computed, onMounted, provide, ref } from 'vue'
import { formContextKey } from './context/formcontext'
interface Props {
  schema: any
  submitText?: string
  cancelText?: string
  initialValues?: any
  effects?: () => void
  submit?: (...args: any) => any
}
const props = withDefaults(defineProps<Props>(), {
  submitText: '提交',
  cancelText: '取消'
})

const compShcema = computed(() => {
  if (props.schema instanceof Function) {
    return props.schema(form.value)
  }
  return props.schema
})

const emits = defineEmits(['submit', 'cancel', 'onSubmitFailed', 'onSubmitSuccess'])
const form = ref()
form.value = createForm({
  validateFirst: true, // 只展示第一个校验错误提示
  initialValues: props.initialValues,
  effects: props.effects
})
onMounted(() => {
  console.log(form.value.values, 'initialValues')
})

setValidateLanguage('cn')

const isSubmitting = ref(false)

const updateSubmittingStatus = (status: boolean) => {
  isSubmitting.value = status
}

const onSubmit = async () => {
  if (props?.submit) {
    // 传入的submit 返回的内容会在onsubmitSuccess
    const res = await props?.submit(form.value.values)
    return res
  } else {
    emits('submit', form.value.values)
  }
  return null
}

const onFormSubmit = async () => {
  if (isSubmitting.value) return
  updateSubmittingStatus(true)
  form.value
    .submit(onSubmit)
    .then((res: any) => {
      onSubmitSuccess(res)
      updateSubmittingStatus(false)
    })
    .catch((err: any) => {
      onSubmitFaild(err)
      updateSubmittingStatus(false)
    })
}
const onSubmitFaild = (args: any) => {
  emits('onSubmitFailed', args)
}
const onFormSubmitCancel = () => {
  emits('cancel', form)
}

const onSubmitSuccess = (res: any) => {
  emits('onSubmitSuccess', res)
}

const fieldSchema = createSchemaField({
  components: {
    Input,
    Switch,
    Radio,
    Password,
    FormItem,
    RRadio,
    RUpload,
    RSelect,
    RVericode
  }
}).SchemaField

provide(formContextKey, form.value)
defineExpose(form.value)
</script>
<style scoped lang="scss">
/* @import url(); 引入css类 */
.default_submit_wrapper {
  margin-top: 16px;
  display: flex;
  justify-content: space-between;

  .submit_btn {
    width: 240px;
    margin-right: 40px;
  }

  .cancel_btn {
    width: 240px;
  }
}
</style>
