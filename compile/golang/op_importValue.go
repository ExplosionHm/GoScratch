package golang

import (
	"bytes"
	"opticode/tree"
)

func (g *Generator) op_importValue(node *tree.Type2, flags tree.Flag) ([]byte, error) {
	var left *tree.NodeValue
	node.Left(left)

	var right *tree.NodeValue
	node.Left(right)

	var buf = new(bytes.Buffer)

	// write import alias
	leftValue, err := g.LookUp(uint32(left.Value()))
	if err != nil {
		return nil, err
	}
	buf.Write(leftValue)
	// write seperator
	buf.WriteByte(' ')
	// write package path
	rightValue, err := g.LookUp(uint32(left.Value()))
	if err != nil {
		return nil, err
	}
	buf.Write(JoinBytes(TokenQuotation.Bytes(), rightValue, TokenQuotation.Bytes()))

	return buf.Bytes(), nil
}
