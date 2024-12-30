const { createRsbuild, mergeRsbuildConfig } = require('@rsbuild/core');
const path = require('path');
const { merge: mergeWebpackConfig } = require('webpack-merge');
const fs = require('fs');

const buildTool = process.argv[2] || 'rsbuild';
const command = process.argv[3] || 'build';
// 通用配置
const commonConfig = require(`./build-configs/common.config.${buildTool}.js`);
// 根据执行文件夹路径获取相应配置
const appName = process.argv[4] || path.basename(process.cwd()) || 'home';
const appConfigPath = path.resolve(__dirname, `./build-configs/${appName}/${buildTool}.config.js`);
const appConfig = require(appConfigPath);

async function run(params) {
  if (buildTool === 'rsbuild') {
    rsbuildCLI(command);
  } else {
    webpackCLI(command);
  }
}
const filePath = path.resolve(__dirname, `./config/${buildTool}.${appName}.config.js`);

const rsbuildCLI = async function (command = 'build') {
  const rsbuildConfig = mergeRsbuildConfig(commonConfig, appConfig);
  const rsbuild = await createRsbuild({ rsbuildConfig });
  if (command === 'dev' || command == 'serve') {
    rsbuild.startDevServer();
  } else {
    rsbuild.build();
  }
};

const webpackCLI = function (command) {
  const webpackConfig = mergeWebpackConfig({}, appConfig);
  const compiler = webpack(webpackConfig);
  if (command === 'dev' || command == 'serve') {
    const serverOptions = {
      ...webpackConfig.devServer
    };
    // 启动webpack Dev Server
    const server = new webpackDevServer(serverOptions, compiler);
    server
      .start()
      .then(() => {})
      .catch((err) => {});
  } else {
    compiler.run((err, stats) => {});
  }
};

run();
