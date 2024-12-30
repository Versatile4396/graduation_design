const { createRsbuild, mergeRsbuildConfig } = require("@rsbuild/core");
const path = require("path");
const command = process.argv[2];
async function run() {
  rsbuildCLI(command);
}
const rsbuildCLI = async function (command = "build") {
  const rsbuildConfig = require("../rsbuild.config.js");
  const rsbuild = await createRsbuild({ rsbuildConfig });
  if (command === "dev" || command == "serve") {
    rsbuild.startDevServer();
  } else {
    rsbuild.build();
  }
};

run();
