package compile

import (
	"fmt"
	"opticode/desktop/tree"
)

// TODO: Should indent
func (gg *GoGenerator) op_operDivAssign(node tree.Node) (string, error) {
	if len(node.Fields) != 2 {
		return "", fmt.Errorf("division assign operation cannot take in more than or less than two values: %d values specified", len(node.Fields))
	}

	var arg1 string
	if node.Fields[0].Flags&tree.IsPointer != 0 {
		if node, ok := gg.Tree.Nodes[node.Fields[0].Value]; ok {
			n, err := gg.Eval(node)
			if err != nil {
				return "", err
			}

			arg1 = n
		}
	} else {
		if node.Fields[0].Flags&tree.HasQuotes != 0 {
			arg1 = "\"" + node.Fields[0].Value + "\""
		} else {
			arg1 = node.Fields[0].Value
		}
	}

	var arg2 string
	if node.Fields[1].Flags&tree.IsPointer != 0 {
		if node, ok := gg.Tree.Nodes[node.Fields[1].Value]; ok {
			n, err := gg.Eval(node)
			if err != nil {
				return "", err
			}

			arg2 = n
		}
	} else {
		if node.Fields[1].Flags&tree.HasQuotes != 0 {
			arg2 = "\"" + node.Fields[1].Value + "\""
		} else {
			arg2 = node.Fields[1].Value
		}
	}
	// TODO: Add type check

	return arg1 + " /= " + arg2, nil
}
