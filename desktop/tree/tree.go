package tree

import (
	"encoding/json"
	"fmt"
	"io"
	"opticode/desktop/log"
	"os"
)

type Tree struct {
	Assets []string        `json:"assets"` //! Implement
	Nodes  map[string]Node `json:"nodes,omitempty"`
	logger *log.Logger     `json:"-"`
}

func NewTree(l *log.Logger) *Tree {
	return &Tree{
		Assets: nil,
		Nodes:  map[string]Node{},
		logger: l,
	}
}

// JS safe
func (t *Tree) OpenFile(path string) error {
	if t == nil {
		return fmt.Errorf("cannot open file: tree is nil")
	}
	file, err := os.OpenFile(path, 0, os.ModeAppend)
	if err != nil {
		return err
	}
	contents, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	var treeData Tree
	json.Unmarshal(contents, &treeData)

	*t = treeData
	return nil
}

func (t *Tree) LookUp() {

}
