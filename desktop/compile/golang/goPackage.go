package golang

import (
	"log"
	"opticode/desktop/tree"
	"opticode/utils"
	"regexp"
)

func (gg *GoGenerator) op_package(node tree.Node) (string, error) {
	isInvalid := regexp.MustCompile(`^(?:\d.*|.*[^a-zA-Z0-9].*)$`)
	//! This check should be client side only
	if isInvalid.MatchString(node.Tab) {
		return "", utils.Err(utils.Fatal, "package identifier cannot start with a number or contain non-alphanumeric characters")
	}

	if nextExists := gg.finishNode(node); !nextExists {
		log.Println("Failed to find next")
	}
	return "package " + node.Tab + "\n\n", nil
}
