<template>
  <FormProvider :form="form">
    <SchemaField :schema="schema" />
  </FormProvider>
  <Button @click="submitd">提交</Button>
</template>

<script setup>
import { Input, FormItem } from "ant-design-vue";
import { createForm } from "@formily/core";
import { FormProvider, createSchemaField } from "@formily/vue";
import { Button } from "ant-design-vue";
const { SchemaField } = createSchemaField({
  components: {
    Input,
    FormItem,
  },
});

const form = createForm();
const schema = {
  type: "object",
  properties: {
    input: {
      type: "string",
      "x-decorator": "FormItem",
      "x-component": "Input",
      "x-component-props": {
        placeholder: "请输入",
      },
      "x-validator": {
        required: true,
      },
    },
  },
};
const submitd = () => {
  console.log("asd");
  form.submit((v) => {
    console.log(v);
  });
};
</script>
