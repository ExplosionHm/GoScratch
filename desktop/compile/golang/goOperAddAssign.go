package golang

import (
	"opticode/desktop/tree"
	"opticode/utils"
)

func (gg *GoGenerator) op_operAddAssign(node tree.Node) (string, error) {
	if len(node.Fields) != 2 {
		return "", utils.Err(utils.Fatal, "addAssign operation cannot take in more than or less than two values: %d values specified", len(node.Fields))
	}

	args, err := gg.evalArgs(node)
	if err != nil {
		return "", err
	}
	// TODO: Add type check

	return args[0] + " += " + args[1], nil
}
