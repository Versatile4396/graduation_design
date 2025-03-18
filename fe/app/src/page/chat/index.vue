<template>
  <div class="chat-container-box">
    <div class="side-bar">
      <div class="search-bar"></div>
      <div class="list-bar"></div>
    </div>
    <div class="chat-container">
      <div class="chat-header"></div>
      <div class="chat-content"></div>
      <div class="chat-send-wrapper">
        <inputChat @send-msg="handleSendMsg"></inputChat>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { getUrlQuery } from "@/utils/common";
import { objectEntries } from "@vueuse/core";
import inputChat from "./component/input-chat.vue";
import { onUnmounted } from "vue";
import Ajax from "@/ajax";

// 获取聊天列表用户信息？
const { uid, toUid } = getUrlQuery();
const query = {
  uid,
  toUid,
};
//@ts-ignore
const queryStr = objectEntries(query)
  .map(
    ([key, value]) => encodeURIComponent(key) + "=" + encodeURIComponent(value)
  )
  .join("&");
// 建立ws链接 发起聊天
const socket = new WebSocket("ws://localhost:5555/api/chat/ws?" + queryStr);

// 连接成功 获取历史消息
socket.onopen = () => {
  // 获取历史消息
  getHistoryMsg();
};

const handleSendMsg = (msg: string) => {
  // 判断socket是否在链接中
  if (socket.readyState === WebSocket.OPEN) {
    const formatMsg = {
      type: 1,
      content: msg,
    };
    socket.send(JSON.stringify(formatMsg));
  }
};

const getHistoryUnreadMsg = () => {
  // 获取历史未读消息
  if (socket.readyState === WebSocket.OPEN) {
    const formatMsg = {
      type: 3,
      content: "",
    };
    socket.send(JSON.stringify(formatMsg));
  }
};
const getHistoryMsg = () => {
  // 获取历史消息
  Ajax.get(`/chat/history?uid=${uid}`).then((res) => {
  });
};

onUnmounted(() => {
  socket.close();
});
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
