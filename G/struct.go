package G

import (
	"go/ast"
)

func Field(ident *ast.Ident, _type ast.Expr, tag *ast.BasicLit) *ast.Field {
	return &ast.Field{Names: []*ast.Ident{ident}, Type: _type, Tag: tag}
}

func Struct(fields ...*ast.Field) *ast.StructType {
	return &ast.StructType{
		Fields: &ast.FieldList{
			List: fields,
		},
	}
}
