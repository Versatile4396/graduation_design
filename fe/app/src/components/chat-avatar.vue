<template>
  <el-popover :width="280">
    <template #reference>
      <el-avatar :size="avatarSize" :src="userInfo.avatar" />
    </template>
    <OverView :user-info="userInfo"></OverView>
  </el-popover>
</template>
<script lang="ts" setup>
import { type UserInfo } from '@/ajax/type/index'
import OverView from '@/components/overview.vue'
import { ElAvatar } from 'element-plus'
type ElAvatarInstance = InstanceType<typeof ElAvatar>
interface Props {
  userInfo: UserInfo
  avatarSize?: ElAvatarInstance['size']
}
const props = withDefaults(defineProps<Props>(), {
  avatarSize: ''
})
console.log(props, 'avatarSize')
</script>
<style scoped lang="scss"></style>
