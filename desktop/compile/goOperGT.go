package compile

import (
	"opticode/desktop/tree"
)

func (gg *GoGenerator) op_operGT(node tree.Node) (string, *Error) {
	if len(node.Fields) != 2 {
		return "", Err(Fatal, "GT operation cannot take in more than or less than two values: %d values specified", len(node.Fields))
	}

	args, err := gg.evalArgs(node)
	if err != nil {
		return "", err
	}

	return args[0] + " > " + args[1], nil
}
