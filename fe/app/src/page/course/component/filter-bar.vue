<template>
  <div class="filter-bar-container-box">
    <span>{{ title }}</span>
    <div class="filter-item"></div>
  </div>
</template>

<script lang="ts" setup>
interface Props {
  title: string
  filterConfig?: {
    label: string
    value: string
  }[]
}

const props = defineProps<Props>()
const emits = defineEmits(['filterChange'])
</script>
<style scoped lang="scss">
.filter-bar-container-box {
}
</style>
