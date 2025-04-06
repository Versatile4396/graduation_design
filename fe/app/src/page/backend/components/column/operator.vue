<template>
  <div class="operator-container-box">
    <el-button type="primary">编辑</el-button>
    <el-popover trigger="click" width="90px" v-model:visible="deleteVisible">
      <template #reference>
        <el-button type="danger"> {{ buttonText[0] }}</el-button>
      </template>
      <div style="display: flex; flex-direction: column; align-items: center">
        <div>你确定要删除吗？</div>
        <div class="delete-confirm" style="margin-top: 8px">
          <el-button type="danger" size="small" @click="handleDeleteClick">
            {{ buttonText[0] }}</el-button
          >
          <el-button type="primary" size="small" @click="handleCancelClick">{{
            buttonText[1]
          }}</el-button>
        </div>
      </div>
    </el-popover>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

interface Props {
  row: any
  handleDelete?: (row: any) => void
  handleEditor?: (row: any) => void
  buttonText: string[]
}
const emits = defineEmits(['delete', 'editor'])
const props = defineProps<Props>()
const deleteVisible = ref(false)
const handleDeleteClick = async () => {
  if (props.handleDelete) {
    await props.handleDelete(props.row)
  }
}
const handleCancelClick = () => {
  deleteVisible.value = false
}
</script>
<style scoped lang="scss">
.operator-container-box {
}
</style>
