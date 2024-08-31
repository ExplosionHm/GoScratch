import fs from "fs";

/**@type {BlockDefinition[]} */
const asset = [
  {
    id: "CreateWindow",
    args: ["string", "int", "int", "bool"],
  },
  {
    id: "SetSmooth",
    args: ["bool"],
    tails: [],
  },
  {
    id: "CreateSprite",
    args: ["string"],
    tails: [],
  },
  {
    id: "Draw",
    args: ["target", "matrix"],
    tails: [],
  },
  {
    id: "Update",
    args: [],
    tails: [],
  },
];
class Import {
  constructor() {
    /**@type {Import[]} */
    this.imports = [];
  }
  /**
   * @param {Import} Import
   */
  add(Import) {
    this.imports.push(Import);
    return this;
  }
  /**
   * @param {Import[]} Imports
   */
  set(Imports) {
    this.imports = Imports;
    return this;
  }
}

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
   * @param {BlockDefinition[]} defintions
   * @returns {string}
   */
  getBlocks(defintions) {
    //(presage)
    //(uid) (assignType) (lib).(parents).(id)(args)
    //(tails)
    let out = "";
    for (const block of this.blocks) {
      const def = defintions.find((v) => v.id === block.id);
      if (!def) {
        console.log(`Definition not found for: ${block.id}`);
        continue;
      }
      if (def.presage) out += def.presage.join("\n") + "\n";
      let cap = "";
      switch (block.assignType) {
        case "assign":
          const inDefintions = this.defintions.find((v) => v.id === block.uid)
            ? true
            : false;
          out += `${block.uid} ${inDefintions ? "=" : ":="} `;
          break;
        case "append":
          cap = ")";
          out = `${block.uid} = append(${block.uid}, `;
          break;
        case "map":
          break;
        default:
          break;
      }
      if (block.lib) out += block.lib + ".";
      if (block.parents) out += block.parents.join(".") + ".";
      out += `${block.id}(${block.args.join(", ")})`;
      if (cap) out += cap;
      out += "\n";
      if (def.tails) out += def.tails.join("\n") + "\n";
    }

    return out;
  }

  /**
   * @param {Defintion[]} defintions
   * @return {string}
   */
  compile(defintions) {
    let out = "";
    out += this.getImports();
    out += this.getDefintions();
    out += this.getBlocks(defintions);
    return out;
  }
}

const i = new Import().add({ url: "goscratch/scratch" }).imports;
const d = [
  {
    id: "test",
    type: "int",
  },
];

/**@type {Block[]} */
const b = [
  {
    lib: "scratch",
    assignType: "assign",
    id: "CreateWindow",
    uid: "window",
    args: ['"GoScratch"', 1024, 768, true],
  },
  {
    assignType: "none",
    id: "SetSmooth",
    parents: ["window"],
    uid: "",
    args: [true],
  },
  {
    lib: "scratch",
    assignType: "assign",
    id: "CreateSprite",
    uid: "s",
    args: [
      '"C:/Users/sandn/Documents/Projects/goscratch/back/images/gopher.png"',
    ],
  },

  //! ADD STATEMENTS
];

const f = new GoScratch(i, d, b);

console.log(f.compile(asset));

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

async () => {
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
};
