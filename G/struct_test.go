package G_test

import (
	"bytes"
	"go/printer"
	"go/token"
	"testing"

	"github.com/jpicht/codegen/G"
	"github.com/stretchr/testify/assert"
)

var expected = `test struct {
	test bool json:",omitempty"
}`

func TestStruct(t *testing.T) {
	S := G.Type(G.Ident("test"), G.Struct(
		G.Field(G.Ident("test"), G.Ident("bool"), G.BasicLit(`json:",omitempty"`)),
	))

	B := &bytes.Buffer{}
	printer.Fprint(B, token.NewFileSet(), S)

	assert.Equal(t, expected, B.String())
}
