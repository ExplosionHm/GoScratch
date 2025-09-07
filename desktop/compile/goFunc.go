package compile

import (
	"log"
	"opticode/desktop/tree"
)

func (gg *GoGenerator) op_func(node tree.Node) (string, error) {
	var result string
	var args string
	var funcBody string
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
				n, err := gg.Eval(node)
				if err != nil {
					return "", err
				}
				funcBody += n
			}
		}
	}
	result += args + ") {\n"
	result += funcBody + "}\n"
	log.Println("exit func")
	if nextExists := gg.finishNode(); !nextExists {
		log.Println("Failed to find next")
	}
	return result, nil
}
