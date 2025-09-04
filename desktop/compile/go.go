package compile

import (
	"fmt"
	"iter"
	"opticode/desktop/tree"
	"os"
	"os/exec"
	"path"
	"strings"
)

const GoEntryId = "_entry"
const GoExitId = "_exit"

func ValidateGoTree(t *tree.Tree) error {
	// ! Implement
	return nil
}

func InitGoProject(dir string, name string) error {
	info, err := os.Stat(dir)
	if err != nil {
		return fmt.Errorf("cannot access dir %q: %w", dir, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("path %q is not a directory", dir)
	}

	if name == "" {
		return fmt.Errorf("module name cannot be empty")
	}

	cmd := exec.Command("go", "mod", "init", name)
	cmd.Dir = dir

	std, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go mod init failed: %w\nOutput:\n%s", err, string(std))
	}
	println("std out:", string(std))

	mainPath := path.Join(dir, "main.go")
	if _, err := os.Stat(mainPath); err == nil {
		return fmt.Errorf("main.go already exists at %q", mainPath)
	}

	f, err := os.Create(mainPath)
	if err != nil {
		return fmt.Errorf("failed to create main.go: %w", err)
	}
	defer f.Close()

	return nil
}

type GoGenerator struct {
	Tree        *tree.Tree
	Libaries    []*GoLibrary
	Builder     strings.Builder
	Indent      string
	NodeTracker []string
	NodeIndexID string
}

func NewGoGenerator(t *tree.Tree, libs []*GoLibrary, indent string) *GoGenerator {
	return &GoGenerator{
		Tree:        t,
		Indent:      indent,
		Libaries:    libs,
		NodeTracker: []string{},
	}
}

func (gg *GoGenerator) Next() iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		node, ok := gg.Tree.Nodes[gg.NodeIndexID]
		if !ok {
			yield("", fmt.Errorf("cannot resolve next node: %s -> ?", gg.NodeIndexID))
			return
		}

		switch node.Opcode {
		case tree.OP_funcCall:
			yield(gg.op_funcCall(node))
			return
		}
	}
}

func (gg *GoGenerator) finishNode() error {
	gg.NodeTracker = append(gg.NodeTracker, gg.NodeIndexID)

	node, ok := gg.Tree.Nodes[gg.NodeIndexID]
	if !ok {
		return fmt.Errorf("cannot resolve next node: %s -> ?", gg.NodeIndexID)
	}
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

		gg.Builder.WriteString(buf)

		if gg.NodeIndexID == GoExitId {
			break
		}
	}

	gg.Builder.WriteString("}")

	return gg.Builder.String(), nil
}
