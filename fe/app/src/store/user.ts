import { getUrlQuery } from "@/utils/common";
import { defineStore } from "pinia";

export const userInfoStore = defineStore("userInfo", () => {
  const query = getUrlQuery();
});
