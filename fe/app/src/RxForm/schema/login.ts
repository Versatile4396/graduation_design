import type { ISchema } from "@formily/vue";

export const loginSchema = {
  type: "object",
  properties: {
    account: {
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
        style: {
          marginTop: "16px",
        },
        placeholder: "输入账号密码",
      },
    },
  },
} as ISchema;

export const registerSchema = {
  type: "object",
  properties: {
    account: {
      type: "string",
      required: true,
      "x-component": "Input",
      "x-component-props": {
        placeholder: "请输入手机号/邮箱号",
      },
      "x-validator": `{{(value)=> {
        if(value.length > 18) {
          return '密码长度不能超过18';
        }
        if(value.length < 4) {
          return '密码长度至少为4个字符';
        }
      }}}`,
    },
    password: {
      required: true,
      type: "string",
      "x-component": "Password",
      "x-component-props": {
        style: {
          marginTop: "16px",
        },
        placeholder: "输入密码",
      },
    },
    confirm_password: {
      type: "string",
      "x-component": "Password",
      "x-component-props": {
        style: {
          marginTop: "16px",
        },
        placeholder: "请再次输入密码",
      },
    },
  },
};
