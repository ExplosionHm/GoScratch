package tree

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Tree struct {
	Assets []string        `json:"assets"` //! Implement
	Nodes  map[string]Node `json:"nodes,omitempty"`
}

func NewTree() *Tree {
	return &Tree{
		Assets: nil,
		Nodes:  map[string]Node{},
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
