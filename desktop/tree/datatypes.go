package tree

type DataType int8

const (
	DT_int int8 = iota
	DT_int8
	DT_int16
	DT_int32
	DT_int64

	DT_uint
	DT_uint8
	DT_uint16
	DT_uint32
	DT_uint64
	DT_uintptr

	DT_string
	// Add other data types
)
