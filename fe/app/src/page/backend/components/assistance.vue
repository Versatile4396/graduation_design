<template>
  <div class="container-box">
    <div class="wrapper">
      <FilterBar
        :filter-config="filterConfig"
        :filter-value="filterValue"
        @change="handleFilterChange"
      />
      <div class="table"></div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { userAritcle } from './config/tableConfig'
import { ref } from 'vue'
import FilterBar from './filter-bar.vue'
const filterValue = ref<any>({})
const handleFilterChange = ref<any>({})

const { filterConfig, dataList, getDataList, tableConfig } = userAritcle()
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
