<template>
  <div class="assistance-container-box">
    <Assistance
      v-for="item in assistanceList"
      :assistance="item"
      :commentVisible="false"
    ></Assistance>
  </div>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue'
import { useAssistance } from '../index'
import Assistance from '@/page/assistance/components/assistance-item.vue'

const { assistanceList, getAssistanceList } = useAssistance()
onMounted(async () => {
  await getAssistanceList()
})
</script>
<style scoped lang="scss">
.assistance-container-box {
}
</style>
