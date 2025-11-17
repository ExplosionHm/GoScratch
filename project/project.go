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
	Misc          map[string]interface{}     `json:"misc"`
}

type Appearances struct {
	Theme int              `json:"theme"`
	Nodes []NodeAppearance `json:"nodes"`
}

type NodeAppearance struct {
	Color uint32
	// Add more attributes
}

type Project struct {
	Header      `json:"header"`
	Appearances `json:"appearances"`
	Assets      []string `json:"assets"`

	Program string `json:"program"`
}

func NewProject(header Header, appearances Appearances, program string, assets ...string) *Project {
	return &Project{
		Header:      header,
		Appearances: appearances,
		Program:     "",
		Assets:      assets,
	}
}

func (p *Project) UpdateProgramFile(buf []byte) error {
	// Check if path exists before writing
	if _, err := os.Stat(p.Program); err != nil {
		return err
	}

	// Write buffer to program file
	os.WriteFile(p.Program, buf, 0)
	return nil
}

func (p *Project) ToJson() ([]byte, error) {
	return json.Marshal(p)
}

func LoadProject(data []byte) (*Project, error) {
	//! Should add safety check
	var project *Project
	err := json.Unmarshal(data, project)
	return project, err
}

func LoadProjectFromFile(path string) (*Project, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return LoadProject(buf)
}
