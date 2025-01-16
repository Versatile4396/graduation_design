export const loginSchema = {
  type: "object",
  properties: {
    input: {
      type: "string",
      "x-component": "Input",
      "x-component-props": {
        placeholder: "请输入手机号/邮箱号",
      },
    },
  },
};
