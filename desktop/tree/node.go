package tree

type NodeOp int32

const (
	OP_Package NodeOp = iota

	// Arithmetic Operators

	OP_OperAdd
	OP_OperSub
	OP_OperMul
	OP_OperDiv
	OP_OperMod

	// Increment & decrement

	OP_OperInc
	OP_OperDec

	// Assignment Operators

	OP_OperAssign
	OP_OperAddAssign
	OP_OperSubAssign
	OP_OperMulAssign
	OP_OperDivAssign
	OP_OperModAssign
	OP_OperBAndAssign
	OP_OperBOrAssign
	OP_OperBXorAssign
	OP_OperLShiftAssign
	OP_OperRShiftAssign
	OP_OperBClearAssign

	// Comparison Operators

	OP_OperEqual
	OP_OperNotEqual
	OP_OperLT
	OP_OperLTorEqual
	OP_OperGT
	OP_OperGTorEqual

	// Logical Operators

	OP_OperAnd
	OP_OperOr
	OP_OperNot

	// Bitwise Operators

	OP_OperBAnd
	OP_OperBOr
	OP_OperBXor
	OP_OperBClear
	OP_OperLShift
	OP_OperRShift

	OP_PointerRef  // "&"
	OP_PointerDref // "*"

	OP_ChanLeft  // "<-"
	Op_ChanRight // "->"

	OP_Func
	OP_FuncCall

	OP_If
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
