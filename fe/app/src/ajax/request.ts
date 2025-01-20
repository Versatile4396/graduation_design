import axios from "axios";
// import { tokenExpire } from "./helper";
import { logB, logN } from "@/utils/log";
// import { Message } from "@/utils/message";

// 处理  类型“AxiosResponse<any, any>”上不存在属性“...”。ts(2339) 脑壳疼！关键一步。
declare module "axios" {
  interface AxiosResponse<T = any> {
    code: Number;
    data: T;
    // 这里追加你的参数
  }
  export function create(config?: AxiosRequestConfig): AxiosInstance;
}

var instance = axios.create({
  baseURL: "http://localhost:5555/api",
  timeout: 50000, //超时时间
  headers: {
    "Content-Type": "application/json;charset=UTF-8",
  },
});

let loading = false;
let requestCount = 0;
// loading效果
const showLoading = () => {
  if (requestCount === 0 && !loading) {
  }
};
// 隐藏loading效果
const hideLoading = () => {};
// 请求拦截器
instance.interceptors.request.use(
  (config) => {
    showLoading();
    logN.success("请求URL", config.url!, "params:", config.params);
    const token = window.localStorage.getItem("token");
    if (token) {
      config.headers.token = token;
    }
    if (config.method === "POST") {
      config.data = JSON.stringify(config.data);
    }
    return config;
  },
  (error) => Promise.reject(error)
);
// 响应拦截器
instance.interceptors.response.use(
  (response) => {
    logB.success("响应数据：", response.data);
    hideLoading();
    if (response.data.code !== 1000) {
      // Message("error", response.data?.message || "未知错误");
    }
    return response.data;
  },
  (error) => {
    // 返回401 token权限没有通过
    if (error.response.status === 401) {
      // tokenExpire();
    }
  }
);
export default instance;
