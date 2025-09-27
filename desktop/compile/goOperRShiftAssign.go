package compile

import (
	"opticode/desktop/tree"
)

// TODO: Should indent
func (gg *GoGenerator) op_operRShiftAssign(node tree.Node) (string, *Error) {
	if len(node.Fields) != 2 {
		return "", Err(Fatal, "right shift assign operation cannot take in more than or less than two values: %d values specified", len(node.Fields))
	}

	args, err := gg.evalArgs(node)
	if err != nil {
		return "", err
	}
	// TODO: Add type check

	return args[0] + " >>= " + args[1], nil
}
