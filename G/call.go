package G

import "go/ast"

func Call(ident *ast.Ident, args ...ast.Expr) *ast.CallExpr {
	return &ast.CallExpr{
		Fun:  ident,
		Args: args,
	}
}
