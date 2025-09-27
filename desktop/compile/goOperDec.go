package compile

import (
	"opticode/desktop/tree"
)

func (gg *GoGenerator) op_operDec(node tree.Node) (string, *Error) {
	if len(node.Fields) != 1 {
		return "", Err(Fatal, "decrement operation cannot take in more than or less than one values: %d values specified", len(node.Fields))
	}

	var arg1 string
	var hasParentheses bool
	if node.Fields[0].Flags&tree.IsPointer != 0 {
		hasParentheses = true
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
	if hasParentheses {
		return "(" + arg1 + ")--", nil
	}
	return arg1 + "--", nil
}
