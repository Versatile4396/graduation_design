<template>
  <div class="page-content">
    <div class="login-form">
      <Tabs v-model:active-key="activeKey">
        <TabPane :key="0" tab="登录">
          <Form
            :schema="loginSchema"
            submit-text="登录"
            @submit="submitHandle"
          ></Form>
        </TabPane>
        <TabPane :key="1" tab="注册">
          <Form
            :schema="registerSchema"
            submit-text="注册"
            @submit="(v) => submitHandle(v, loginType.Register)"
          ></Form>
        </TabPane>
      </Tabs>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Tabs, TabPane } from "ant-design-vue";
import Form from "@/RxForm/index.vue";
import { loginSchema, registerSchema } from "@/RxForm/schema/login";
import { ref } from "vue";

enum loginType {
  Login,
  Register,
}

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
