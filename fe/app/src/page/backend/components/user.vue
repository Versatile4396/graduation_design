<template>
  <div class="user-manage-container-box">
    <FilterBar :filter-config="userFilterConfig" @change="handleFilterChange" />
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import FilterBar from './filter-bar.vue'
import { userFilterConfig } from './filterConfig'

const handleFilterChange = (val: any) => {
  console.log(val, 'handleClickv')
}
</script>
<style scoped lang="scss">
.user-manage-container-box {
  height: 100%;
  padding: 24px;
  background-color: #fff;
}
</style>
