package compile

import (
	"fmt"
	"opticode/desktop/tree"
	"strings"
)

func (gg *GoGenerator) op_funcCall(node tree.Node) (string, error) {
	var result string
	def := lookupFunc(gg.Libaries, node.Opid)
	if def == nil {
		return "", fmt.Errorf("undefined function: %s", node.Opid)
	}
	result = gg.Indent + node.Opid + "("

	for i, v := range node.Fields {
		ty := v.GetType()
		argIndex := nearestArgument(def.Arguments, i)
		if ty.Name() == def.Arguments[argIndex][1] || strings.HasSuffix(def.Arguments[argIndex][1], "Type") {
			// TODO: Can be improved
			if v.HasQuotes {
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

// TODO: Only take in the length
func nearestArgument(args [][]string, i int) int {
	length := len(args)
	if i <= 0 || length < i {
		// is within length
		return i
	}

	return length - 1 // Return last
}
