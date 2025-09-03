package compile

type Library struct {
	Version string             `yaml:"version"`
	Types   []string           `yaml:"types"`
	Funcs   map[string]FuncDef `yaml:"funcs"`
}

type FuncDef struct {
	Added     float64    `yaml:"added"`
	Types     []TypeDef  `yaml:"types,omitempty"`
	Arguments [][]string `yaml:"arguments,omitempty"`
	Returns   *string    `yaml:"returns"`
}

type TypeDef map[string]string

func (l *Library) LookupFunc(opid string) *FuncDef {
	if def, ok := l.Funcs[opid]; ok {
		return &def
	}
	return nil
}
