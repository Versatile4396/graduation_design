<template>
  <div class="container">
    <div class="editor-style">
      <div class="editor-nav">
        <input placeholder="输入文章标题..." v-model="aTitle" spellcheck="false" maxlength="80" class="title-input" />
        <div class="operator">
          <span>
            <el-button style="width: 100px" @click="saveDraft">存入草稿箱</el-button>
          </span>
          <span>
            <el-popover v-model:visible="popoverStatus" trigger="click" placement="bottom" :width="580"
              popper-style="box-shadow: rgb(14 18 22 / 35%) 0px 10px 38px -10px, rgb(14 18 22 / 20%) 0px 10px 20px -15px; padding: 20px;">
              <template #reference>
                <el-button style="width: 100px" type="primary">发布</el-button>
              </template>
              <PublicForm @submit="submitHandle"></PublicForm>
            </el-popover>
          </span>
          <!-- <span>
          <el-popover :width="200" trigger="hover" content="切换成富文本编辑器">
            <template #reference>
              <Switch class="switch-icon"></Switch>
            </template>
          </el-popover>
        </span> -->
          <span>
            <Avatar></Avatar>
          </span>
        </div>
      </div>
      <Toolbar style="border-bottom: 0.5px solid #ccc;" :editor="editorRef" :defaultConfig="toolbarConfig"
        :mode="mode" />
      <div class="editor-container">
        <Editor v-model="aContent" :defaultConfig="editorConfig" :mode="mode" @onCreated="handleCreated" />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import "@wangeditor/editor/dist/css/style.css"; // 引入 css
import { Boot } from "@wangeditor/editor";
import markdownModule from "@wangeditor/plugin-md";
import { Editor, Toolbar } from "@wangeditor/editor-for-vue";
import { onBeforeUnmount, ref, shallowRef, onMounted } from "vue";
import PublicForm from "./public-form.vue";
import Avatar from "../components/navigator/components/avatar.vue";
import { Switch } from "@element-plus/icons-vue";
import { getUrlQuery } from "@/utils/common";
import { FormType } from "@/ajax/type/ariticle";
import Ajax from "@/ajax";
import { Message } from "@/utils/message";
import { userInfoStore } from "@/store/user"
Boot.registerModule(markdownModule);
// 编辑器实例，必须用 shallowRef
const editorRef = shallowRef();
const { userInfo } = userInfoStore()

// 内容 HTML
const aContent = ref("");
const aTitle = ref("");

const mode = "default";
const formType = Number(getUrlQuery().scene);
// 模拟 ajax 异步获取内容
onMounted(() => {
  // 
  const editorContainer = document.querySelector(".editor-container") as any;
  if (editorContainer) {
    editorContainer.style.height = "calc(100vh - 120px)";
  }

  if (formType === FormType.create) {
  } else if (formType === FormType.editor) {
    // 拉取文章内容
  }
});

type InsertFnType = (url: string, poster: string, href: string) => void
const toolbarConfig = {};
const editorConfig = {
  placeholder: "请输入内容...", MENU_CONF: {
    uploadImage: {
      server: 'http://127.0.0.1:5555/api/upload/image',
      fieldName: 'file',
      // 单个文件的最大体积限制，默认为 2M
      maxFileSize: 1 * 1024 * 1024, // 1M
      // 最多可上传几个文件，默认为 100
      maxNumberOfFiles: 10,
      // 选择文件时的类型限制，默认为 ['image/*'] 。如不想限制，则设置为 []
      allowedFileTypes: ['image/*'],
      // 超时时间，默认为 10 秒
      timeout: 5 * 1000, // 5 秒
      customInsert(res: any, insertFn: InsertFnType) {
        insertFn(res.data.image_url, res.data.image_name, "")
      }
    },
  }
} as any;


// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
  const editor = editorRef.value;
  if (editor == null) return;
  editor.destroy();
});

const handleCreated = (editor: any) => {
  editorRef.value = editor; // 记录 editor 实例，重要！
};


// 发布
const popoverStatus = ref(false)
const submitHandle = (value: any) => {
  if (aContent.value === "") {
    Message.info("请输入文章内容");
    return;
  }
  if (aTitle.value === "") {
    Message.info("请输入文章标题");
    return;
  }
  const params = {
    ...value,
    tag_id: 1,
    topic_id: 1,
    content: aContent.value,
    title: aTitle.value,
    user_id: Number(getUrlQuery().uid)
  }
  Ajax.post("/article/create", params).then((_) => {
    popoverStatus.value = false;
  })
};


const saveDraft = () => { };
</script>
<style lang="scss" scoped>
.container {
  width: 1370px;
  margin: 0 auto;

  .editor-style {
    margin: auto 0;
    display: flex;
    flex-direction: column;
    align-items: center;

    .editor-nav {
      width: 100%;
      display: flex;
      justify-content: center;
      align-items: center;
      height: 60px;
      background-color: #fff;

      .title-input {
        margin: 0;
        padding: 0;
        font-size: 24px;
        font-weight: 500;
        color: #1d2129;
        border: none;
        outline: none;
        height: 100%;
        width: 100%;
        padding-left: 20px;
      }

      .title-input:focus {
        outline: none;
        border: none;
      }

      .operator {
        span {
          display: flex;
          align-items: center;
          margin-left: 12px;
        }

        display: flex;
        align-content: center;
        justify-self: center;
        padding-right: 12px;

        .switch-icon {
          margin-left: 12px;
          height: 20px;
          width: 20px;
          color: gray;
          cursor: pointer;
        }

        .switch-icon:hover {
          background-color: #dddee4;
        }

        .avatar {
          margin-left: 12px;
        }
      }
    }

    .editor-container {
      overflow: hidden;
      width: 100%;
      min-height: calc(100vh - 101px);
      background-color: #fff;
    }
  }
}
</style>
