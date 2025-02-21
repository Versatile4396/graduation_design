<template>
    <div class="comments-input-container" v-if="isShow">
        <el-input v-model="commentText" resize="none" :rows="4" type="textarea" :placeholder="placeholder"
            @keyup.enter="handleComment">
        </el-input>
        <div class="submit-btn">
            <el-button width="120px" type="primary" :disabled="!commentText" @click="handleComment">发送</el-button>
        </div>
    </div>

</template>

<script lang='ts' setup>
import { ref } from 'vue';


interface Props {
    isShow: boolean
    placeholder?: string
}
const props = withDefaults(defineProps<Props>(), {
    isShow: false,
    placeholder: "友善评论，平等表达",
});
const emits = defineEmits([
    "comment",
]);

const commentText = ref('')

const handleComment = () => {
    emits("comment", commentText.value)
    commentText.value = ''
}

</script>
<style scoped lang="scss">
.comments-input-container {
    .submit-btn {
        display: flex;
        flex-direction: row-reverse;
        margin-top: 12px;
    }
}
</style>