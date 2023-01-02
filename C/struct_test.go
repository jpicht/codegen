package C_test

import (
	"bytes"
	"go/printer"
	"go/token"
	"testing"

	"github.com/jpicht/codegen/C"
	"github.com/jpicht/codegen/T"
	"github.com/stretchr/testify/assert"
)

var expected = `test struct {
	test bool json:",omitempty"
}`

func TestStruct(t *testing.T) {
	fs := token.NewFileSet()

	S := C.Struct("test").
		AddField("test", T.Bool, `json:",omitempty"`).AST

	B := &bytes.Buffer{}
	printer.Fprint(B, fs, S)

	assert.Equal(t, expected, B.String())
}
