// import qs from "query-string";
import { useLocalStorage } from "@vueuse/core";
const userLocalInfo = useLocalStorage<string>("userinfo", "");
const token = useLocalStorage<string>("token", "");
export type Obj = Record<string, any>;
import { getAllLabType } from "../api/methods/labTypeReq";
import { labQueryAll } from "@/api/methods/labReq";
const labTypeReflect = useLocalStorage<string>("labTypeinfo", "");
const labReflect = useLocalStorage<string>("labinfo", "");

export const getUrlQuery = (url: string = window.location.href) => {
  const paramArr = url.slice(url.indexOf("?") + 1).split("&");
  const params: any = {};
  paramArr.map((param: any) => {
    const [key, val] = param.split("=");
    params[key] = decodeURIComponent(val);
  });
  return params;
};
export const setToken = (tokenStr: string) => {
  token.value = tokenStr;
};
export const setUserInfo = (userStr: string) => {
  userLocalInfo.value = userStr;
};
export const clearLocalInfo = () => {
  userLocalInfo.value = "";
  token.value = "";
};
// 获取所有的映射关系
export const getReflect = async () => {
  const res = await getAllLabType();
  const res1 = await labQueryAll();
  labTypeReflect.value = JSON.stringify(res.data);
  labReflect.value = JSON.stringify(res1.data);
};
