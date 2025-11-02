package project

import (
	"encoding/json"
	"os"
)

type LibraryPathFlag uint32

type Header struct {
	Name          string                     `json:"name"`
	Id            string                     `json:"id"`
	Owner         string                     `json:"owner"`
	Collaborators []string                   `json:"collaborators"`
	Language      string                     `json:"language"`
	Libraries     map[string]LibraryPathFlag `json:"libraries"`
}

type Project struct {
	Header      `json:"header"`
	Appearances `json:"appearances"`
	Assets      []string `json:"assets"`

	Program []byte
}

type Appearances struct {
	Theme int              `json:"theme"`
	Nodes []NodeAppearance `json:"nodes"`
}

type NodeAppearance struct {
	Color uint32
}

func NewProject(header Header, appearances Appearances, assets ...string) *Project {
	return &Project{
		Header:      header,
		Appearances: appearances,
		Program:     nil,
		Assets:      assets,
	}
}

func (p *Project) UpdateProgramFile(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	p.Program = content
	return nil
}

func (p *Project) ToJson() ([]byte, error) {
	return json.Marshal(p)
}

func LoadProject(data []byte) *Project {
	//! Should add safety check
	var project *Project
	err := json.Unmarshal(data, project)
	if err != nil {
		return nil
	}
	return project
}
