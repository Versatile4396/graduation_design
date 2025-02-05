<template>
  <div class="editor-style">
    <div class="editor-nav">
      <input
        placeholder="输入文章标题..."
        spellcheck="false"
        maxlength="80"
        class="title-input"
      />
      <div class="operator">
        <span>
          <el-button style="width: 100px" @click="saveDraft"
            >存入草稿箱</el-button
          >
        </span>
        <span>
          <el-button style="width: 100px" type="primary" @click="publicArticle"
            >发布</el-button
          >
        </span>
        <!-- <span>
          <el-popover :width="200" trigger="hover" content="切换成富文本编辑器">
            <template #reference>
              <Switch class="switch-icon"></Switch>
            </template>
          </el-popover>
        </span> -->
        <span>
          <el-avatar class="avatar" />
        </span>
      </div>
    </div>
    <div class="editor-container">
      <Toolbar
        style="border-bottom: 0.5px solid #ccc; width: 1400px"
        :editor="editorRef"
        :defaultConfig="toolbarConfig"
        :mode="mode"
      />
      <Editor
        v-model="valueHtml"
        style="min-height: 800px; overflow-y: hidden; width: 1400px"
        :defaultConfig="editorConfig"
        :mode="mode"
        @onCreated="handleCreated"
      />
    </div>
  </div>
</template>

<script lang="ts" setup>
import "@wangeditor/editor/dist/css/style.css"; // 引入 css
import { Boot } from "@wangeditor/editor";
import markdownModule from "@wangeditor/plugin-md";
import { Editor, Toolbar } from "@wangeditor/editor-for-vue";
import { onBeforeUnmount, ref, shallowRef, onMounted } from "vue";
import { Switch } from "@element-plus/icons-vue";
import { getUrlQuery } from "@/utils/common";
import { FormType } from "@/ajax/type/ariticle";

Boot.registerModule(markdownModule);
// 编辑器实例，必须用 shallowRef
const editorRef = shallowRef();

// 内容 HTML
const valueHtml = ref("");

const mode = "default";
const formType = Number(getUrlQuery().scene);
// 模拟 ajax 异步获取内容
onMounted(() => {
  if (formType === FormType.create) {
  } else if (formType === FormType.editor) {
    // 拉取文章内容
  }
});

const toolbarConfig = {};
const editorConfig = { placeholder: "请输入内容..." };

// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
  const editor = editorRef.value;
  if (editor == null) return;
  editor.destroy();
});

const handleCreated = (editor: any) => {
  editorRef.value = editor; // 记录 editor 实例，重要！
};

const publicArticle = () => {
  console.log("sads");
};
const saveDraft = () => {};
</script>
<style lang="scss" scoped>
.editor-style {
  margin: auto 0px;
  background-color: #fff;
  display: flex;
  flex-direction: column;
  align-items: center;
  min-height: 100vh;
  .editor-nav {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 60px;
    width: 1400px;
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
    flex: 1;
    height: 100%;
  }
}
</style>
