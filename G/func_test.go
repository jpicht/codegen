package G_test

import (
	"bytes"
	"go/printer"
	"go/token"
	"testing"

	"github.com/jpicht/codegen/G"
	"github.com/stretchr/testify/assert"
)

var func_expected = `func (x *Y) testfn(p1 string, p2 int) (r1 string, r2 bool) {
}`

func TestFunc(t *testing.T) {
	F := G.Func(
		G.Ident("testfn"),
		G.PointerReciever(G.Ident("x"), G.Ident("Y")),
		G.Arguments(
			G.Field(G.Ident("p1"), G.String, nil),
			G.Field(G.Ident("p2"), G.Int, nil),
		),
		G.Results(
			G.Field(G.Ident("r1"), G.String, nil),
			G.Field(G.Ident("r2"), G.Bool, nil),
		),
		G.Block(),
	)

	B := &bytes.Buffer{}
	printer.Fprint(B, token.NewFileSet(), F)

	assert.Equal(t, func_expected, B.String())
}
