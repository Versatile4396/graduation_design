import { MessageSchema } from '@/proto/message_pb'
import { create, toBinary, fromBinary } from "@bufbuild/protobuf"
import * as Constant from './constant'
import { nextTick, ref } from 'vue'
import Ajax from '@/ajax'
import type { userInfo } from '@/ajax/type/user'
import type { ChatInstance } from '../type'
import { watch } from 'vue'


export const useWebSocket = (uid: string, chatWrapperDom: any) => {
  const historyMessage = ref<ChatInstance[]>([])
  const currentChat = ref<userInfo>({})
  var peer = new RTCPeerConnection();
  const socket = new WebSocket('ws://localhost:5555/api/chat/ws?uid=' + uid)
  const heartCheck = {
    timeout: 10000, // 10s
    timeoutObj: null,
    serverTimeoutObj: null,
    num: 3,
    start: function () {
      var self = this
      var _num = this.num
      this.timeoutObj && clearTimeout(this.timeoutObj)
      this.serverTimeoutObj && clearTimeout(this.serverTimeoutObj)
      this.timeoutObj = setTimeout(function () {
        //这里发送一个心跳，后端收到后，返回一个心跳消息，
        let data = {
          type: 'heatbeat',
          content: 'ping'
        }
        if (socket.readyState === 1) {
          const messagePB = create(MessageSchema, data)
          socket.send(toBinary(MessageSchema, messagePB))
        }
        self.serverTimeoutObj = setTimeout(function () {
          _num--
          if (_num <= 0) {
            console.log("the ping num is more then 3, close socket!")
            socket.close();
          }
        }, self.timeout) as any;
      }, this.timeout) as any
    }
  }

  const sendMessage = (messageData: any,) => {
    let toUid = messageData.toUid;

    let data = {
      contentType: Constant.TEXT, // 消息类型，1.文本 2.图片 3.文件 4.语音 5.视频 6.位置 7.自定义,
      to: toUid,
      ...messageData,
    }
    const dataReflect = {
      ...data,
      type: Constant.MESSAGE_TRANS_TYPE,
      contentType: data.messageType,
      fromUserId: Number(data.from),
      toUserId: Number(data.to),
    }
    // 发送给服务端的也要再本地存储一份
    historyMessage.value.push(dataReflect)
    const messagePB = create(MessageSchema, data)
    socket.send(toBinary(MessageSchema, messagePB))
  }
  watch(historyMessage, () => {
    // 前端处理滚动条
    nextTick(() => {
      if (chatWrapperDom?.value) {
        chatWrapperDom.value.scrollTo({
          top: chatWrapperDom.value.scrollHeight,
          behavior: 'smooth'
        })
      }
    })
  }, { immediate: true })

  const webrtcConnection = () => {

    /**
     * 对等方收到ice信息后，通过调用 addIceCandidate 将接收的候选者信息传递给浏览器的ICE代理。
     * @param {候选人信息} e 
     */
    peer.onicecandidate = (e) => {
      if (e.candidate) {
        // rtcType参数默认是对端值为answer，如果是发起端，会将值设置为offer
        let candidate = {
          type: 'answer_ice',
          iceCandidate: e.candidate
        }
        let message = {
          content: JSON.stringify(candidate),
          type: Constant.MESSAGE_TRANS_TYPE,
        }
        sendMessage(message);
      }

    };
    const state = {
      onlineType: 1, // 在线视频或者音频： 1视频，2音频
      video: {
        height: 400,
        width: 540
      },
      share: {
        height: 540,
        width: 750
      },
      currentScreen: {
        height: 0,
        width: 0
      },
      videoCallModal: false,
      callName: '',
      fromUserUuid: '',
    }

    /**
     * 当连接成功后，从里面获取语音视频流
     * @param {包含语音视频流} e 
     */
    peer.ontrack = (e) => {
      if (e && e.streams) {
        if (state.onlineType === 1) {
          let remoteVideo = document.getElementById("remoteVideoReceiver")!;
          //@ts-ignore
          remoteVideo.srcObject = e.streams[0];
        } else {
          let remoteAudio = document.getElementById("audioPhone")!;
          //@ts-ignore
          remoteAudio.srcObject = e.streams[0];
        }
      }
    };
  }

  socket.onopen = () => {
    heartCheck.start()
    console.log("connected")
    webrtcConnection()
    // this.props.setSocket(socket);
  }

  socket.onmessage = (message) => {
    heartCheck.start()

    let reader = new FileReader();
    reader.readAsArrayBuffer(message.data);
    reader.onload = (e) => {
      // @ts-ignore
      const messagePB = fromBinary(MessageSchema, new Uint8Array(e.target?.result))
      if (messagePB.type === "heatbeat" || messagePB.from === "System") {
        return;
      }
      const reflectMessage = {
        ...messagePB,
        type: Constant.MESSAGE_TRANS_TYPE,
        contentType: messagePB.messageType,
        fromUserId: Number(messagePB.from),
        toUserId: Number(messagePB.to),
      }
      //@ts-ignore
      historyMessage.value.push(reflectMessage)
      nextTick(() => {
        if (chatWrapperDom?.value) {
          chatWrapperDom.value.scrollTo({
            top: chatWrapperDom.value.scrollHeight,
            behavior: 'smooth'
          })
        }
      })
    }
  }
  // 拿历史消息和用户信息
  const getHistoryMessage = async (uid: string, toUid: string) => {
    // 获取历史消息和用户信息
    const { data } = await Ajax.get("/chat/message", { uid, toUid, messageType: 1 })
    historyMessage.value = data?.messageInfo || [];
    currentChat.value = data?.toUserInfo || {};
  }

  const chatList = ref<userInfo[]>([])
  const getChatList = async () => {
    const { data } = await Ajax.get('/chat/friend', { uid })
    chatList.value = data || []
  }
  return {
    chatList,
    getChatList,
    socket,
    heartCheck,
    historyMessage,
    currentChat,
    sendMessage,
    getHistoryMessage,
  }
}