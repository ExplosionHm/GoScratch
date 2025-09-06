package compile

import (
	"log"
	"opticode/desktop/tree"
)

func (gg *GoGenerator) op_if(node tree.Node) (string, error) {
	log.Println("if")
	gg.IncreaseIndent()
	var result string

	result += gg.Indent() + "if ("
	for _, arg := range node.Fields {
		// Break if not part of condition
		if arg.Flags&tree.IsCondition == 0 {
			break
		}

		if arg.Flags&tree.IsPointer != 0 {

		}
	}
	result += ") {\n"

	//* TEMP
	tmp, err := gg.op_funcCall(gg.Tree.Nodes[gg.NodeTracker[len(gg.NodeTracker)-1]])
	if err != nil {
		return "", err
	}
	result += tmp + gg.Indent() + "}\n"
	err = gg.finishNode()
	if err != nil {
		return "", err
	}
	gg.DecreaseIndent()
	return result, nil
}
