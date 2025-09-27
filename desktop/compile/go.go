package compile

import (
	"fmt"
	"iter"
	"log"
	"opticode/desktop/tree"
	"strings"
	"sync"
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

	mutex *sync.Mutex
}

func NewGoGenerator(t *tree.Tree, libs []*GoLibrary, indent string) *GoGenerator {
	return &GoGenerator{
		Tree:        t,
		indent:      indent,
		Libaries:    libs,
		NodeTracker: []string{},
		mutex:       &sync.Mutex{},
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

func (gg *GoGenerator) Next() iter.Seq2[string, *Error] {
	return func(yield func(string, *Error) bool) {
		for {
			if gg.NodeIndexID == GoExitId {
				log.Println("exit cause: _exit node")
				break
			}
			node, ok := gg.Tree.Nodes[gg.NodeIndexID]
			if !ok {
				log.Println("exit cause: unresolvable node")
				if !yield("", Err(Fatal, "cannot resolve next node: %s -> ? (code: 1)", gg.NodeIndexID)) {
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

func (gg *GoGenerator) Eval(node tree.Node, isParent ...bool) (string, *Error) {
	switch node.Opcode {
	case tree.OP_Package:
		log.Println("package")
		return gg.op_package(node)
	case tree.OP_Import:
		log.Println("import")
		return gg.op_import(node)

	// Increment & decrement
	case tree.OP_OperInc:
		log.Println("operInc")
		return gg.op_operInc(node)
	case tree.OP_OperDec:
		log.Println("operDec")
		return gg.op_operDec(node)
	// Assignment Operators
	case tree.OP_OperAssign:
		log.Println("operAssign")
		return gg.op_operAssign(node)
	case tree.OP_OperAddAssign:
		log.Println("operAddAssign")
		return gg.op_operAddAssign(node)
	case tree.OP_OperSubAssign:
		log.Println("operSubAssign")
		return gg.op_operSubAssign(node)
	case tree.OP_OperMulAssign:
		log.Println("operMulAssign")
		return gg.op_operMulAssign(node)
	case tree.OP_OperDivAssign:
		log.Println("operDivAssign")
		return gg.op_operDivAssign(node)
	case tree.OP_OperModAssign:
		log.Println("operModAssign")
		return gg.op_operModAssign(node)
	case tree.OP_OperBAndAssign:
		log.Println("operBAndAssign")
		return gg.op_operBAndAssign(node)
	case tree.OP_OperBOrAssign:
		log.Println("operBOrAssign")
		return gg.op_operBorAssign(node)
	case tree.OP_OperBXorAssign:
		log.Println("operBXorAssign")
		return gg.op_operBXorAssign(node)
	case tree.OP_OperLShiftAssign:
		log.Println("operLShiftAssign")
		return gg.op_operLShiftAssign(node)
	case tree.OP_OperRShiftAssign:
		log.Println("operRShiftAssign")
		return gg.op_operRShiftAssign(node)
	case tree.OP_OperBClearAssign:
		log.Println("operBCLearAssign")
		return gg.op_operBclearAssign(node)
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
		return gg.op_func(node)
	case tree.OP_FuncCall:
		log.Println("funcCall")
		return gg.op_funcCall(node, isParent...)
	case tree.OP_If:
		log.Println("if")
		return gg.op_if(node)
	case tree.OP_Return:
		log.Println("return")
		return gg.op_return(node)
	default:
		return "", Err(Warn, "invalid opcode: %d", node.Opcode)
	}
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

	for buf, err := range gg.Next() {
		if err != nil {
			if err.Code == Fatal {
				return "", err
			}
			log.Println(err.Error())
		}

		_, err := gg.Builder.WriteString(buf)
		if err != nil {
			return "", err
		}
		if gg.NodeIndexID == GoExitId {
			break
		}
	}

	for _, l := range gg.Libaries {
		l.Dispose()
	}
	return gg.Builder.String(), nil
}

func (gg *GoGenerator) finishNode(node tree.Node) (nextExists bool) {
	gg.NodeTracker = append(gg.NodeTracker, gg.NodeIndexID)
	if node.Next == nil {
		log.Printf("Finished node: %s, Next: null", gg.NodeIndexID)
		return false
	}
	log.Printf("Finished node: %s Next: %s", gg.NodeIndexID, *node.Next)
	gg.NodeIndexID = *node.Next
	return true
}

func (gg *GoGenerator) evalArgs(node tree.Node) ([]string, *Error) {
	var result []string = []string{}
	for _, f := range node.Fields {
		if f.Flags&tree.IsPointer != 0 {
			if node, ok := gg.Tree.Nodes[f.Value]; ok {
				n, err := gg.Eval(node)
				if err != nil {
					return nil, err
				}

				result = append(result, n)
			}
		} else {
			if f.Flags&tree.HasQuotes != 0 {
				result = append(result, "\""+f.Value+"\"")
			} else {
				result = append(result, f.Value)
			}
		}
	}
	return result, nil
}
