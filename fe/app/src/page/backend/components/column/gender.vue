<template>
  <el-tag type="success" v-if="gender === 1">男</el-tag>
  <el-tag type="primary" v-else-if="gender === 2">女</el-tag>
  <el-tag type="warning" v-else="gender === 0">未知</el-tag>
</template>

<script lang="ts" setup>
interface Props {
  gender?: number
}
const props = defineProps<Props>()
console.log(props.gender, 'props.gender')
</script>

<style scoped lang="scss"></style>
