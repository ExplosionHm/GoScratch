package tree

type NodeOp int32

const (
	OP_assign NodeOp = iota
	OP_if
	OP_func
)

type Node struct {
	Opcode   NodeOp `json:"opcode"`
	Fields   []any  `json:"values"`
	Children []Node `json:"children,omitempty"`

	// Only used if the node is a parent
	Position [2]int `json:"-,omitempty"`
	Flags    Flags
}

func (n Node) IsParent() bool {
	if n.Children == nil {
		return false
	}
	return len(n.Children) > 0
}

func NewNode(op NodeOp, flags Flags, fields ...any) Node {
	return Node{
		Opcode:   op,
		Fields:   fields,
		Children: nil,

		Position: [2]int{},
		Flags:    flags,
	}
}
