import { createApp } from "vue";
import App from "./App.vue";
import router from "@/router/index";

const vueInstance = createApp(App);
vueInstance.use(router)
vueInstance.mount("#app")
