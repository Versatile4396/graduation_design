import { setToken, setLocalStorageUserInfo } from "@/utils/common";
import { defineStore } from "pinia";
import { ref } from "vue";
import { type userInfo } from "@/ajax/type/user";
export const userInfoStore = defineStore("userInfo", () => {
  const userInfo = ref<Partial<userInfo>>({});
  const setUserInfo = (_userInfo: userInfo) => {
    userInfo.value = _userInfo;
    setLocalStorageUserInfo(_userInfo);
    setToken(_userInfo.access_token, _userInfo.refresh_token);
  };
  return { userInfo, setUserInfo };
});
