package golang

import (
	"fmt"
	"opticode/desktop/tree"

	flatbuffers "github.com/google/flatbuffers/go"
)

type GoFile struct {
}

func (gf *GoFile) Write(dir string) error {
	return nil
}

type Generator struct {
	buf *[]byte
	lut *map[uint32][]byte
}

func NewGenerator(lut *map[uint32][]byte, buf *[]byte) *Generator {
	return &Generator{
		buf,
		lut,
	}
}

func (g *Generator) LookUp(i uint32) ([]byte, error) {
	if v, ok := (*g.lut)[i]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("Look-up failed: cannot find item with index %d", i)
}

// Compiles buffer into go files.
// lut -> Look-up table for strings used in the program
// buf -> flatbuffer's buffer
func Compile(lut map[uint32][]byte, buf []byte) ([]*GoFile, error) {
	program := tree.GetRootAsProgram(buf, 0)

	nodesLength := program.NodesLength()
	gen := NewGenerator(&lut, &buf)

	// Looping world start here
	var node *tree.Node
	program.Nodes(node, nodesLength-1)

	unionTable := new(flatbuffers.Table)

	if node.Node(unionTable) {
		nodeType := node.NodeType()
		switch nodeType {
		case tree.NodeUnionType1:
			type1 := new(tree.Type1)
			type1.Init(unionTable.Bytes, unionTable.Pos)

			gen.EvalType1(node.Opcode(), type1, node.Flags())
		case tree.NodeUnionType2:

		}
	}
	return nil, nil
}

func (g *Generator) EvalType1(opcode tree.Opcode, node *tree.Type1, flags uint32) ([]byte, error) {
	switch opcode {
	case 0:
		return g.op_package(node, flags)
	}
	return nil, nil
}
