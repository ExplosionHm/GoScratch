package golang

type TokenKind uint16

const (
	TokenEOF TokenKind = iota
	TokenSpace
	TokenNewLine
	TokenQuotation
	TokenParenLeft  // (
	TokenParenRight // )
	TokenTab
	TokenPackage
	TokenImport
)

var tokens = [][]byte{
	[]byte(""),
	[]byte(" "),
	[]byte("\n"),
	[]byte("\""),
	[]byte("("),
	[]byte(")"),
	[]byte("    "),
	[]byte("package"),
	[]byte("import"),
}

func (t TokenKind) Bytes() []byte {
	if int(t) >= len(tokens) {
		return nil
	}
	return tokens[t]
}

// JoinBytes efficiently concatenates multiple byte slices into one.
// It performs exactly ONE allocation, regardless of the number of slices.
func JoinBytes(slices ...[]byte) []byte {
	// Compute total length first
	totalLen := 0
	for _, s := range slices {
		totalLen += len(s)
	}

	// Allocate the final slice with exact capacity
	result := make([]byte, totalLen)

	// Copy each slice in sequence
	offset := 0
	for _, s := range slices {
		copy(result[offset:], s)
		offset += len(s)
	}

	return result
}
