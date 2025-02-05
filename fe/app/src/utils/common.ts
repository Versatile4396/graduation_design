// import qs from "query-string";
import type { userInfo } from "@/ajax/type/user";
import { useLocalStorage } from "@vueuse/core";
import qs from "query-string";

const userLocalInfo = useLocalStorage<Partial<userInfo>>("userinfo", {});
const acccess_token = useLocalStorage<string>("acccess_token", "");
const refresh_token = useLocalStorage<string>("refresh_token", "");

export type Obj = Record<string, any>;

export const getUrlQuery = () => {
  const query = qs.parse(window.location.search);
  if (query.uid === undefined) {
    query.uid = "";
  }
  return query;
};

export const setUrlQuery = (key: string, value: any) => {
  const url = new URL(window.location.href);
  const searchParams = url.searchParams;
  // 设置指定的参数键值对
  searchParams.set(key, value);
  // 更新URL（不刷新页面，仅改变当前历史记录项的URL）
  window.history.replaceState({}, "", url);
  return url.href;
};

export const setToken = (
  aToken: string | undefined,
  rToken: string | undefined
) => {
  acccess_token.value = aToken;
  refresh_token.value = rToken;
};
export const setLocalStorageUserInfo = (uinfo: userInfo) => {
  // 设置query参数
  setUrlQuery("uid", "");
  userLocalInfo.value = uinfo;
};
export const clearLocalInfo = () => {
  userLocalInfo.value = {};
  acccess_token.value = "";
  refresh_token.value = "";
};

export { userLocalInfo };
