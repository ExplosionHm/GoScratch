package golang

import (
	"fmt"
	"opticode/desktop/tree"
	"sync"

	"log"

	fb "github.com/google/flatbuffers/go"
)

type GoFile struct {
}

func (gf *GoFile) Write(dir string) error {
	return nil
}

type DeserializeNode struct {
	Flags   tree.Flag
	Path    string
	Content *[]byte
	Span    [2]uint32
}

type Generator struct {
	program *tree.Program
	buf     *[]byte
	lut     *map[uint32][]byte

	nodeOffsets map[int64]int
	// nodeId -> deserialized node
	nodes map[int64]*DeserializeNode
}

func NewGenerator(program *tree.Program, lut *map[uint32][]byte, buf *[]byte) *Generator {
	return &Generator{
		program:     program,
		buf:         buf,
		lut:         lut,
		nodeOffsets: make(map[int64]int),
		nodes:       make(map[int64]*DeserializeNode), //! Should estimate total size
	}
}

func (g *Generator) LookUp(i uint32) ([]byte, error) {
	if v, ok := (*g.lut)[i]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("look-up failed: cannot find item with index %d", i)
}

// Can return nil
func (g *Generator) GetNode(id int64) *tree.Node {
	i, ok := g.nodeOffsets[id]
	if !ok {
		return nil
	}
	var node *tree.Node
	g.program.Nodes(node, i)

}

func (g *Generator) Write(id int64, flags tree.Flag, path string, content *[]byte) {
	g.nodes[id] = &DeserializeNode{
		Flags:   flags,
		Path:    path,
		Content: content,

		Span: [2]uint32{},
	}
}

func (g *Generator) Export() ([]*GoFile, error) {
	return nil, nil
}

// Compiles buffer into go files.
// lut -> Look-up table for strings used in the program
// buf -> flatbuffer's buffer
func Compile(lut map[uint32][]byte, buf []byte) ([]*GoFile, error) {
	program := tree.GetRootAsProgram(buf, 0)

	nodesLength := program.NodesLength()
	gen := NewGenerator(&lut, &buf)

	const maxRoutines = 5 // maximum allowed concurrent goroutines

	sem := make(chan struct{}, maxRoutines) // semaphore channel
	var wg sync.WaitGroup

	for i := range nodesLength - 1 {
		sem <- struct{}{} // acquire a "slot" (pauses if full)
		wg.Add(1)

		go func(index int) {
			defer wg.Done()
			defer func() { <-sem }() // release the slot when done

			var node *tree.Node
			program.Nodes(node, index)

			gen.Eval(node)
		}(i)
	}

	wg.Wait()
	fmt.Println("All tasks completed")

	return nil, nil
}

func (g *Generator) Eval(node *tree.Node) {
	unionTable := new(fb.Table)

	if node.Node(unionTable) {
		nodeType := node.NodeType()
		out := []byte{}
		switch nodeType {
		case tree.NodeUnionType1:
			type1 := new(tree.Type1)
			type1.Init(unionTable.Bytes, unionTable.Pos)
			var err error
			out, err = g.EvalType1(node.Opcode(), type1, node.Flags())
			if err != nil {
				log.Println(err) //! Add proper error handling
				return
			}
		case tree.NodeUnionType2:
			type2 := new(tree.Type2)
			type2.Init(unionTable.Bytes, unionTable.Pos)
			var err error
			out, err = g.EvalType2(node.Opcode(), type2, node.Flags())
			if err != nil {
				log.Println(err) //! Add proper error handling
				return
			}
		case tree.NodeUnionType3:
			type3 := new(tree.Type3)
			type3.Init(unionTable.Bytes, unionTable.Pos)
			var err error
			out, err = g.EvalType3(node.Opcode(), type3, node.Flags())
			if err != nil {
				log.Println(err) //! Add proper error handling
				return
			}
		}
		g.Write(node.Id(), node.Flags(), "main.go", &out)
	}
}

func (g *Generator) EvalType1(opcode tree.Opcode, node *tree.Type1, flags tree.Flag) ([]byte, error) {
	switch opcode {
	case 0:
		return g.op_package(node, flags)
	case 1:
		return g.op_import(node, flags)
	}
	return nil, nil
}

func (g *Generator) EvalType2(opcode tree.Opcode, node *tree.Type2, flags tree.Flag) ([]byte, error) {
	switch opcode {
	}
	return nil, nil
}

func (g *Generator) EvalType3(opcode tree.Opcode, node *tree.Type3, flags tree.Flag) ([]byte, error) {
	switch opcode {
	}
	return nil, nil
}
