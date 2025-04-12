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

    let data = {
      from: uid,
      to: String(currentChat.value.user_id),
      ...messageData,
    }

    // 发送给服务端的也要再本地存储一份
    const messagePB = create(MessageSchema, data)
    socket.send(toBinary(MessageSchema, messagePB))

    const getDataReflectValue = (data: any) => {
      if (data.contentType === 3) {
        data.contentType = GetContentTypeBySuffix(data.fileSuffix || "")
      }
      return {
        ...data,
        fromUserId: Number(data.from),
        toUserId: Number(data.to),
      }
    }
    const dataReflect = getDataReflectValue(data)
    historyMessage.value.push(dataReflect)
    // 前端处理滚动条
    nextTick(() => {
      if (chatWrapperDom.value) {
        chatWrapperDom.value.scrollTo({
          top: chatWrapperDom.value.scrollHeight,
          behavior: 'smooth'
        })
      }
    })
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

  const localPeer = new RTCPeerConnection()
  const remoteAudioStream = ref(null)
  /**
   * 语音通话模块
   * webrtc 绑定事件
   */
  const webrtcConnection = () => {
    /**
    * 对等方收到ice信息后，通过调用 addIceCandidate 将接收的候选者信息传递给浏览器的ICE代理。
    * @param {候选人信息} e 
    */
    localPeer.onicecandidate = (e) => {
      if (e.candidate) {
        // rtcType参数默认是对端值为answer，如果是发起端，会将值设置为offer
        let candidate = {
          type: 'offer_ice',
          iceCandidate: e.candidate
        }
        let message = {
          content: JSON.stringify(candidate),
          type: Constant.MESSAGE_TRANS_TYPE,
        }
        sendMessage(message)
      }
    }
    /**
     * 当连接成功后，从里面获取语音视频流
     * @param {包含语音视频流} e
     */
    localPeer.ontrack = (e) => {
      if (e && e.streams) {
        let remoteAudio = document.getElementById("remoteAudioPhone");
        //@ts-ignore
        remoteAudio!.srcObject = e.streams[0];
      }
    };
  }

  /**
   * 视频通话模块
   */

  return {
    chatList,
    getChatList,
    socket,
    heartCheck,
    historyMessage,
    currentChat,
    sendMessage,
    getHistoryMessage,
    // 音视频通话模块
    webrtcConnection,
    localPeer,
  }
}

export const GetContentTypeBySuffix = (suffix: string) => {
  if (['png', 'jpg', 'jpeg', 'gif'].includes(suffix)) {
    return Constant.ContentTypeEnum.IMAGE
  }
  if (['mp4', 'avi', 'mov'].includes(suffix)) {
    return Constant.ContentTypeEnum.VIDEO
  }
  if (['mp3', 'wav'].includes(suffix)) {
    return Constant.ContentTypeEnum.AUDIO
  }
  return Constant.ContentTypeEnum.FILE
}