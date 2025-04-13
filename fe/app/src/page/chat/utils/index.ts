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
    console.log(data, "sendMessage ===message")
    const messagePB = create(MessageSchema, data)
    socket.send(toBinary(MessageSchema, messagePB))
    // webrtc 内容不需要再本地存储
    if (messageData.type === Constant.MESSAGE_TRANS_TYPE) {
      return
    }
    // 发送给服务端的也要再本地存储一份
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
      console.log(messagePB, "onMessage ===message")
      if (messagePB.type === Constant.MESSAGE_TRANS_TYPE) {
        try {
          // 处理webrtc
          handleWebrtcOnMessage(messagePB)
          return
        } catch (error) {

        }
      }
      const reflectMessage = {
        ...messagePB,
        contentType: messagePB.contentType,
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
  const remotePeer = new RTCPeerConnection()
  const chatModalStatus = ref(false)
  const caller = ref(false)
  const called = ref(false)
  const calling = ref(false)
  const peer = ref()
  const localStream = ref()
  const handleWebrtcOnMessage = async (messagePB: any) => {
    if (
      messagePB.contentType >= Constant.DIAL_MEDIA_START &&
      messagePB.contentType <= Constant.DIAL_MEDIA_END
    ) {
      dealMediaCall(messagePB);
      return;
    }
    const { type } = JSON.parse(messagePB.content)
    if (type === 'offer') {
      // 用户B需要创建自己的RTCPeerConnection，
      // 添加本地音视频流，设置远端描述信息，生成answer，并且通过信令服务器发送给用户A
      const offer = JSON.parse(messagePB.content);
      peer.value = new RTCPeerConnection()
      const stream = await getLocalStream(true)
      peer.value.addStream(stream)
      // 设置远端描述信息
      await peer.value.setRemoteDescription(offer);
      const answer = await peer.value.createAnswer()
      // 发送answer给信令服务器
      await peer.value.setLocalDescription(answer);
      sendMessage({
        type: Constant.MESSAGE_TRANS_TYPE,
        content: JSON.stringify(answer)
      })
    } else if (type === "answer") {
      const answer = JSON.parse(messagePB.content);
      peer.value.setRemoteDescription(answer)
      // 通过监听onicecandidate事件获取candidate信息
      peer.value.onicecandidate = (event: any) => {
        if (event.candidate) {
          // 通过信令服务器发送candidate信息给用户B
          sendMessage({
            type: Constant.MESSAGE_TRANS_TYPE,
            content: JSON.stringify(event.candidate),
          })
        }
      }
    } else if (type === "candidate") {
      const candidate = JSON.parse(messagePB.content)
      peer.value.addIceCandidate(candidate);
      peer.value.onicecandidate = (event: any) => {
        if (event.candidate) {
          // 通过信令服务器发送candidate信息给用户A
          sendMessage({
            type: Constant.MESSAGE_TRANS_TYPE,
            content: JSON.stringify(event.candidate)
          })
        }
      }
    } else {
      calling.value = true
      const candidate = JSON.parse(messagePB.content)
      console.log(candidate, "candidate")
      await peer.value.addIceCandidate(candidate);
      peer.value.ontrack = (event: any) => {
        const remoteVideo = document.getElementById("remoteVideo") as HTMLVideoElement
        remoteVideo!.srcObject = event.stream
        remoteVideo!.play()
      }
    }
  }

  const dealMediaCall = async (message: any) => {
    if ([Constant.DIAL_AUDIO_ONLINE, Constant.DIAL_VIDEO_ONLINE].includes(message.contentType)) {
      chatModalStatus.value = true
      called.value = true
    }
    if (
      [Constant.CANCELL_AUDIO_ONLINE,
      Constant.CANCELL_VIDEO_ONLINE].includes(message.contentType)
    ) {
      chatModalStatus.value = false
      stopVideoOnline()
      return;
    }
    if (
      [Constant.REJECT_AUDIO_ONLINE,
      Constant.REJECT_VIDEO_ONLINE].includes(message.contentType)
    ) {
      chatModalStatus.value = false
      stopVideoOnline()
      return;
    }

    if (
      [Constant.ACCEPT_VIDEO_ONLINE,
      Constant.ACCEPT_AUDIO_ONLINE].includes(message.contentType)
    ) {
      // caller 创建RTCPeerConnection 添加本地音视频流，生成offer，并且通过信令服务器将offer发送给用户B
      // 创建RTCPeerConnection
      peer.value = new RTCPeerConnection()
      // 添加本地音视频流
      peer.value.addStream(localStream.value)
      // 生成offer
      const offer = await peer.value.createOffer({
        offerToReceiveAudio: 1,
        offerToReceiveVideo: 1
      })
      // 设置本地描述的offer
      await peer.value.setLocalDescription(offer);
      // 通过信令服务器将offer发送给用户B
      sendMessage({
        type: Constant.MESSAGE_TRANS_TYPE,
        content: JSON.stringify(offer),
      })
      return;
    }
  };

  // 获取本地音视频流
  const getLocalStream = async (local: boolean = false) => {
    const localVideoDom = document.getElementById("localVideo") as HTMLVideoElement
    if (local) {
      const videoElement = document.createElement('video');
      videoElement.src = "/src/page/chat/utils/local.mp4"
      videoElement.muted = true;
      videoElement.autoplay = true;

      // 等待视频加载并播放
      await new Promise((resolve, reject) => {
        videoElement.onloadedmetadata = () => {
          videoElement.play().then(() => {
            resolve(true);
          }).catch((error) => {
            reject(error);
          });
        };
        videoElement.onerror = () => {
          reject(new Error('视频加载失败'));
        };
      });

      // 使用 MediaStream API 从视频元素创建媒体流
      // @ts-ignore
      const mediaStream = videoElement.captureStream();
      return mediaStream;

    }
    const stream = await navigator.mediaDevices.getUserMedia({ // 获取音视频流
      audio: true,
      video: true
    })
    localStream.value = stream
    if (localVideoDom) {
      localVideoDom!.srcObject = stream
      localVideoDom!.play()
    }
    return stream
  }
  const stopVideoOnline = () => {
    let preview = document.getElementById("localVideo") as HTMLVideoElement;
    //@ts-ignore
    if (preview && preview.srcObject && preview.srcObject.getTracks()) {
      //@ts-ignore
      preview.srcObject.getTracks().forEach((track) => track.stop());
    }

    //@ts-ignore
    let remoteVideo = document.getElementById("remoteVideo") as HTMLVideoElement;
    //@ts-ignore
    if (remoteVideo && remoteVideo.srcObject && remoteVideo.srcObject.getTracks()) {
      //@ts-ignore
      remoteVideo.srcObject.getTracks().forEach((track) => track.stop());
    }
    localStream.value?.getTracks().forEach((track: any) => {
      track.stop()
    })
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
    // 音视频通话模块
    called,
    caller,
    calling,
    localPeer,
    remotePeer,
    chatModalStatus,
    getLocalStream,
    stopVideoOnline
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
