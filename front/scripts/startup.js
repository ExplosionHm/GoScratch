import { app, BrowserWindow, co } from "electron";

let window;

app.whenReady().then(() => {
  window = new BrowserWindow({
    width: 1024,
    height: 768,
    title: "GoScratch Editor",
    autoHideMenuBar: true,
  });
});
