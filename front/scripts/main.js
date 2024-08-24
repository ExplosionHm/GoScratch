import fs from "fs";

const asset = {
  "+types": {
    int: (v) => {
      switch (typeof v) {
        case "undefined":
          return 0;
        case "symbol":
          return 0;
        case "string":
          return parseInt(v);
        case "object":
          return 0;
        case "number":
          return Math.floor(v);
        case "function":
          return 0;
        case "boolean":
          return Number(v);
        case "bigint":
          return Number(v);
        default:
          return 0;
      }
    },
    bool: (v) => {
      return Boolean(v);
    },
    WindowConfig: {
      parent: "pixelgl",
    },
  },
  CreateWindow: {
    args: ["string", "int", "int", "bool"],
    tails: [],
  },
  SetSmooth: {
    args: ["bool"],
    tails: [],
  },
  CreateSprite: {
    args: ["string"],
    tails: [],
  },
  Draw: {
    args: ["target", "matrix"],
    tails: [],
  },
  Update: {
    args: [],
    tails: [],
  },
};

const p = {
  /**@type {load[]} */
  start: [
    //global parents: "pixel", "pixelgl"
    {
      parents: ["scratch"],
      assign_type: "append", //! add appending
      id: "CreateWindow",
      uid: "windows",
      args: ['"GoScratch"', 1024, 768, true],
    },
    {
      parents: ["windows"],
      id: "SetSmooth",
      args: [true],
    },
    {
      parents: ["scratch"],
      id: "CreateSprite",
      uid: "s",
      args: ['"images/gopher.png"'],
    },
  ],
  /**@type {load[]} */
  loop: [
    {
      parents: ["s"],
      id: "Draw",
      args: ["windows", "pixel.IM.Moved(windows.MousePosition())"],
    },
    {
      parents: ["windows"],
      id: "Update",
      args: [],
    },
  ],
};

/**
 * @param {load[]} payload
 */
function parse(payload) {
  let out = "";
  for (const load of payload) {
    const h = asset[load.id];
    if (!h) continue;
    let parents;
    parents = load.parents.join(".") + ".";
    if (load.uid) out += `${load.uid} := `;
    out += `${parents || ""}${load.id}`;
    out += `(${load.args.join(", ")})`;
    for (const tail of h.tails) {
      out += `\n${tail.replace("%s", load.uid)}`;
    }
    out += "\n";
  }
  return out;
}

(async () => {
  const file = fs.readFileSync("front/test.txt", {
    encoding: "utf8",
  });

  if (!file) throw new Error();

  let out = file;
  for (const [id, value] of Object.entries(p)) {
    const target = `/*@INSERT ${id}*/`;
    const pos = out.indexOf(target) + target.length;
    if (target.length > pos) return;
    out =
      out.slice(0, pos) +
      "\n" +
      String(parse(value)) +
      out.slice(pos, out.length);
  }
  fs.writeFileSync("front/test.txt", out);
})();
