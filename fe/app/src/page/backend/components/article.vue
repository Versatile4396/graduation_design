<template>
  <div class="container-box">
    <div class="wrapper">
      <FilterBar
        :filter-config="filterConfig"
        :filter-value="filterValue"
        @change="handleFilterChange"
      />
      <div class="table">
        <el-table :data="dataList" v-if="dataList?.length > 0">
          <el-table-column
            v-for="table in tableConfig"
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
import { userAritcle } from './config/tableConfig'
import { onMounted, ref } from 'vue'
import FilterBar from './filter-bar.vue'
const filterValue = ref<any>({})
const handleFilterChange = async (val: any) => {
  await getDataList(val)
}

const { filterConfig, dataList, getDataList, tableConfig } = userAritcle()
onMounted(async () => {
  await getDataList()
})
</script>
<style scoped lang="scss">
.container-box {
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
