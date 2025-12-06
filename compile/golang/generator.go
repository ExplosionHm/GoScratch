package golang

import (
	"fmt"
	schema "opticode/compile/golang/golang"
)

type GoFile struct {
}

func (gf *GoFile) Write(dir string) error {
	return nil
}

type DeserializeNode struct {
	Flags   schema.Flag // Node flags
	Module  uint8       // Index of the nodes' module
	Content *[]byte     // Evaluated node content
	Span    [2]uint32   // File line span
}

type Generator struct {
	program *schema.Program
	buf     *[]byte

	nodeOffsets map[int64]int
	// nodeId -> deserialized node
	nodes map[int64]*DeserializeNode
}

func NewGenerator(program *schema.Program, buf *[]byte) *Generator {
	var nodeOffset = make(map[int64]int)

	for i := range program.NodesLength() {
		var n *schema.Node
		program.Nodes(n, i)

		nodeOffset[n.Id()] = i
	}

	return &Generator{
		program:     program,
		buf:         buf,
		nodeOffsets: nodeOffset,
		nodes:       make(map[int64]*DeserializeNode), //! Should estimate total size
	}
}

func (g *Generator) LookUpStr(i uint32) ([]byte, error) {
	var str *schema.StringEntry
	ok := g.program.LutByKey(str, i)
	if !ok {
		return nil, fmt.Errorf("look-up failed: cannot find item with index %d", i)
	}
	return str.Value(), nil
}

func (g *Generator) LookUpType(t uint32) (*schema.TypeDef, error) {
	var _type *schema.TypeEntry
	ok := g.program.TypesByKey(_type, t)
	if !ok {
		return nil, fmt.Errorf("look-up failed: cannot find item with index %d", t)
	}
	return _type.Value(nil), nil
}

func (g *Generator) GetNode(id int64) *schema.Node {
	i, ok := g.nodeOffsets[id]
	if !ok {
		return nil
	}
	var node *schema.Node
	g.program.Nodes(node, i)

	return node
}

func (g *Generator) Write(id int64, flags schema.Flag, module uint8, content *[]byte) {
	g.nodes[id] = &DeserializeNode{
		Flags:   flags,
		Module:  module,
		Content: content,

		Span: [2]uint32{},
	}
}

func (g *Generator) Export() ([]*GoFile, error) {
	return nil, nil
}
