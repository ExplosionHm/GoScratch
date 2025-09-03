package tree

type NodeOp int32

const (
	OP_assign NodeOp = iota
	OP_if
	OP_funcCall
	OP_func
)

type Node struct {
	Opcode NodeOp  `json:"opcode"`
	Opid   string  `json:"opid"`
	Fields []Value `json:"fields"`
	Parent *string `json:"parent"`
	Next   string  `json:"next"`

	// Only used if the node is a parent
	Position [2]int `json:"position,omitempty"`
	Flags    Flags  `json:"flags,omitempty"`
	Tab      string `json:"tab"`
}

func (n Node) IsParent() bool {
	return n.Parent == nil
}

func NewNode(op NodeOp, flags Flags, fields ...Value) Node {
	return Node{
		Opcode: op,
		Fields: fields,
		Parent: nil,
		Next:   "",

		Position: [2]int{},
		Flags:    flags,
	}
}
