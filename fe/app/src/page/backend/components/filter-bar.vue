<template>
  <div class="filter-bar-container-box">
    <div v-for="item in filterConfig" :key="item.label">
      <div class="filter-item">
        <el-select
          @change="(v:any)=>handleSelectChange(item.key!,v)"
          v-if="item.type == FilterType.select"
          v-bind="item.componentProps"
          v-model="filterValue[item.key!]"
        >
          <el-option
            v-for="option in item.componentProps?.options || []"
            :key="option.value"
            :label="option.label"
            :value="option.value"
          ></el-option>
        </el-select>
        <el-input
          v-if="item.type == FilterType.input"
          v-bind="item.componentProps"
          v-model="filterValue[item.key!]"
          @input="
            (v) => {
              handleInputChange(item.key!, v)
            }
          "
        ></el-input>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { type FilterConfigType, FilterType } from './filterConfig'
interface Props {
  filterConfig: FilterConfigType[]
  filterValue: any
}
const emits = defineEmits(['change'])
const props = defineProps<Props>()

const value = ref<any>({})
const handleInputChange = (key: string, v: string) => {
  value.value[key] = v
  handleFilterChange(value.value)
}

const handleSelectChange = (key: string, v: string) => {
  value.value[key] = v
  handleFilterChange(value.value)
}

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
