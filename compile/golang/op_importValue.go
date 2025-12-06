package golang

import (
	"bytes"
	schema "opticode/compile/golang/golang"
)

func (g *Generator) op_importValue(node *schema.BinaryNode, flags schema.Flag) ([]byte, error) {
	var left *schema.NodeValue
	node.Left(left)

	var right *schema.NodeValue
	node.Left(right)

	var buf = new(bytes.Buffer)

	// write import alias
	leftValue, err := g.LookUpStr(uint32(left.Value()))
	if err != nil {
		return nil, err
	}
	buf.Write(leftValue)
	// write seperator
	buf.WriteByte(' ')
	// write package path
	rightValue, err := g.LookUpStr(uint32(left.Value()))
	if err != nil {
		return nil, err
	}
	buf.Write(JoinBytes(TokenQuotation.Bytes(), rightValue, TokenQuotation.Bytes()))

	return buf.Bytes(), nil
}
