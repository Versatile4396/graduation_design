// import qs from "query-string";
import { useLocalStorage } from "@vueuse/core";
const userLocalInfo = useLocalStorage<string>("userinfo", "");
const token = useLocalStorage<string>("token", "");
import qs from "query-string";
export type Obj = Record<string, any>;

export const getUrlQuery = () => {
  const query = qs.parse(window.location.search);
  if (query.uid === undefined) {
    query.uid = "";
  }
  return query;
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
