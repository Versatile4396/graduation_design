import { createApp } from "vue";
import App from "./App.vue";
import router from "@/router/index";
import "@/assets/style/global.css";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import { createPinia } from "pinia";

const pinia = createPinia();
const vueInstance = createApp(App);
vueInstance.use(ElementPlus);
vueInstance.use(router);
vueInstance.use(pinia);
vueInstance.mount("#app");
