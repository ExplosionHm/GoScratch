package golang

import (
	"log"
	"opticode/desktop/tree"
	"opticode/utils"
	"strings"
)

func (gg *GoGenerator) op_funcCall(node tree.Node, isParent ...bool) (string, error) {
	var IsParent bool
	var result string
	if len(isParent) > 0 && isParent[0] {
		IsParent = true
		gg.IncreaseIndent()
		defer gg.DecreaseIndent()
		result = gg.Indent()
	}

	def := lookupFunc(gg.Libaries, node.Opid)
	if def == nil {
		return "", utils.Err(utils.Fatal, "undefined function: %s", node.Opid)
	}
	result += node.Opid + "("

	for i, v := range node.Fields {
		argIndex := utils.Nearest(len(def.Arguments), i)
		if v.Type == def.Arguments[argIndex][1] || strings.HasSuffix(def.Arguments[argIndex][1], "Type") { // TODO: Remove hardcoded value "Type"

			if v.Flags&tree.IsPointer != 0 {
				if node, ok := gg.Tree.Nodes[v.Value]; ok {
					n, err := gg.Eval(node)
					if err != nil {
						return "", err
					}
					result += n
					continue
				}
			}

			// TODO: Can be improved
			if v.Flags&tree.HasQuotes != 0 {
				if i > 0 {
					result += ", " + "\"" + v.Value + "\""
				} else {
					result += "\"" + v.Value + "\""
				}
			} else {
				if i > 0 {
					result += "," + v.Value
				} else {
					result += v.Value
				}
			}
		} else {
			return "", utils.Err(utils.Fatal, "syntax error: type mismatch for %s", node.Opid)
		}
	}

	result += ")"
	if IsParent {
		result += "\n"
		if nextExists := gg.finishNode(node); !nextExists {
			log.Println("Failed to find next-")

		}
	}
	return result, nil
}

func lookupFunc(libs []*GoLibrary, opid string) *FuncDef {
	for _, lib := range libs {
		def := lib.LookupFunc(opid)
		if def != nil {
			return def
		}
	}
	return nil
}
