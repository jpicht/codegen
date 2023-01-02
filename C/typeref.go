package C

import "go/ast"

type TypeRef struct {
	Name    string
	Package string
}

func (tr *TypeRef) Ident() *ast.Ident {
	return &ast.Ident{Name: tr.Name}
}

func (tr *TypeRef) Import() (*ast.ImportSpec, bool) {
	if tr.Package == "" {
		return nil, false
	}
	return &ast.ImportSpec{Path: &ast.BasicLit{Value: tr.Package}}, true
}

var (
	Bool    = TypeRef{"bool", ""}
	Float32 = TypeRef{"float32", ""}
	Float64 = TypeRef{"float64", ""}
	Int     = TypeRef{"int", ""}
	Int8    = TypeRef{"int8", ""}
	Int16   = TypeRef{"int16", ""}
	Int32   = TypeRef{"int32", ""}
	Int64   = TypeRef{"int64", ""}
	UInt    = TypeRef{"uint", ""}
	UInt8   = TypeRef{"uint8", ""}
	UInt16  = TypeRef{"uint16", ""}
	UInt32  = TypeRef{"uint32", ""}
	UInt64  = TypeRef{"uint64", ""}
	String  = TypeRef{"string", ""}
)
