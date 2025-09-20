package compile

import (
	"opticode/desktop/tree"
	"strings"
)

func (gg *GoGenerator) op_return(node tree.Node) (string, error) {
	var result string
	gg.IncreaseIndent()
	var args = []string{}
	for _, arg := range node.Fields {
		if arg.Flags&tree.IsPointer != 0 {
			n, err := gg.Eval(node)
			if err != nil {
				return "", err
			}

			args = append(args, n)
			continue
		}

		if arg.Flags&tree.HasQuotes != 0 {
			args = append(args, "\""+arg.Value+"\"")
		} else {
			args = append(args, arg.Value)
		}
	}
	result = gg.Indent() + "return " + strings.Join(args, ",") + "\n"

	gg.DecreaseIndent()
	return result, nil
}
