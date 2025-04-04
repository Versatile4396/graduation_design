<template>
  <div class="page-content">
    <div class="login-form">
      <el-tabs v-model="activeKey" type="border-card">
        <el-tab-pane label="登录" name="login">
          <RxForm
            :schema="loginSchema"
            submit-text="登录"
            cancel-text="去注册"
            :initial-values="{
              username: 'username',
              password: '123456'
            }"
            :submit="submitHandle"
            @onSubmitSuccess="onSubmitSuccess"
            @cancel="() => (activeKey = loginType.Register)"
          ></RxForm>
        </el-tab-pane>
        <el-tab-pane label="注册" name="register">
          <RxForm
            ref="registerForm"
            :schema="registerSchema"
            :initial-values="{
              username: 'username',
              email: 'asd@qq.com',
              confirm_password: '123456',
              password: '123456',
              gender: 1,
              nickname: '章三',
              overview: '一个叫张三的man'
              
            }"
            submit-text="注册"
            cancel-text="去登录"
            :submit="(f: any) => submitHandle(f, loginType.Register)"
            @cancel="() => (activeKey = loginType.Login)"
            @onSubmitSuccess="onSubmitSuccess"
          ></RxForm>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ElTabPane, ElTabs } from 'element-plus'
import { ref } from 'vue'
import RxForm from '@/page/components/RxForm/index.vue'
import { loginSchema, registerSchema } from '@/page/components/RxForm/schema/login'
import Ajax from '@/ajax'
import { userInfoStore } from '@/store/user'
import router from '@/router'
import type { userInfo } from '@/ajax/type/user'

const { setUserInfo } = userInfoStore()
enum loginType {
  Login = 'login',
  Register = 'register'
}
const submitHandle = async (values: any, type = loginType.Login) => {
  try {
    if (type === loginType.Login) {
      const { data } = await Ajax.post('/user/login', values)
      return data
    } else {
      const { data } = await Ajax.post('/user/register', values)
      return data
    }
  } catch (error) {}
}
const onSubmitSuccess = (data: userInfo) => {
  // 登录注册成功之后的回掉 存储token信息 存储userInfo
  setUserInfo(data)
  // 携带uid跳转主页
  router.push({
    name: 'home',
    query: {
      uid: data.user_id
    }
  })
}
const activeKey = ref(loginType.Login)
</script>
<style scoped lang="scss">
/* @import url(); 引入css类 */
.page-content {
  position: fixed;
  width: 100vw;
  height: 100vh;

  background-image: url('../../assets//image/login.png');
  background-size: cover;
  background-attachment: fixed;

  .login-form {
    padding: 16px;
    width: 360px;
    background-color: #fff;
    border-radius: 12px;
    position: absolute;
    right: 12%;
    top: 50%;
    transform: translateY(-50%);
  }
}
</style>
