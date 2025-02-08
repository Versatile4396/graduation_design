<template>
    <el-select v-bind="$attrs" v-model:model-value="rValue">
        <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" />
    </el-select>
</template>

<script lang='ts' setup>
import { ElSelect } from 'element-plus'
import { ref, watch } from 'vue';
interface Props {
    options: {
        label: string
        value: string
    }[]
    value: Number
}
const props = defineProps<Props>()
const options = props.options
const emits = defineEmits(['change'])
const rValue = ref(props.value)
watch(rValue, () => {
    emits('change', rValue.value)
})
</script>
<style scoped lang="scss"></style>