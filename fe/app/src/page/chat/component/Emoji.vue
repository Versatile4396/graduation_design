<template>
    <div class="emoji-container">
        <!-- 最近使用 -->
        <div class="use-wrap" v-if="historyEmojiList.length">
            <div class="emoji-title">最近使用</div>
            <div class="recently-use-emoji">
                <div class="emoji-item" v-for="(item, index) in historyEmojiList" :key="index">
                    <img :src="`../../../assets/gif/${item}.gif`" alt="">
                </div>
            </div>
        </div>
        <!-- 所有表情列表 -->
        <div class="emoji-title">小黄脸表情包</div>
        <div class="all-emoji">
            <div class="emoji-item" v-for="(item, index) in 203" :key="index" @click="appendEmoji(item)">
                <img :src="`/src/assets/gif/${index + 1}.gif`" alt="">
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const emits = defineEmits(['selectEmoji'])
const appendEmoji = (item: number) => {
    emits('selectEmoji', item)
}
const historyEmojiList = ref<number[]>([])
</script>

<style lang="scss" scoped>
.emoji-container {
    width: 100%;
    height: 350px;
    padding: 5px 10px 15px;
    box-sizing: border-box;
    overflow: auto;

    &::-webkit-scrollbar {
        display: none;
    }
}

.emoji-title {
    font-size: 12px;
    margin: 10px 0 5px;
}

.recently-use-emoji,
.all-emoji {
    width: 100%;
    display: flex;
    flex-wrap: wrap;

    .emoji-item {
        width: 30px;
        height: 30px;
        margin: 5px 10px;
        border-radius: 4px;
        cursor: pointer;

        img {
            width: 100%;
            height: 100%;
        }
    }

    .emoji-item:hover {
        background-color: $fill-color-2;
    }
}
</style>