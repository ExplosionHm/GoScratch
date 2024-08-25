import { ipcRenderer } from "electron";

document.addEventListener("DOMContentLoaded", main);

function main() {
  document.getElementById("run")?.addEventListener("click", (e) => {
    ipcRenderer.send("run");
  });
  document.getElementById("compile")?.addEventListener("click", (e) => {
    ipcRenderer.send("compile");
  });
}
