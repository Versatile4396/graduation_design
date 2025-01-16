<template>
  <div class="page-content">
    <div class="login-form">
      <Form :schema="loginSchema"></Form>
    </div>
  </div>
</template>

<script lang="ts" setup>
import Form from "@/RxForm/index.vue";
import { loginSchema } from "@/RxForm/login";
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
