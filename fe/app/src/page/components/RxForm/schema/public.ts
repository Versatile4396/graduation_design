import type { Form } from "@formily/core";

export const publicSchema = (form: Form) => {
  return {
    type: "object",
    properties: {
      CategoryId: {
        title: "分类",
        type: "string",
        required: true,
        "x-decorator": "FormItem",
        "x-decorator-props": {
          labelWidth: 70,
        },
        "x-component": "CustomRadio",
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

      Cover: {
        title: "文章封面",
        "x-decorator": "FormItem",
        "x-decorator-props": {
          labelWidth: 70,
        },
        "x-component": "CustomUpload",
      },
      topic_title: {
        title: "创作话题",
        "x-decorator-props": {
          labelWidth: 70,
        },
        "x-decorator": "FormItem",
        "x-component": "Input",
        "x-component-props": {
          placeholder: "请输入创作话题",
        },
      },
      abstract: {
        title: "文章摘要",
        type: "string",
        "x-decorator": "FormItem",
        "x-decorator-props": {
          labelWidth: 70,
        },
        "x-component": "Input.TextArea",
        "x-component-props": {
          placeholder: "请输入文章摘要",
          maxLength: 100,
        },
      },
    },
  };
};
