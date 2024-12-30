const { defineConfig } = require("@rsbuild/core");
const { pluginVue } = require("@rsbuild/plugin-vue");
const { pluginLess } = require("@rsbuild/plugin-less");
const { pluginSass } = require("@rsbuild/plugin-sass");
const path = require("node:path");

const cwdPath = process.cwd();

module.exports = defineConfig({
  dev: {
    client: {
      protocol: "ws",
      host: "127.0.0.1",
    },
  },
  server: {
    port: 5050,
  },
  html: {
    title: "跨专业学习论坛",
    inject: "body",
    template: path.resolve(cwdPath, "./index.rs.html"),
  },
  plugins: [pluginVue(), pluginLess(), pluginSass()],
  source: {
    entry: {
      index: path.resolve(cwdPath, "./src/main.ts"),
    },
    alias: {
      "@": path.resolve(cwdPath, "./src"),
    },
  },
  output: {
    distPath: {
      root: path.resolve(cwdPath, "./dist"),
    },
    sourceMap: {
      js:
        process.env.NODE_ENV === "production"
          ? // 生产模式使用高质量的 source map 格式
            "source-map"
          : // 开发模式使用性能更好的 source map 格式
            "cheap-module-source-map",
    },
    externals: [/node_modules/],
  },
});
