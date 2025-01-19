export const loginSchema = {
  type: "object",
  properties: {
    account: {
      type: "string",
      required: true,
      "x-component": "Input",
      "x-component-props": {
        placeholder: "请输入手机号/邮箱号",
      },
      "x-validator": [
        {
          minimum: 5,
        },
      ],
    },
    password: {
      type: "string",
      required: true,
      "x-component": "InputPassword",
      "x-component-props": {
        style: {
          marginTop: "16px",
        },
        placeholder: "输入账号密码",
      },
    },
  },
};

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
      "x-component": "InputPassword",
      "x-component-props": {
        style: {
          marginTop: "16px",
        },
        placeholder: "输入密码",
      },
    },
    confirm_password: {
      type: "string",
      "x-component": "InputPassword",
      "x-component-props": {
        style: {
          marginTop: "16px",
        },
        placeholder: "请再次输入密码",
      },
    },
  },
};
