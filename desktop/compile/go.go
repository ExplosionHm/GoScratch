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
				yield("", fmt.Errorf("cannot resolve next node: %s -> ? (code: 1)", gg.NodeIndexID))
				return
			}
			log.Println("Processing node: ", gg.NodeIndexID, node.Opcode)

			switch node.Opcode {
			case tree.OP_FuncCall:
				if !yield(gg.op_funcCall(node)) {
					return
				}
			case tree.OP_If:
				if !yield(gg.op_if(node)) {
					return
				}
			}
		}
	}
}

func (gg *GoGenerator) finishNode() error {
	gg.NodeTracker = append(gg.NodeTracker, gg.NodeIndexID)

	node, ok := gg.Tree.Nodes[gg.NodeIndexID]
	if !ok {
		return nil
	}
	log.Printf("Finished node: %s Next: %s", gg.NodeIndexID, node.Next)
	gg.NodeIndexID = node.Next
	return nil
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
	if err := gg.finishNode(); err != nil {
		return "", err
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

	return gg.Builder.String(), nil
}
