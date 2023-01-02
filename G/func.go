package G

import "go/ast"

type FuncDeclHelper func(*ast.FuncDecl)

func Func(elements ...any) *ast.FuncDecl {
	f := &ast.FuncDecl{
		Type: &ast.FuncType{
			Params: &ast.FieldList{},
		},
	}
	for _, elem := range elements {
		switch tt := elem.(type) {
		case *ast.BlockStmt:
			f.Body = tt
		case *ast.Ident:
			f.Name = tt
		case FuncDeclHelper:
			tt(f)
		default:
			panic("what is dis?")
		}
	}
	return f
}

func Arguments(args ...*ast.Field) FuncDeclHelper {
	return func(fd *ast.FuncDecl) {
		fd.Type.Params = &ast.FieldList{List: args}
	}
}

func PointerReciever(ident *ast.Ident, Type ast.Expr) FuncDeclHelper {
	return Reciever(ident, Star(Type))
}

func Results(results ...*ast.Field) FuncDeclHelper {
	return func(fd *ast.FuncDecl) {
		fd.Type.Results = &ast.FieldList{List: results}
	}
}

func Reciever(ident *ast.Ident, Type ast.Expr) FuncDeclHelper {
	return func(fd *ast.FuncDecl) {
		fd.Recv = &ast.FieldList{
			List: []*ast.Field{
				Field(ident, Type, nil),
			},
		}
	}
}
