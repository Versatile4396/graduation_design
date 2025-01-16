import { createApp } from "vue";
import App from "./App.vue";
import router from "@/router/index";
import "@/assets/style/global.css";
import 'ant-design-vue/dist/reset.css';

const vueInstance = createApp(App);
vueInstance.use(router)
vueInstance.mount("#app")
