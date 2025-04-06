<template>
  <div class="user-manage-container-box">
    <div class="wrapper">
      <FilterBar
        :filter-config="userFilterConfig"
        :filter-value="filterValue"
        @change="handleFilterChange"
      />
      <div class="table">
        <el-table :data="userList" v-if="userList.length > 0">
          <el-table-column
            v-for="table in userTableConfig"
            :prop="table.prop"
            :label="table.label"
            :width="table?.width ?? '180'"
          >
            <template #default="scope" v-if="typeof table?.render === 'function'">
              <component :is="table.render(scope.row)" :key="scope.row" />
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import FilterBar from './filter-bar.vue'
import { userFilterConfig } from './config/filterConfig'
import { userTableConfig } from './config/tableConfig'
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
  const { data } = await Ajax.post('/user/list')
  userList.value = data
}

onMounted(async () => {
  await getUserList()
})
</script>
<style scoped lang="scss">
.user-manage-container-box {
  padding: 24px;
  width: 100%;
  height: 100%;
  background-color: #fff;
  display: flex;
  .wrapper {
    margin: 0 auto;

  }
}
</style>
