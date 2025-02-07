<template>
  <div class="container">
    <div
      class="tag"
      :class="{ 'active-tag': item.value === checkedTag }"
      v-for="item in tagConfig"
      @click="handleChange(item.value)"
    >
      {{ item.label }}
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from "vue";

interface Props {
  value: any;
  options: any[];
}
const props = defineProps<Props>();

const emits = defineEmits(["change"]);
const checkedTag = ref(props.value);
const tagConfig = computed(() => {
  return props.options.map((item) => {
    return {
      label: item.label,
      value: item.value,
    };
  });
});
const handleChange = (val: any) => {
  checkedTag.value = val;
  emits("change", val);
};
</script>
<style scoped lang="scss">
.container {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
  .tag {
    text-align: center;
    padding: 2px;
    border-radius: 4px;
    background-color: rgb(244, 245, 245);
    cursor: pointer;
  }
  .active-tag {
    color: #1d7dfa;
    background-color: #e8f3ff !important;
  }
  .tag:hover {
    background-color: #e5e6eb;
  }
}
</style>
