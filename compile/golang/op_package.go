package golang

import (
	schema "opticode/compile/golang/golang"
)

func (g *Generator) op_package(node *schema.IndexedNode, flags schema.Flag) ([]byte, error) {
	id, err := g.LookUpStr(node.Id())
	if err != nil {
		return nil, err
	}

	return JoinBytes(TokenPackage.Bytes(), TokenSpace.Bytes(), id), nil
}
