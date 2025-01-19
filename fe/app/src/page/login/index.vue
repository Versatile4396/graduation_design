<template>
  <div class="page-content">
    <div class="login-form">
      <Tabs v-model:active-key="activeKey">
        <TabPane :key="0" tab="登录">
          <FormItem
            label="用户名"
            name="username"
            :rules="[
              { required: true, message: 'Please input your username!' },
            ]"
          >
            <Input v-model:value="formState.username" />
          </FormItem>

          <FormItem
            label="密码"
            name="password"
            :rules="[{ required: true, message: '请输入密码' }]"
          >
            <InputPassword v-model:value="formState.password" />
          </FormItem>

          <FormItem name="remember" :wrapper-col="{ offset: 8, span: 16 }">
            <Checkbox v-model:checked="formState.remember"
              >Remember me</Checkbox
            >
          </FormItem>

          <FormItem :wrapper-col="{ offset: 8, span: 16 }">
            <Button type="primary" html-type="submit">Submit</Button>
          </FormItem>
        </TabPane>
        <TabPane :key="1" tab="注册"> </TabPane>
      </Tabs>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {
  Tabs,
  TabPane,
  FormItem,
  Checkbox,
  InputPassword,
  Input,
  Button,
} from "ant-design-vue";
import { ref, reactive } from "vue";

enum loginType {
  Login,
  Register,
}

interface FormState {
  username: string;
  password: string;
  remember: boolean;
}

const formState = reactive<FormState>({
  username: "",
  password: "",
  remember: true,
});

const submitHandle = (form: any, type = loginType.Login) => {
  if (type) {
    console.log(form.values);
  } else {
    console.log(form.values);
  }
};
const activeKey = ref(loginType.Login);
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
