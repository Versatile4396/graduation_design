<template>
    <div class="aritcle-detail-container-box">
        <Editor v-model="aContent" :defaultConfig="editorConfig" :mode="mode" @onCreated="handleCreated" />
    </div>
</template>

<script lang='ts' setup>

const aContent = ref("");

</script>
<style scoped lang="scss">
.aritcle-detail-container-box {
    margin: 0 auto;
    width: 820px;
    border-radius: 4px 4px 0 0;
    padding: 2.667rem 2.667rem 0 2.667rem;
    height: 3000px;
    background-color: #fff;
}
</style>