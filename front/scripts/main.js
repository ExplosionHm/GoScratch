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
  },
  CreateWindow: {
    args: ["*Screen", "string", "int", "int"],
    tails: ["defer %s.Release()"],
  },
  NewBuffer: {
    tails: [
      "if err != nil {",
      'log.Fatalf("%v - failed to create screen buffer", err)',
      "}",
      "defer %s.Release()",
    ],
  },
};

const p = [
  //global parents: "pixel", "pixelgl"
  {
    id: "CreateWindow",
    uid: "w",
    args: ["screen", '"GoScratch"', 400, 600],
  },
  {
    parent: "screen",
    id: "NewBuffer",
    uid: "screenBuffer",
    args: ["rect"],
  },
];

/**
 * @param {load[]} payload
 */
function parse(payload) {
  let out = "";
  for (const load of payload) {
    const h = asset[load.id];
    if (!h) continue;
    out += `${load.uid} := ${load.parent || ""}${load.id}`;
    out += `(${load.args.join(", ")})`;
    for (const tail of h.tails) {
      out += `\n${tail.replace("%s", load.uid)}`;
    }
    out += "\n";
  }
  return out;
}

(() => {
  const file = fs.readFileSync("front/test.txt", {
    encoding: "utf8",
  });

  if (!file) throw new Error();

  let out = "";
  const target = "/*@INSERT start*/";
  const pos = file.indexOf(target) + target.length;
  if (target.length > pos) return;
  out =
    file.slice(0, pos) + "\n" + String(parse(p)) + file.slice(pos, file.length);
  fs.writeFileSync("front/test.txt", out);
})();
