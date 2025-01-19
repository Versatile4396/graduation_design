<template>
  <FormProvider v-if="form" :form="form">
    <component :is="fieldSchema" :schema="props.schema"></component>
    <div v-if="$slots['submit_wrapper']">
      <slot name="submit_wrapper"></slot>
    </div>
    <div v-else class="default_submit_wrapper">
      <Button
        class="submit_btn"
        :block="true"
        type="primary"
        @click="submitHandle"
        >{{ submitText }}</Button
      >
      <Button
        class="cancel_btn"
        :block="true"
        type="default"
        @click="cancelHandle"
        >{{ cancelText }}</Button
      >
    </div>
  </FormProvider>
</template>
<script lang="ts">
export default {
  name: "Form",
};
</script>
<script lang="ts" setup>
import { createForm, setValidateLanguage } from "@formily/core";

import { createSchemaField, FormProvider } from "@formily/vue";
import {
  Button,
  Input,
  Switch,
  Radio,
  RadioGroup,
  Select,
  InputPassword,
  FormItem,
} from "ant-design-vue";
interface Props {
  schema: any;
  submitText?: string;
  cancelText?: string;
  initialValues?: any;
  effects?: () => void;
}
const props = withDefaults(defineProps<Props>(), {
  submitText: "提交",
  cancelText: "取消",
});
const emits = defineEmits(["submit", "cancel"]);
const form = createForm({
  validateFirst: true, // 只展示第一个校验错误提示
  initialValues: props.initialValues,
  effects: props.effects,
});
setValidateLanguage("cn");

const submitHandle = () => {
  form.submit();
  emits("submit", form);
};
const cancelHandle = () => {
  emits("cancel", form);
};

const fieldSchema = createSchemaField({
  components: {
    Input,
    Switch,
    Radio,
    RadioGroup,
    Select,
    InputPassword,
    FormItem,
  },
}).SchemaField;
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
