<template>
  <div class="course-container">
    <div class="left-box">
      <div class="course-filter">
        <filterBar title="课程"></filterBar>
        <filterBar title="分类"></filterBar>
      </div>
    </div>
    <div class="right-box"></div>
  </div>
</template>

<script lang="ts" setup>
import filterBar from './component/filter-bar.vue'
</script>
<style scoped lang="scss">
.course-container {
  width: 1000px;
  margin: 0 auto;
  display: flex;
  .left-box {
    width: 800px;
    margin-right: 16px;
    background-color: #fff;
    padding: 16px;
    border-radius: 4px;

    .course-filter {
      height: 200px;
      display: flex;
      flex-direction: column;
    }
  }
  .right-box {
    padding: 16px;
    border-radius: 4px;
    background-color: #fff;
    width: 200px;
  }
}
</style>
