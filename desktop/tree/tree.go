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

func (t *Tree) OpenFile(path string) error {
	if t == nil {
		return fmt.Errorf("cannot open file: tree is nil")
	}
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file does not exist")
		}
		return fmt.Errorf("an error has occurred while checking path")
	}
	if info.IsDir() {
		return fmt.Errorf("path has to point to a file")
	}
	file, err := os.OpenFile(path, 0, 0)
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
