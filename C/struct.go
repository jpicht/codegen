package C

import "go/ast"

type Base struct {
	AST     *ast.TypeSpec
	Imports []*ast.Ident
}

type StructBuilder struct {
	Base
	S *ast.StructType
}

func Struct(name string) *StructBuilder {
	a := &ast.TypeSpec{}
	a.Name = &ast.Ident{Name: name}
	s := &ast.StructType{
		Fields: &ast.FieldList{},
	}
	a.Type = s
	return &StructBuilder{
		Base: Base{
			AST: a,
		},
		S: s,
	}
}

func (sb *StructBuilder) AddField(name string, Type TypeRef, Tag string) *StructBuilder {
	f := &ast.Field{
		Names: []*ast.Ident{{Name: name}},
		Type:  Type.Ident(),
	}
	if Tag != "" {
		f.Tag = &ast.BasicLit{Value: Tag}
	}
	sb.S.Fields.List = append(sb.S.Fields.List, f)
	return sb
}
