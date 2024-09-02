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
  const image = document.getElementById("ob");

  document.addEventListener("mousedown", startDrag);

  function startDrag(e) {
    const rect = image.getBoundingClientRect();
    const x = document.clientX - rect.left;
    const y = e.clientY - rect.top;

    let isDragging = false;

    document.addEventListener("mousemove", drag);
    document.addEventListener("mouseup", stopDrag);

    function drag(e) {
      if (!isDragging) return;

      const dx = e.clientX - x;
      const dy = e.clientY - y;

      image.style.position = "absolute";
      image.style.left = `${rect.left + dx}px`;
      image.style.top = `${rect.top + dy}px`;

      isDragging = true;
    }

    function stopDrag() {
      document.removeEventListener("mousemove", drag);
      document.removeEventListener("mouseup", stopDrag);
      isDragging = false;
    }
  }
}
