<template>
  <div class="page-content">
    <div class="login-form">
      <el-tabs v-model:active-key="activeKey">
        <el-tab-pane :key="0" label="登录">
          <Form
            :schema="loginSchema"
            submit-text="登录"
            cancel-text="注册"
            @submit="submitHandle"
            @onSubmitSuccess="onSubmitSuccess"
          ></Form>
        </el-tab-pane>
        <el-tab-pane :key="1" label="注册">
          <Form
            :schema="registerSchema"
            submit-text="注册"
            cancel-text="登录"
            @submit="(f: any) => submitHandle(f, loginType.Register)"
          ></Form>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ElTabPane, ElTabs } from "element-plus";
import { ref } from "vue";
import Form from "@/RxForm/index.vue";
import { loginSchema, registerSchema } from "@/RxForm/schema/login";
import Ajax from "@/ajax";

enum loginType {
  Login,
  Register,
}

const submitHandle = async (values: any, type = loginType.Login) => {
  if (type) {
    Ajax.post("/login", values);
    console.log(values);
  } else {
    console.log(values);
    Ajax.post("/user/register", values);
    console.log(values);
  }
};
const onSubmitSuccess = (res: any) => {
  console.log("onSubmitSuccess", res);
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
