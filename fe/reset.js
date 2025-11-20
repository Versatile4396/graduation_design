const fs = require("fs");
const path = require("path");

function deleteNodeModules(dir) {
  const files = fs.readdirSync(dir);
  for (const file of files) {
    const curPath = path.join(dir, file);
    const stat = fs.statSync(curPath);
    if (stat.isDirectory()) {
      if (file === "node_modules") {
        // 如果是node_modules文件夹，则删除它
        fs.rmSync(curPath, { recursive: true, force: true });
      } else {
        // 如果是其他文件夹，继续递归查找
        deleteNodeModules(curPath);
      }
    }
  }
}

// 从当前目录开始删除 realy
deleteNodeModules(process.cwd());


// wuyu