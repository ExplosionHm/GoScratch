package tree

type Flags int

const (
	HintTest1 Flags = 1 << iota
	HintTest2
	HintTest3
	HintTest4
)
