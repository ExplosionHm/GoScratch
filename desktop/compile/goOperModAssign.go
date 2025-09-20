package compile

import (
	"fmt"
	"opticode/desktop/tree"
)

// TODO: Should indent
func (gg *GoGenerator) op_operModAssign(node tree.Node) (string, error) {
	if len(node.Fields) != 2 {
		return "", fmt.Errorf("mod assign operation cannot take in more than or less than two values: %d values specified", len(node.Fields))
	}

	args, err := gg.evalArgs(node)
	if err != nil {
		return "", err
	}
	// TODO: Add type check

	return args[0] + " %= " + args[1], nil
}
