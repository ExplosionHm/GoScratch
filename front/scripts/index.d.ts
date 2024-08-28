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

declare type cmd = [string, string[]];
