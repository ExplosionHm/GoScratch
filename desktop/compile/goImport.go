package compile

import (
	"log"
	"opticode/desktop/tree"
	"strings"
)

func (gg *GoGenerator) op_import(node tree.Node) (string, error) {
	var args = []string{}
	for _, arg := range node.Fields {
		args = append(args, arg.Type+" \""+arg.Value+"\"")
	}

	if nextExists := gg.finishNode(node); !nextExists {
		log.Println("Failed to find next")
	}
	return "import (\n" + strings.Join(args, "\n") + "\n)\n\n", nil
}
