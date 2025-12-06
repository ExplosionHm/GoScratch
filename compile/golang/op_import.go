package golang

import (
	"bytes"
	"fmt"
	schema "opticode/compile/golang/golang"
)

func (g *Generator) op_import(node *schema.IndexedNode, flags schema.Flag) ([]byte, error) {
	length := node.FieldsLength()

	buf := new(bytes.Buffer)

	for i := range length {
		var field *schema.NodeValue
		node.Fields(field, i)

		if field.Flags()&schema.ValueFlagPointer != 0 {
			// is pointer
			node := g.GetNode(field.Value())
			if node == nil {
				return nil, fmt.Errorf("attempt to access undefined node: %d", field.Value())
			}
			out, err := g.Eval(node)
			if err != nil {
				return nil, err
			}
			//! This preforms an allocation each cycle. (not good)
			buf.Write(JoinBytes(TokenTab.Bytes(), out, TokenNewLine.Bytes()))
		} else {
			return nil, fmt.Errorf("import node fields can only be pointers")
		}
	}

	if length > 1 {
		//! messy and needs work
		return JoinBytes(TokenImport.Bytes(), TokenSpace.Bytes(), TokenParenLeft.Bytes(), TokenNewLine.Bytes(), buf.Bytes(), TokenParenRight.Bytes()), nil
	}
	//! includes the tab seperator (not intended)
	return JoinBytes(TokenImport.Bytes(), buf.Bytes()), nil
}
