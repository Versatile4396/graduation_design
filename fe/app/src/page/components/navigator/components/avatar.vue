<template>
    <el-dropdown placement="top-start" trigger="click">
        <el-avatar class="avatar" :src="userInfo.avatar" />

        <template #dropdown>
            <el-dropdown-menu>
                <el-dropdown-item @click="handlePersonal">个人中心</el-dropdown-item>
                <el-dropdown-item @click="handleLoginOut">退出登录</el-dropdown-item>
            </el-dropdown-menu>
        </template>
    </el-dropdown>
</template>

<script lang='ts' setup>
import { userInfoStore } from '@/store/user';
import { storeToRefs } from 'pinia';
import router, { routerName } from "@/router";


const { userInfo } = storeToRefs(userInfoStore());
const handleLoginOut = () => {
    router.push({ name: routerName.LOGIN });
    router.replace({ query: {} });
    userInfoStore().setUserInfo({});
};
const handlePersonal = () => {
    router.push({ name: routerName.PERSONAL });
};

</script>
<style scoped lang="scss"></style>