package compile

import (
	"fmt"
	"opticode/desktop/tree"
)

func (gg *GoGenerator) op_operAddAssign(node tree.Node) (string, error) {
	if len(node.Fields) != 2 {
		return "", fmt.Errorf("addAssign operation cannot take in more than or less than two values: %d values specified", len(node.Fields))
	}

	args, err := gg.evalArgs(node)
	if err != nil {
		return "", err
	}
	// TODO: Add type check

	return args[0] + " += " + args[1], nil
}
