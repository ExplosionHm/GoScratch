package compile

import (
	"log"
	"opticode/desktop/tree"
)

func (gg *GoGenerator) op_if(node tree.Node) (string, error) {
	gg.IncreaseIndent()
	var result string
	var ifCondition string
	var ifBody string
	var elseBody string

	result += gg.Indent() + "if "
	for _, arg := range node.Fields {
		if arg.Flags&tree.IsCondition != 0 {
			if arg.Flags&tree.IsPointer != 0 {
				if node, ok := gg.Tree.Nodes[arg.Value]; ok {
					n, err := gg.Eval(node)
					if err != nil {
						return "", err
					}

					ifCondition += n
				}
			} else {
				// TODO: Find out if this works (probably not well)
				ifCondition += arg.Value
			}
			continue
		}

		// Should have `IsPointer` check but a pointer is garenteed in this senario
		if node, ok := gg.Tree.Nodes[arg.Value]; ok {
			n, err := gg.Eval(node)
			if err != nil {
				return "", err
			}
			if arg.Flags&tree.IsIfBody != 0 {
				ifBody += n
			} else if arg.Flags&tree.IsElseBody != 0 {
				elseBody += n
			}
		}
	}
	result += ifCondition + " {\n" + ifBody + gg.Indent() + "}"

	// TODO: Add support for if...else chaining

	if len(elseBody) > 0 {
		result += " else {\n" + elseBody + gg.Indent() + "}\n"
	} else {
		result += "\n"
	}

	if nextExists := gg.finishNode(node); !nextExists {
		log.Println("Failed to find next")
	}
	gg.DecreaseIndent()
	return result, nil
}
