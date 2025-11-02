package golang

import (
	"opticode/desktop/tree"
)

func (g *Generator) op_package(node *tree.Type1, flags tree.Flag) ([]byte, error) {
	id, err := g.LookUp(node.Id())
	if err != nil {
		return nil, err
	}
	//! Should try preallocating to avoid reallocations (bad)
	return JoinBytes(TokenPackage.Bytes(), TokenSpace.Bytes(), id), nil
}
