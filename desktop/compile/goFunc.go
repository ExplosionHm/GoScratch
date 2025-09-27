package compile

import (
	"log"
	"opticode/desktop/tree"
)

func (gg *GoGenerator) op_func(node tree.Node) (string, error) {
	var result string
	var args string
	var returns = []string{}
	var body string
	log.Println("Enter func")
	result = "func "
	// TODO: func ref implementation
	result += node.Opid + "("
	for _, arg := range node.Fields {
		log.Println("Enter field range")
		if arg.Flags&tree.IsFuncArg != 0 {
			if len(args) > 0 {
				args += "," + arg.Value + " " + arg.Type
			} else {
				args += arg.Value + " " + arg.Type
			}
		}

		// Should have `IsPointer` check but a pointer is garenteed in this senario
		if arg.Flags&tree.IsFuncBody != 0 {
			if node, ok := gg.Tree.Nodes[arg.Value]; ok {
				n, err := gg.Eval(node, true)
				if err != nil {
					return "", err
				}
				body += n
			}
		}
		// Should have `IsPointer` check but a pointer is garenteed to not be in this senario
		if arg.Flags&tree.IsFuncReturns != 0 {
			returns = append(returns, arg.Value+" "+arg.Type)
		}

	}
	result += args + ")"
	if len(returns) > 0 { //! this is bad code
		if len(returns) > 1 {
			result += "("
		}
		for i, re := range returns {
			if i > 0 {
				result += "," + re
			} else {
				result += re
			}
		}
		if len(returns) > 1 {
			result += ")"
		}
	}
	result += " {\n" + body + "}\n\n"
	log.Println("exit func", *node.Parent)
	if nextExists := gg.finishNode(node); !nextExists {
		log.Println("Failed to find next")
	}
	return result, nil
}
