package tree

import "reflect"

type ValueFlag int32

const (
	HasQuotes ValueFlag = 1 << iota
	IsPointer           // Points to another node

	IsCondition // Inside if statement condition
	IsIfBody    // Inside if statement body
	IsElseBody  // Inside else statement body

	IsFuncRef
	IsFuncTypes
	IsFuncArg
	IsFuncBody
)

type Value struct {
	Value string    `json:"value"`
	Type  string    `json:"type"`
	Flags ValueFlag `json:"flags"`
}

func (v Value) GetType() reflect.Type {
	return reflect.TypeOf(v)
}
