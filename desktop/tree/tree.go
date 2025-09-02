package tree

type Tree struct {
	Assets any    `json:"-"` //! Implement
	Nodes  []Node `json:"nodes,omitempty"`
}

func NewTree() *Tree {
	return &Tree{
		Assets: nil,
		Nodes:  []Node{},
	}
}

func (*Tree) GenerateCode(outdir string) {

}
