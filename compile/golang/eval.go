package golang

import (
	"fmt"
	schema "opticode/compile/golang/golang"

	fb "github.com/google/flatbuffers/go"
)

func (g *Generator) Eval(node *schema.Node) ([]byte, error) {
	unionTable := new(fb.Table)

	if node.Node(unionTable) {
		nodeType := node.NodeType()
		out := []byte{}
		switch nodeType {
		case schema.NodeUnionIndexedNode:
			type1 := new(schema.IndexedNode)
			type1.Init(unionTable.Bytes, unionTable.Pos)
			var err error
			out, err = g.EvalType1(node.Opcode(), type1, node.Flags())
			if err != nil {
				return nil, err
			}
		case schema.NodeUnionBinaryNode:
			type2 := new(schema.BinaryNode)
			type2.Init(unionTable.Bytes, unionTable.Pos)
			var err error
			out, err = g.EvalType2(node.Opcode(), type2, node.Flags())
			if err != nil {
				return nil, err
			}
		case schema.NodeUnionUnaryNode:
			type3 := new(schema.UnaryNode)
			type3.Init(unionTable.Bytes, unionTable.Pos)
			var err error
			out, err = g.EvalType3(node.Opcode(), type3, node.Flags())
			if err != nil {
				return nil, err
			}
		case schema.NodeUnionNONE:
			return nil, fmt.Errorf("failed to determine node type of node: %d", node.Id())
		}
		g.Write(node.Id(), node.Flags(), 0, &out)
	}
	return nil, fmt.Errorf("failed to access union of node: %d", node.Id())
}

func (g *Generator) EvalType1(opcode schema.Opcode, node *schema.IndexedNode, flags schema.Flag) ([]byte, error) {
	switch opcode {
	case 0:
		return g.op_package(node, flags)
	case 1:
		return g.op_import(node, flags)
	}
	return nil, nil
}

func (g *Generator) EvalType2(opcode schema.Opcode, node *schema.BinaryNode, flags schema.Flag) ([]byte, error) {
	switch opcode {
	case 2:
		return g.op_importValue(node, flags)
	}
	return nil, nil
}

func (g *Generator) EvalType3(opcode schema.Opcode, node *schema.UnaryNode, flags schema.Flag) ([]byte, error) {
	switch opcode {
	}
	return nil, nil
}
