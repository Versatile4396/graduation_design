// import qs from "query-string";
import type { userInfo } from "@/ajax/type/user";
import { useLocalStorage } from "@vueuse/core";
const userLocalInfo = useLocalStorage<string>("userinfo", "");
const acccess_token = useLocalStorage<string>("acccess_token", "");
const refresh_token = useLocalStorage<string>("refresh_token", "");
import qs from "query-string";
export type Obj = Record<string, any>;

export const getUrlQuery = () => {
  const query = qs.parse(window.location.search);
  if (query.uid === undefined) {
    query.uid = "";
  }
  return query;
};
export const setToken = (aToken: string, rToken: string) => {
  acccess_token.value = aToken;
  refresh_token.value = rToken;
};
export const setLocalStorageUserInfo = (uinfo: userInfo) => {
  userLocalInfo.value = JSON.stringify(uinfo);
};
export const clearLocalInfo = () => {
  userLocalInfo.value = "";
  acccess_token.value = "";
};
