package compile

import (
	"fmt"
	"log"
	"opticode/desktop/tree"
	"strings"
)

func (gg *GoGenerator) op_funcCall(node tree.Node) (string, error) {
	gg.IncreaseIndent()
	var result string
	def := lookupFunc(gg.Libaries, node.Opid)
	if def == nil {
		return "", fmt.Errorf("undefined function: %s", node.Opid)
	}
	result = gg.Indent() + node.Opid + "("

	for i, v := range node.Fields {
		argIndex := nearestArgument(len(def.Arguments), i)
		if v.Type == def.Arguments[argIndex][1] || strings.HasSuffix(def.Arguments[argIndex][1], "Type") { // TODO: Remove hardcoded value "Type"
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
			return "", fmt.Errorf("syntax error: type mismatch for %s", node.Opid)
		}
	}

	result += ")\n"
	if nextExists := gg.finishNode(node); !nextExists {
		log.Println("Failed to find next-")
	}
	gg.DecreaseIndent()
	return result, nil
}

func lookupFunc(libs []*GoLibrary, opid string) *FuncDef {
	for _, lib := range libs {
		def := lib.Library.LookupFunc(opid)
		if def != nil {
			return def
		}
	}
	return nil
}
