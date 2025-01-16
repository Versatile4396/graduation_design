<template>
  <FormProvider v-if="form" :form="form">
    <component :is="fieldSchema" :schema="props.schema"></component>
    <div v-if="$slots['submit_wrapper']">
      <slot name="submit_wrapper"></slot>
    </div>
    <div v-else class="default_submit_wrapper">
      <Button type="primary" @click="submitHandle">提交</Button>
      <Button type="default" @click="cancelHandle">取消</Button>
    </div>
  </FormProvider>
</template>
<script>
export default {
  name: "RxForm",
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
} from "ant-design-vue";
interface Props {
  schema: any;
}
const props = defineProps<Props>();
const emits = defineEmits(["submit", "cancel"]);

const form = createForm({ validateFirst: true });
setValidateLanguage("cn");

const submitHandle = () => {
  emits("submit");
};
const cancelHandle = () => {
  emits("cancel");
};

const fieldSchema = createSchemaField({
  components: {
    Input,
    Switch,
    Radio,
    RadioGroup,
    Select,
  },
}).SchemaField;
</script>
<style scoped lang="scss">
/* @import url(); 引入css类 */
.default_submit_wrapper {
  display: flex;
  justify-content: space-between;
  width: 300px;
}
</style>
