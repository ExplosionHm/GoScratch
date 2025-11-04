package golang

import (
	"bytes"
	"fmt"
	"opticode/desktop/tree"
)

func (g *Generator) op_import(node *tree.Type1, flags tree.Flag) ([]byte, error) {
	length := node.FieldsLength()

	buf := new(bytes.Buffer)

	for i := range length {
		var field *tree.NodeValue
		node.Fields(field, i)

		if field.Flags()&tree.ValueFlagPointer != 0 {
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
			buf.Write(JoinBytes(out, TokenNewLine.Bytes()))
		} else {
			return nil, fmt.Errorf("import node fields can only be pointers")
		}
	}

	return nil, nil
}
