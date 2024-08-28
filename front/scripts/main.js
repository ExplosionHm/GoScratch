import fs from "fs";
import { exec } from "child_process";

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

class GoScratch {
  /**
   * @param {Import[]} imports
   * @param {Defintion[]} defintions
   * @param {Block[]} blocks
   */
  constructor(imports, defintions, blocks) {
    this.imports = imports;
    this.defintions = defintions;
    this.blocks = blocks;
  }
  /**
   * @returns {string}
   */
  getImports() {
    let out = "";
    for (const Import of this.imports) {
      out += `${Import.implements ? "_" : ""}"${Import.url}"\n`;
    }
    return out;
  }
  /**
   * @returns {string}
   */
  getDefintions() {
    let out = "";
    let value = "";
    for (const defintion of this.defintions) {
      if (!defintion.type) {
        if (!defintion.value) value = "any";
        else value = `= `;
      }
      //! TODO
      //if only type then just put the type
      //if only value then just put " = value"
      //if both type and value then try converting value to said type
      out += `${defintion.id} ${defintion.type}\n`;
    }
    return out;
  }
  /**
   * @returns {string}
   */
  getBlocks() {}
}

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
