import type { Form } from "@formily/core";

export const publicSchema = (form: Form) => {
  return {
    type: "object",
    properties: {
      category_id: {
        title: "分类",
        type: "string",
        required: true,
        "x-decorator": "FormItem",
        "x-decorator-props": {
          labelWidth: 70,
        },
        "x-component": "RRadio",
        "x-component-props": {
          options: [
            {
              label: "动物科学",
              value: 1,
            },
            {
              label: "水产养殖",
              value: 2,
            },
            {
              label: "动物医学",
              value: 3,
            },
            {
              label: "计算机科学技术",
              value: 4,
            },
            {
              label: "软件工程",
              value: 5,
            },
          ],
        },
      },

      cover: {
        title: "文章封面",
        "x-decorator": "FormItem",
        "x-decorator-props": {
          labelWidth: 70,
        },
        "x-component": "RUpload",
        "x-component-props": {
          desc: "建议尺寸：192*128px (封面仅展示在首页信息流中)",
        },
      },
      topic_id: {
        title: "创作话题",
        "x-decorator-props": {
          labelWidth: 70,
        },
        "x-decorator": "FormItem",
        "x-component": "RSelect",
        "x-component-props": {
          id: "topic_title",
          placeholder: "请添加创作话题，最多添加一个话题",
          options: [
            {
              label: "选项1",
              value: 1,
            },
            {
              label: "选项2",
              value: 2,
            },
            {
              label: "选项3",
              value: 3,
            },
            {
              label: "选项4",
              value: 4,
            },
            {
              label: "选项5",
              value: 5,
            },
            {
              label: "选项6",
              value: 6,
            },
            {
              label: "选项7",
              value: 7,
            },
            {
              label: "选项8",
              value: 8,
            },
          ],
          teleported: false,
        },
      },
      abstract: {
        title: "文章摘要",
        type: "string",
        "x-decorator": "FormItem",
        "x-decorator-props": {
          labelWidth: 70,
        },
        "x-component": "Input",
        "x-component-props": {
          type: "textarea",
          placeholder: "请输入文章摘要",
          maxlength: 100,
          minlength: 20,
          "show-word-limit": true,
          autosize: {
            minRows: 5,
            maxRows: 5,
          },
        },
      },
    },
  };
};
