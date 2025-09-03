package tree

import "reflect"

type Value struct {
	Value     string `json:"value"`
	Type      string `json:"type"`
	HasQuotes bool   `json:"hasQuotes"`
}

func (v Value) GetType() reflect.Type {
	return reflect.TypeOf(v)
}
