<template>
    <div class="comment-container-box">
        <div class="first-comment-box">
            <el-avatar class="avatar" :src="userInfo.avatar" />
            <div class="comment-content">
                <div class="username">
                    {{ userInfo.username }}
                </div>
                <div class="comment-text">
                    {{ commentInfo.content || " a" }}
                </div>
                <div class="comment-brief">
                    <span class="create-time">
                        {{ timeFormat(commentInfo.create_time) }}
                    </span>
                    <span class="reply" @click="secondaryReview">
                        <svg-icon iconName="icon-pinglun" size="20px" top="1.5px"></svg-icon>
                        回复
                    </span>
                </div>
            </div>
        </div>
        <div class="secondary-comment-box">
            <CommentInput :isShow="commentInfo.second_comments_status!" @comment="handleComment"></CommentInput>
        </div>
    </div>
</template>

<script lang='ts' setup>
import { computed, ref } from "vue";
import { type IComment } from "./types"
import CommentInput from "./comment-input.vue";

interface Props {
    commentInfo:
    IComment
}
const props = withDefaults(defineProps<Props>(), {
});

const emits = defineEmits([
    "secondaryReview",
    "sendComment"
]);
const secondaryReview = () => {
    emits("secondaryReview")
}
const handleComment = (comment: string) => {
    emits("sendComment", comment, props.commentInfo.comment_id)
}
const userInfo = computed(() => {
    return props.commentInfo.user_info
})
const timeFormat = (time: string) => {
    // 2025-02-19T10:33:20+08:00
    const timeStr = time.split("T")[0]
    return timeStr.slice(2)
}


</script>
<style scoped lang="scss">
.comment-container-box {
    padding: 4px 0;
    display: flex;
    flex-direction: column;

    .first-comment-box {
        display: flex;

        .avatar {
            width: 40px;
            height: 40px;
            border-radius: 50%;
            margin-right: 12px;
        }

        .comment-content {
            display: flex;
            flex-direction: column;
            row-gap: 4px;

            .username {
                font-size: 16px;
                line-height: 24px;
                color: #575c6b;
            }

            .comment-text {
                font-size: 16px;
                font-weight: 400;
                line-height: 28px;
                display: -webkit-box;
                overflow: hidden;
                text-overflow: ellipsis;
                -webkit-box-orient: vertical;
                -webkit-line-clamp: 6;
                color: #262932;
            }

            .comment-brief {
                display: flex;
                color: #575c6b;
                align-items: center;

                .reply {
                    display: flex;
                    gap: 8px;
                    align-items: center;
                    margin-left: 12px;
                    font-size: 14px;
                    line-height: 22px;
                    cursor: pointer;
                }
            }
        }
    }

}

.secondary-comment-box {
    margin: 8px 0 0 52px;
}
</style>