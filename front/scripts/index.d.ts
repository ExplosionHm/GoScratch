declare type Position = {
  x: number;
  y: number;
};

declare enum BlockTypes {
  Hat = 0,
  Stack,
  Boolean,
  Reporter,
  C,
}

declare type Attribute = {
  id: string;
  modifier: {
    //! TODO
    type: undefined;
    value: undefined;
  };
};

//(presage)
//(uid) (assignType) (lib).(parents).(id)(args)
//(tails)
declare type Block = {
  lib?: "scratch" | string;
  assignType?: "assign" | "append" | "map";
  parents?: [];
  id: string;
  uid: string;
  args: any[];
  type: BlockTypes;
  apperence: {
    pos: Position;
    text: string;
    attributes: Attribute[];
  };
};

declare type Defintion = {
  id: string;
  value?: string;
  type?: string;
};

declare type Import = {
  implements?: false;
  url: string;
};

declare type Payload = {
  "+imports": Import[];
  "+defintions": Defintion[];
  "+blocks": Block[];
};

declare type BlockDefinition = {
  id: string;
  presage?: string[];
  args: JsToGoTypes[];
  tails?: string[];
};

declare enum JsToGoTypes {
  boolean = "bool",
  string = "string",
  int = "int",
  bigint = "int64", //A.K.A Long
  number = "float64",
}

declare type BlockCollection = {
  default: Block;
  definition: BlockDefinition;
};

declare type cmd = [string, string[]];
