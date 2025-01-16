<template>
  <div class="page-content">
    <div class="login-form">
      <FormProvider :form="form">
        <Field
          name="name"
          required
          :decorator="[FormItem]"
          :component="[Input, { placeholder: '请输入邮箱/手机号' }]"
        />
        <Field
          name="password"
          required
          :decorator="[FormItem]"
          :component="[Input, { type: 'password', placeholder: '请输入密码' }]"
          :reactions="createPasswordEqualValidate('confirm_password')"
        />
        <Field
          name="confirm_password"
          required
          :decorator="[FormItem]"
          :component="[Input, { type: 'password', placeholder: '确认密码' }]"
          :reactions="createPasswordEqualValidate('password')"
        />
        <!-- <FormConsumer>
          <template #default="{ form }">
            <div style="white-space: pre">
              {{ JSON.stringify(form.values, null, 2) }}
            </div>
          </template>
        </FormConsumer> -->
      </FormProvider>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { createForm, isVoidField, setValidateLanguage } from "@formily/core";
import {
  FormProvider,
  FormConsumer,
  Field,
  connect,
  mapProps,
} from "@formily/vue";
import { Form, Input } from "ant-design-vue";

const FormItem = connect(
  Form.Item,
  mapProps(
    { validateStatus: true, title: "label", required: true },
    (_, field) => {
      return {
        help: !isVoidField(field)
          ? field.selfErrors.length
            ? field.selfErrors
            : undefined
          : undefined,
        extra: field.description,
      };
    }
  )
);
setValidateLanguage("en");

const form = createForm({ validateFirst: true });
const createPasswordEqualValidate = (equalName: string) => (field: any) => {
  if (
    form.values.confirm_password &&
    field.value &&
    form.values[equalName] !== field.value
  ) {
    field.selfErrors = ["Password does not match Confirm Password."];
  } else {
    field.selfErrors = [];
  }
};
</script>
<style scoped lang="scss">
/* @import url(); 引入css类 */
.page-content {
  position: fixed;
  width: 100vw;
  height: 100vh;

  background-image: url("../../assets//image/login.png");
  background-size: cover;
  background-attachment: fixed;

  .login-form {
    padding: 16px;
    width: 360px;
    height: 500px;
    background-color: #fff;
    border-radius: 12px;
    position: absolute;
    right: 12%;
    top: 50%;
    transform: translateY(-50%);
  }
}
</style>
