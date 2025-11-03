package golang

import (
	"opticode/desktop/tree"
)

func (g *Generator) op_import(node *tree.Type1, flags tree.Flag) ([]byte, error) {
	id, err := g.LookUp(node.Id())
	if err != nil {
		return nil, err
	}
	return JoinBytes(TokenPackage.Bytes(), TokenSpace.Bytes(), id), nil
}
