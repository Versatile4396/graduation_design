import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { fileURLToPath } from "node:url";

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    port: 5050,
  },
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },

  css: {
    // css预处理器
    preprocessorOptions: {
      scss: {
        additionalData: `
          @use "@/assets/style/global.scss";
        `,
      },
    },
  },
});
