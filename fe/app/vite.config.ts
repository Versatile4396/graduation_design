import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { fileURLToPath } from "node:url";

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    cors: true,
    open: true,
    port: 5199,
    proxy: {
      "/api": {
        target: "http://karrysmile.top:8100/api", // 设置代理目标
        changeOrigin: true, // 是否改变请求源地址
        rewrite: (path: any) => path.replace(/^\/api/, ""), // 将 /api 替换为空字符串
      },
    },
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
