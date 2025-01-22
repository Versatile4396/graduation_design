<template>
  <div class="main-container">
    <div class="index-nav">
      <div>inde-nav</div>
      <div>inde-nav</div>
      <div>inde-nav</div>
      <div>inde-nav</div>
      <div>inde-nav</div>
    </div>
    <div class="right-container">
      <div class="article-content">
        <div>argicl</div>
        <div>argicl</div>
        <div>argicl</div>
      </div>
      <div class="irrelevant-container">
        <div class="irelevant-content">reaasdasd</div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup></script>
<style scoped lang="scss">
/* @import url(); 引入css类 */
.main-container {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  .index-nav {
    position: fixed;
    background-color: #fff;
    min-width: 180px;
    margin-right: 20px;
    padding: 8px;
  }
  .right-container {
    margin-left: 215px;
    padding: 8px;
    display: flex;
    .article-content {
      background-color: #fff;
      min-width: 720px;
      margin-right: 16px;
    }
    .irrelevant-container {
      .irelevant-content{
        background-color: #fff;
        padding: 8px;
      }
      min-width: 260px;
    }
  }
}
</style>
