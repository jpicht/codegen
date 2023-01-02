package G

import "go/ast"

func BasicLit(lit string) *ast.BasicLit {
	return &ast.BasicLit{Value: lit}
}

func Ident(ident string) *ast.Ident {
	return &ast.Ident{Name: ident}
}

func Star(x ast.Expr) ast.Expr {
	return &ast.StarExpr{
		X: x,
	}
}

func Type(ident *ast.Ident, expr ast.Expr) *ast.TypeSpec {
	return &ast.TypeSpec{Name: ident, Type: expr}
}
