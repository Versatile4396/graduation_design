import { ElMessage } from "element-plus";

class MessageClass {
  info(message: string) {
    ElMessage({
      message,
      type: "info",
    });
  }
  success(message: string) {
    ElMessage({
      message,
      type: "success",
    });
  }
  err(message: string) {
    ElMessage({
      message,
      type: "error",
    });
  }
  warn(message: string) {
    ElMessage({
      message,
      type: "warning",
    });
  }
}

const Message = new MessageClass();

export { Message };
