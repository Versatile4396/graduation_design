<template>
  <div class="followe-container-box" v-if="isFollower">
    <followBar
      v-for="item in followerList"
      :userInfo="item"
      :is-follower="true"
      @getFollowList="handleResetFollowList"
    ></followBar>
  </div>
  <div class="followe-container-box" v-else>
    <followBar
      v-for="item in followedList"
      :userInfo="item"
      :is-follower="false"
      @getFollowList="handleResetFollowList"
    ></followBar>
  </div>
</template>

<script lang="ts" setup>
// 关注列表
import { onMounted } from 'vue'
import { useFollower } from '../index'
import { getUrlQuery } from '@/utils'
import followBar from './components/follow-bar.vue'
interface Props {
  isFollower?: boolean
}
const { isFollower } = defineProps<Props>()
const { followerList, followedList, getFollowerList } = useFollower()
const handleResetFollowList = async () => {
  const { uid } = getUrlQuery()
  await getFollowerList(Number(uid))
}
onMounted(async () => {
  handleResetFollowList()
})
</script>
<style scoped lang="scss">
.followe-container-box {
}
</style>
