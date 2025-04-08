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
import { useUserList } from './config/tableConfig'
import { debounce } from 'lodash-es'
const filterValue = ref<any>({})

const handleFilterChange = debounce(async (value: any) => {
  filterValue.value = value
  await getUserList(filterValue.value)
}, 500)

const { userList, getUserList, userTableConfig, userFilterConfig } = useUserList()

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
