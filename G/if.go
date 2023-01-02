package G

import "go/ast"

func IF(condition ast.Expr, then *ast.BlockStmt, else_ ast.Stmt) *ast.IfStmt {
	return &ast.IfStmt{
		Cond: condition,
		Body: then,
		Else: else_,
	}
}
