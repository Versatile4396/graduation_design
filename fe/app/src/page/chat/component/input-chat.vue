<template>
  <div class="container">
    <div class="top-icon">
      <!-- emoji图标 -->
      <el-popover
        placement="top-end"
        trigger="click"
        width="450"
        v-model:visible="emojiPopoverStatus"
      >
        <template #reference>
          <div class="icon-box">
            <svg
              t="1690621072710"
              class="icon"
              viewBox="0 0 1024 1024"
              version="1.1"
              xmlns="http://www.w3.org/2000/svg"
              p-id="30103"
              width="25"
              height="25"
            >
              <path
                d="M515.328 965.2224c-249.9072 0-453.2224-203.3152-453.2224-453.2224s203.3152-453.2224 453.2224-453.2224 453.2224 203.3152 453.2224 453.2224-203.3152 453.2224-453.2224 453.2224z m0-845.0048c-216.0128 0-391.7824 175.7696-391.7824 391.7824s175.7696 391.7824 391.7824 391.7824 391.7824-175.7696 391.7824-391.7824-175.7696-391.7824-391.7824-391.7824z"
                fill="#4f4f4f"
                p-id="30105"
              ></path>
              <path
                d="M385.8944 409.4976m-53.6064 0a53.6064 53.6064 0 1 0 107.2128 0 53.6064 53.6064 0 1 0-107.2128 0Z"
                fill="#4f4f4f"
                p-id="30106"
              ></path>
              <path
                d="M636.8768 409.4976m-53.6064 0a53.6064 53.6064 0 1 0 107.2128 0 53.6064 53.6064 0 1 0-107.2128 0Z"
                fill="#4f4f4f"
                p-id="30107"
              ></path>
              <path
                d="M510.208 708.1984c-122.88 0-183.7568-103.1168-186.3168-107.52a30.72 30.72 0 0 1 53.1456-30.8224c1.9968 3.3792 46.4896 76.9024 133.1712 76.9024 86.784 0 131.328-73.7792 133.1712-76.9024a30.78144 30.78144 0 0 1 41.984-11.008 30.6688 30.6688 0 0 1 11.1616 41.8304c-2.56 4.352-63.488 107.52-186.3168 107.52z"
                fill="#4f4f4f"
                p-id="30108"
              ></path>
            </svg>
          </div>
        </template>
        <Emoji @selectEmoji="handleSelectEmoji"></Emoji>
      </el-popover>
      <el-upload
        class="avatar-uploader"
        action="http://127.0.0.1:5555/api/upload/image"
        :show-file-list="false"
        :on-success="handleAvatarSuccess"
        :before-upload="beforeAvatarUpload"
      >
        <svg-icon iconName="icon-shangchuantupian" color="#4F4F4F"></svg-icon>
      </el-upload>
      <el-popover placement="top">
        <template #reference>
          <svg-icon
            @mousedown="startAudio"
            @mouseup="stopAudio"
            icon-name="icon-luyin"
            color="#4F4F4F"
            class="icon-luyin"
          ></svg-icon>
        </template>
        按下开始录音
      </el-popover>
    </div>
    <!-- 自定义输入框 -->
    <div
      id="message-input"
      contenteditable="true"
      spellcheck="false"
      autofocus
      ref="messageInputDom"
      class="message-input"
      placeholder="请输入聊天内容"
    />
    <div class="msg-footer">{{ footerMsg }}</div>
  </div>
</template>

<script setup lang="ts">
import Emoji from './Emoji.vue'
import { computed, onMounted, ref } from 'vue'
import type { UploadProps } from 'element-plus'
import { ElMessage } from 'element-plus'
import SvgIcon from '@/assets/iconfont/SvgIcon.vue'
import Recorder from 'recorder-core'
import 'recorder-core/src/engine/wav.js'

interface Props {
  enterLock?: boolean
  footerMsg?: string
}
const props = withDefaults(defineProps<Props>(), {
  enterLock: false,
  footerMsg: '按 Enter 发送消息'
})

const emits = defineEmits(['sendMsg', 'sendAudio'])

const emojiPopoverStatus = ref(false)

const beforeAvatarUpload: UploadProps['beforeUpload'] = (rawFile) => {
  if (rawFile.type !== 'image/jpeg') {
    ElMessage.error('Avatar picture must be JPG format!')
    return false
  } else if (rawFile.size / 1024 / 1024 > 2) {
    ElMessage.error('Avatar picture size can not exceed 2MB!')
    return false
  }
  return true
}

const handleAvatarSuccess: UploadProps['onSuccess'] = (response) => {
  const imageUrl = response.data.image_url
  const imageNode = `<img src="${imageUrl}" width="50" height="50" style="vertical-align: middle; margin:2px;">`
  handleInsertImage(imageNode)
}

const messageInputDom = ref<HTMLInputElement>()

const handleSendMessage = () => {
  const message = messageInputDom.value?.innerHTML || ''
  if (message.trim() !== '') {
    emits('sendMsg', message)
    if (document.activeElement != messageInputDom.value) {
      messageInputDom.value?.focus()
    }
    messageInputDom.value!.innerHTML = ''
  }
}

const handleInsertImage = (imageUrl: string) => {
  if (document.activeElement != messageInputDom.value) {
    messageInputDom.value?.focus()
    // 将光标移动到输入框的末尾
    const range = document.createRange()
    range.selectNodeContents(messageInputDom.value!)
    range.collapse(false)
    const selection = window.getSelection()
    selection?.removeAllRanges()
    selection?.addRange(range)
  }
  document.execCommand('insertHTML', false, imageUrl)
}

const handleSelectEmoji = (index: number) => {
  // 构建emoji的HTML字符串
  const emojiImg = `<img src="/src/assets/gif/${index}.gif" width="25" height="25" style="vertical-align: middle; margin:2px;">`
  handleInsertImage(emojiImg)
  emojiPopoverStatus.value = false
}
let audiorecorder: any = null
let wave: any
const startAudio = async () => {
  audiorecorder = Recorder({
    type: 'wav', //录音格式，可以换成wav等其他格式
    sampleRate: 16000, //录音的采样率，越大细节越丰富越细腻
    bitRate: 16, //录音的比特率，越大音质越好
    onProcess: (
      buffers: any,
      powerLevel: any,
      _bufferDuration: any,
      bufferSampleRate: any,
      _newBufferIdx: any,
      _asyncEnd: any
    ) => {
      //录音实时回调，大约1秒调用12次本回调
      //可实时绘制波形，实时上传（发送）数据
      if (wave) {
        wave.input(buffers[buffers.length - 1], powerLevel, bufferSampleRate)
      }
    }
  })
  if (!audiorecorder) {
    ElMessage.error('录音失败 当前浏览器不支持录音功能')
    return
  }
  await audiorecorder.open(
    () => {
      audiorecorder.start()
      console.log('录音开始')
    },
    (msg: any, isUserNotAllow: any) => {
      //用户拒绝了录音权限，或者浏览器不支持录音
      console.log((isUserNotAllow ? 'UserNotAllow,' : '') + '无法录音:' + msg)
    }
  )
}
let recBlob: any = null
const stopAudio = () => {
  if (!audiorecorder) {
    console.error('未打开录音')
    return
  }
  audiorecorder.stop((blob: any, _duration: any) => {
    console.log(blob, '====>')
    recBlob = blob
    //简单利用URL生成本地文件地址，此地址只能本地使用，比如赋值给audio.src进行播放，赋值给a.href然后a.click()进行下载（a需提供download="xxx.mp3"属性）
    const localUrl = (window.URL || window.webkitURL).createObjectURL(blob)
    console.log('录音成功====>', blob, localUrl, '时长:' + _duration + 'ms')

    let reader = new FileReader()
    reader.readAsArrayBuffer(blob)
    reader.onload = (e: any) => {
      let imgData = e.target.result
      // 上传文件必须将ArrayBuffer转换为Uint8Array
      let data = {
        content: undefined,
        contentType: 3,
        fileSuffix: 'wav',
        file: new Uint8Array(imgData)
      }
      emits('sendAudio', data)
    }
    audiorecorder.close()
  })
}

onMounted(() => {
  messageInputDom.value?.addEventListener('keydown', (event) => {
    if (event.key === 'Enter') {
      event.preventDefault()
      if (props.enterLock) {
        return
      }
      handleSendMessage()
    }
  })
})
const value = computed(() => {
  return messageInputDom.value?.innerHTML || ''
})

const clearMsg = () => {
  messageInputDom.value!.innerHTML = ''
}

defineExpose({ value, clearMsg })
</script>

<style scoped lang="scss">
.icon-luyin {
  cursor: pointer;
  margin-left: 4px;
  outline: none;
}
.container {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  position: relative;
  padding: 10px 15px;
  box-sizing: border-box;

  .top-icon {
    width: 100%;
    display: flex;

    .icon-box {
      margin-right: 10px;
      cursor: pointer;
    }
  }

  .message-input {
    padding-top: 8px;
    font-size: 14px;
    overflow-y: auto;
    color: #262932;
    flex: 1;
    width: 100%;
    resize: none;
    border: none;
    outline: none;
  }

  .msg-footer {
    bottom: 6px;
    right: 16px;
    height: 24px;
    padding: 0;
    font-size: 14px;
    line-height: 22px;
    color: #c3c8d0;
    text-align: right;
  }
}
</style>
