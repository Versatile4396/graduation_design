<template>
  <el-upload
    :style="{ '--custom-width': width, '--custom-height': height }"
    class="avatar-uploader"
    action="http://127.0.0.1:5555/api/upload/image"
    :show-file-list="false"
    :on-success="handleAvatarSuccess"
    :before-upload="beforeAvatarUpload"
    v-if="!props.multiple"
  >
    <img v-if="imageUrl" :src="imageUrl" class="avatar" />
    <el-icon v-else class="avatar-uploader-icon">
      <Plus />
    </el-icon>
  </el-upload>
  <el-upload
    :style="{ '--custom-width': width, '--custom-height': height }"
    class="avatar-uploader"
    action="http://127.0.0.1:5555/api/upload/image"
    v-model:file-list="imageUrlList"
    list-type="picture-card"
    :before-upload="beforeAvatarUpload"
    :on-preview="handlePictureCardPreview"
    :on-change="handleImageListChange"
    v-else
  >
    <img v-if="imageUrl" :src="imageUrl" class="avatar" />
    <el-icon v-else class="avatar-uploader-icon">
      <Plus />
    </el-icon>
  </el-upload>
  <el-dialog v-model="dialogVisible">
    <img w-full :src="dialogImageUrl" alt="Preview Image" />
  </el-dialog>
  <span v-if="!!props.desc">{{ props.desc }}</span>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import type { UploadProps } from 'element-plus'

interface Props {
  url?: string
  desc?: string
  value?: string
  width?: string
  height?: string
  borderRadius?: string
  multiple?: boolean
}
const props = withDefaults(defineProps<Props>(), {
  width: '192px',
  height: '128px',
  multiple: false
})

const emits = defineEmits(['change'])
const imageUrl = ref('')
if (props.value) {
  imageUrl.value = props.value!
  emits('change', props.value)
}
const handleAvatarSuccess: UploadProps['onSuccess'] = (response) => {
  imageUrl.value = response.data.image_url
  emits('change', response.data.image_url)
}
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

const imageUrlList = ref([])
const dialogImageUrl = ref()
const dialogVisible = ref(false)
const handleImageListChange = () => {
  const emitsImageList = imageUrlList.value?.map((img: any) => {
    return img.response.data.image_url
  })
  emits('change', JSON.stringify(emitsImageList))
}
const handlePictureCardPreview: UploadProps['onPreview'] = (uploadFile) => {
  dialogImageUrl.value = uploadFile.url!
  dialogVisible.value = true
}
</script>

<style scoped>
.avatar-uploader .avatar {
  width: var(--custom-width);
  height: var(--custom-height);
}
</style>

<style>
.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.desc-text {
  margin-top: 10px;
  font-size: 12px;
  color: #8c939d;
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: var(--custom-width);
  height: var(--custom-height);
  text-align: center;
}
</style>
