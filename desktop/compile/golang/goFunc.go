package golang

import (
	"log"
	"opticode/desktop/tree"
	"strings"
)

func (gg *GoGenerator) op_func(node tree.Node) (string, error) {
	var result strings.Builder
	var args = []string{}
	var returns = []string{}
	var body strings.Builder
	// TODO: func ref implementation
	result.WriteString("func " + node.Opid + "(")
	for _, arg := range node.Fields {
		if arg.Flags&tree.IsFuncArg != 0 {
			args = append(args, arg.Value+" "+arg.Type)
		}

		// Should have `IsPointer` check but a pointer is garenteed in this senario
		if arg.Flags&tree.IsFuncBody != 0 {
			if node, ok := gg.Tree.Nodes[arg.Value]; ok {
				n, err := gg.Eval(node, true)
				if err != nil {
					return "", err
				}
				body.WriteString(n)
			}
		}
		// Should have `IsPointer` check but a pointer is garenteed to not be in this senario
		if arg.Flags&tree.IsFuncReturns != 0 {
			returns = append(returns, arg.Value+" "+arg.Type)
		}

	}

	result.WriteString(strings.Join(args, ", ") + ")")
	var reLen = len(returns)
	if reLen > 0 {
		re := strings.Join(returns, ", ")
		if reLen > 1 {
			result.WriteString("(" + re + ")")
		} else {
			result.WriteString(re)
		}
	}
	result.WriteString(" {\n" + body.String() + "}\n\n")
	if nextExists := gg.finishNode(node); !nextExists {
		log.Println("Failed to find next")
	}
	return result.String(), nil
}
