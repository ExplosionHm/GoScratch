import { ipcRenderer } from "electron";

document.addEventListener("DOMContentLoaded", main);

function main() {
  let lastCompiled;
  document.getElementById("run")?.addEventListener("click", (e) => {
    ipcRenderer.send("run");
  });
  document.getElementById("compile")?.addEventListener("click", (e) => {
    ipcRenderer.send("compile");

    document.getElementById("runCompile").disabled = false;
  });
  document.getElementById("runCompile")?.addEventListener("click", (e) => {
    ipcRenderer.send("runCompiled");
  });
}
