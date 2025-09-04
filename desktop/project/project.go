package project

import "opticode/desktop/tree"

type LibraryPathFlag uint32

type ProjectHeader struct {
	Name      string
	Id        string
	Libraries map[string]LibraryPathFlag
}

type Project struct {
	ProjectHeader
	Tree *tree.Tree
}

func NewProject(header ProjectHeader, dir string) *Project {
	return &Project{
		header,
		tree.NewTree(),
	}
}
