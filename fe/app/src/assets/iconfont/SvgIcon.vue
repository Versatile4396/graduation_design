<template>
    <svg :class="svgClass" :style="{ '--size': size }" aria-hidden="true">
        <use :xlink:href="iconClassName" :fill="fillColor" />
    </svg>
</template>
<script setup lang="ts">
import { computed } from 'vue';
const props = defineProps({
    iconName: {
        type: String,
        required: true
    },
    className: {
        type: String,
        default: ''
    },
    color: {
        type: String,
        default: '#8b919e'
    },
    size: {
        type: String,
        default: '1.5rem'
    }
});
// 图标在 iconfont 中的名字
const iconClassName = computed(() => {
    return `#${props.iconName}`;
})
const fillColor = computed(() => {
    if (!props.color) {
        return '#8b919e'
    }
    return props.color;
})
// 给图标添加上类名
const svgClass = computed(() => {
    if (props.className) {
        return `svg-icon ${props.className}`;
    }
    return 'svg-icon';
});
</script>
<style scoped>
.svg-icon {
    width: var(--size);
    height: var(--size);
    position: relative;
    fill: currentColor;
    vertical-align: -2px;
}
</style>