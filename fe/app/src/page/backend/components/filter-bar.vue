<template>
  <div class="filter-bar-container-box">
    <div v-for="item in filterConfig" :key="item.label">
      <div class="filter-item">
        <component
          :is="item.component"
          v-bind="item.componentProps"
          v-model="item.value"
          :filter-config="item"
          @change="handleFilterChange"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { type FilterConfigType } from './filterConfig'
interface Props {
  filterConfig: FilterConfigType[]
}
const emits = defineEmits(['change'])
const props = defineProps<Props>()
const handleFilterChange = (v: any) => {
  emits('change', v)
}
</script>
<style scoped lang="scss">
.filter-bar-container-box {
  display: grid;
  gap: 12px;
  grid-template-columns: 1fr 1fr 1fr 1fr;
  align-items: center;
}
</style>
