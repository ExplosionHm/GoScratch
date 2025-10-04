package golang

import (
	"log"
	"opticode/desktop/tree"
	"strings"
)

func (gg *GoGenerator) op_import(node tree.Node) (string, error) {
	gg.IncreaseIndent()
	defer gg.DecreaseIndent()
	var args = []string{}
	for _, arg := range node.Fields {
		args = append(args, arg.Type+" \""+arg.Value+"\"")
	}

	if nextExists := gg.finishNode(node); !nextExists {
		log.Println("Failed to find next")
	}

	result := "import (\n" + gg.Indent() + strings.Join(args, "\n"+gg.Indent()) + "\n)\n\n"
	return result, nil
}
