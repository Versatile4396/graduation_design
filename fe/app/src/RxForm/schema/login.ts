import type { ISchema } from "@formily/vue";

export const loginSchema = {
  type: "object",
  properties: {
    username: {
      type: "string",
      "x-decorator": "FormItem",
      "x-component": "Input",
      "x-component-props": {
        placeholder: "请输入手机号/邮箱号",
      },
      "x-validator": {
        required: true,
        minLength: 4,
      },
    },
    password: {
      type: "string",
      "x-validator": {
        required: true,
        minLength: 6,
      },
      "x-decorator": "FormItem",
      "x-component": "Password",
      "x-component-props": {
        placeholder: "输入账号密码",
      },
    },
  },
} as ISchema;

export const registerSchema = {
  type: "object",
  properties: {
    username: {
      type: "string",
      required: true,
      "x-decorator": "FormItem",
      "x-component": "Input",
      "x-component-props": {
        placeholder: "输入用户名",
      },
      "x-validator": {
        maxLength: 16,
      },
    },
    email: {
      type: "string",
      "x-decorator": "FormItem",
      "x-component": "Input",
      "x-component-props": {
        placeholder: "输入你的邮箱",
      },
      "x-validator": {
        required: true,
        format: "email",
      },
    },

    password: {
      type: "string",
      "x-validator": {
        required: true,
        minLength: 6,
      },
      "x-decorator": "FormItem",
      "x-component": "Password",
      "x-component-props": {
        placeholder: "输入密码",
      },
    },
    confirm_password: {
      type: "string",
      "x-validator": {
        required: true,
        minLength: 6,
      },
      "x-decorator": "FormItem",
      "x-component": "Password",
      "x-component-props": {
        placeholder: "再次输入密码",
      },
    },
    gender: {
      type: "number",
      title: "性别",
      "x-decorator": "FormItem",
      "x-component": "Radio.Group",
      "x-component-props": {
        default: 2,
      },
      initialValue: 2,
      enum: [
        {
          label: "男",
          value: 0,
        },
        {
          label: "女",
          value: 1,
        },
        {
          label: "武装直升机",
          value: 2,
        },
        {
          label: "胖东来洗发水",
          value: 3,
        },
      ],
    },
  },
};
