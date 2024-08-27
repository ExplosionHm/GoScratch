import { app, BrowserWindow, dialog, ipcMain } from "electron";
import child_process from "child_process";
import { fileURLToPath } from "url";
import { dirname, join } from "path";

const __dirname = dirname(fileURLToPath(import.meta.url));

let window;

app.whenReady().then(() => {
  window = new BrowserWindow({
    width: 1024,
    height: 768,
    title: "GoScratch Editor",
    darkTheme: true,
    webPreferences: {
      nodeIntegration: true,
      preload: join(__dirname, "preload.mjs"),
    },
    autoHideMenuBar: true,
  });
  window.loadFile("index.html");
});

/**
 *
 * @param {cmd[]} cmds
 */
function runCmds(...cmds) {
  /**@type {child_process.ChildProcessWithoutNullStreams} */
  let child;
  for (const [command, args] of cmds) {
    console.log(command, args);
    child = child_process.spawn(command, args, {
      encoding: "utf8",
      shell: true,
    });
    child.on("error", (error) => {
      console.error("An error occured", error);
    });

    child.stdout.setEncoding("utf8");
    child.stdout.on("data", (data) => {
      data = data.toString();
      console.log(data);
    });

    child.stderr.setEncoding("utf8");
    child.stderr.on("data", (data) => {
      console.log(data);
    });
  }
  child.on("close", (code) => {
    //Here you can get the exit code of the script
    switch (code) {
      case 0:
        console.log("Closed child process.");
        break;
    }
  });
}

ipcMain.on("run", (_) => {
  runCmds([
    "C:/Users/sandn/Documents/Projects/goscratch/back/run.bat",
    ["C:/Users/sandn/Documents/Projects/goscratch/back/main.go"],
  ]);
});

ipcMain.on("compile", (_) => {
  runCmds([
    "C:/Users/sandn/Documents/Projects/goscratch/back/compile.bat",
    ["C:/Users/sandn/Documents/Projects/goscratch/back/main.go"],
  ]);
});

ipcMain.on("runCompiled", (_) => {
  runCmds(["C:/Users/sandn/Documents/Projects/goscratch/back/main.exe", []]);
});
