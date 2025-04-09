import { MessageSchema } from '@/proto/message_pb'
import { create, toBinary } from "@bufbuild/protobuf"
import * as Constant from './constant'


export const useWebSocket = (uid: string) => {
  var peer = new RTCPeerConnection();
  var lockConnection = false;
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
        //onmessage拿到返回的心跳就说明连接正常
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

  const sendMessage = (messageData: any, propsToUid?: string, propsMessageType?: number) => {
    let toUid = messageData.toUid;
    if (null == toUid) {
      toUid = propsToUid;
    }
    let data = {
      ...messageData,
      messageType: propsMessageType, // 消息类型，1.单聊 2.群聊
      fromUsername: localStorage.username,
      from: localStorage.uuid,
      to: toUid,
    }
    const messagePB = create(MessageSchema, data)
    socket.send(toBinary(MessageSchema, messagePB))
  }

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

  return {
    socket,
    heartCheck
  }
}
