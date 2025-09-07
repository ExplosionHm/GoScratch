package compile

import (
	"fmt"
	"iter"
	"log"
	"opticode/desktop/tree"
	"strings"
)

const GoEntryId = "_entry"
const GoExitId = "_exit"

func ValidateGoTree(t *tree.Tree) error {
	// ! Implement
	return nil
}

type GoGenerator struct {
	Tree        *tree.Tree
	Libaries    []*GoLibrary
	Builder     strings.Builder
	FuncDefs    []string
	indent      string
	indentMul   int
	NodeTracker []string
	NodeIndexID string
}

func NewGoGenerator(t *tree.Tree, libs []*GoLibrary, indent string) *GoGenerator {
	return &GoGenerator{
		Tree:        t,
		indent:      indent,
		Libaries:    libs,
		NodeTracker: []string{},
	}
}

func (gg *GoGenerator) SetIndent(mul int) {
	gg.indentMul = mul
}

func (gg *GoGenerator) IncreaseIndent() {
	gg.indentMul++
}

func (gg *GoGenerator) DecreaseIndent() {
	if gg.indentMul <= 0 {
		return
	}
	gg.indentMul--
}

func (gg *GoGenerator) Indent() string {
	return strings.Repeat(gg.indent, gg.indentMul)
}

func (gg *GoGenerator) Next() iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		for {
			if gg.NodeIndexID == GoExitId {
				break
			}
			node, ok := gg.Tree.Nodes[gg.NodeIndexID]
			if !ok {
				if !yield("", fmt.Errorf("cannot resolve next node: %s -> ? (code: 1)", gg.NodeIndexID)) {
					return
				}
			}
			log.Println("Processing node:", gg.NodeIndexID, node.Opcode)
			log.Println(node)
			if !yield(gg.Eval(node)) {
				return
			}
		}
	}
}

func (gg *GoGenerator) Eval(node tree.Node) (string, error) {
	switch node.Opcode {
	// Comparison Operators
	case tree.OP_OperEqual:
		log.Println("operEqual")
		return gg.op_operEqual(node)
	case tree.OP_OperNotEqual:
		log.Println("operNotEqual")
		return gg.op_operNotEqual(node)
	case tree.OP_OperLT:
		log.Println("operLT")
		return gg.op_operLT(node)
	case tree.OP_OperLTorEqual:
		log.Println("operLTorEqual")
		return gg.OP_OperLTorEqual(node)
	case tree.OP_OperGT:
		log.Println("operGT")
		return gg.op_operGT(node)
	case tree.OP_OperGTorEqual:
		log.Println("operGTorEqual")
		return gg.op_operGTorEqual(node)
		// Logical Operators
	case tree.OP_OperAnd:
		log.Println("operAnd")
		return gg.op_operAnd(node)
	case tree.OP_OperOr:
		log.Println("operOr")
		return gg.op_operOr(node)
	case tree.OP_OperNot:
		log.Println("operNot")
		return gg.op_operNot(node)
	case tree.OP_Func:
		log.Println("func")
		def, err := gg.op_func(node)
		if err != nil {
			return "", err
		}
		gg.FuncDefs = append(gg.FuncDefs, def)
		return "", nil
	case tree.OP_FuncCall:
		log.Println("funcCall")
		return gg.op_funcCall(node)
	case tree.OP_If:
		log.Println("if")
		return gg.op_if(node)
	default:
		return "", fmt.Errorf("invalid opcode: %d", node.Opcode)
	}
}

func (gg *GoGenerator) finishNode() (nextExists bool) {
	gg.NodeTracker = append(gg.NodeTracker, gg.NodeIndexID)

	node, ok := gg.Tree.Nodes[gg.NodeIndexID]
	if !ok {
		return false
	}
	log.Printf("Finished node: %s Next: %s", gg.NodeIndexID, node.Next)
	gg.NodeIndexID = node.Next
	return true
}

func (gg *GoGenerator) Generate() (string, error) {
	if len(gg.Tree.Nodes) == 0 {
		return "", fmt.Errorf("cannot find entry: tree is empty")
	}
	_, ok := gg.Tree.Nodes[GoEntryId]
	if !ok {
		return "", fmt.Errorf("cannot find entry")
	}

	gg.NodeIndexID = GoEntryId

	gg.Builder.WriteString("package main\n\n")
	gg.Builder.WriteString("//_functions")
	gg.Builder.WriteString("\nfunc main() {\n")
	if nextExists := gg.finishNode(); !nextExists {
		return "", fmt.Errorf("failed to compile: program too short")
	}

	for buf, err := range gg.Next() {
		if err != nil {
			return "", err
		}

		gg.Builder.WriteString(buf) //TODO: Handler returns

		if gg.NodeIndexID == GoExitId {
			break
		}
	}

	gg.Builder.WriteString("}")
	log.Printf("length: %d", len(gg.FuncDefs))
	for _, def := range gg.FuncDefs {
		log.Println("TESTTSTSTWFW")
		log.Println("Adding func defs")
		gg.Builder.WriteString("\n" + def)
	}

	for _, l := range gg.Libaries {
		l.Dispose()
	}
	return gg.Builder.String(), nil
}
