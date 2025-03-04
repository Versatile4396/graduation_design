<template>
    <div class="chat-container-box">
        <div class="side-bar">
            <div class="search-bar"></div>
            <div class="list-bar"></div>
        </div>
        <div class="chat-container">
            <div class="chat-header"></div>
            <div class="chat-content"></div>
            <div class="chat-send-wrapper"></div>
        </div>
    </div>
</template>

<script lang='ts' setup>
</script>
<style scoped lang="scss">
.chat-container-box {
    display: flex;
    margin: 0 auto;
    width: 1200px;
    height: 708px;
    border-radius: 4px;
    background-color: #fff;

    .side-bar {
        width: 340px;
        height: 100%;
        border-right: 1px solid #e4e6ea;

        .search-bar {
            height: 68px;
        }

        .list-bar {
            display: flex;
            flex-direction: column;
        }
    }

    .chat-container {
        flex: 1;
        display: flex;
        flex-direction: column;

        .chat-header {
            height: 68px;
            border-bottom: 1px solid #e4e6ea;
        }

        .chat-content {
            flex: 1;
        }

        .chat-send-wrapper {
            border-top: 1px solid #e4e6ea;
            height: 180px;
        }
    }
}
</style>