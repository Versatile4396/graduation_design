<template>
  <div class="user-manage-container-box">
    <FilterBar
      :filter-config="userFilterConfig"
      :filter-value="filterValue"
      @change="handleFilterChange"
    />
    <div class="table">
      <el-table :data="userList"></el-table>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import FilterBar from './filter-bar.vue'
import { userFilterConfig } from './filterConfig'
import Ajax from '@/ajax'
const filterValue = ref<any>({})

const PageNation = {
  page: 1,
  pageSize: 10,
  total: 0
}

const handleFilterChange = async (val: any) => {
  filterValue.value = val
  await getUserList()
}

const userList = ref<any>([])

const getUserList = async () => {
  Ajax.get('/user/list', {
    page: filterValue.value.page,
    pageSize: filterValue.value.pageSize,
    name: filterValue.value.name,
    role: filterValue.value.role
  })
}
</script>
<style scoped lang="scss">
.user-manage-container-box {
  height: 100%;
  padding: 24px;
  background-color: #fff;
}
</style>
