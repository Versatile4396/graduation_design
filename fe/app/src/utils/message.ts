// import { ExclamationCircleOutlined } from "@ant-design/icons-vue";
// import { message, Modal } from "ant-design-vue";
// import { createVNode } from "vue";

// type MessageType =
//   | "error"
//   | "success"
//   | "danger"
//   | "info"
//   | "warning"
//   | "loading";

// type Msg = string;

// export const Message = (
//   type: MessageType,
//   msg: Msg,
//   callback?: (instance: any) => unknown,
//   duration = 2
// ) => {
//   const instance = message[type](msg, duration);
//   setTimeout(() => callback && callback(instance), duration);
//   return instance;
// };
// export const ModalMessage = (
//   msg: Msg,
//   title: string = "提示",
//   onOkCallBack?: Function,
//   okParams?: any,
//   onCancelCallBack?: Function,
//   cancelParams?: any
// ) => {
//   Modal.confirm({
//     title: title,
//     icon: createVNode(ExclamationCircleOutlined),
//     content: createVNode("div", { style: "color:red;" }, msg),
//     onOk() {
//       onOkCallBack && onOkCallBack(okParams);
//     },
//     onCancel() {
//       onCancelCallBack && onCancelCallBack(cancelParams);
//     },
//   });
// };
