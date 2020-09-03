package updatevrecv

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "updatevrecv is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "updatevrecv",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	inspect.Preorder([]ast.Node{new(ast.FuncDecl)}, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			obj := pass.TypesInfo.ObjectOf(n.Name).(*types.Func)
			t := obj.Type().(*types.Signature)

			// exclude non method
			if t.Recv() == nil {
				return
			}

			// exclude methods with pointer receiver
			if _, ok := under(t.Recv().Type()).(*types.Pointer); ok {
				return
			}

			ast.Inspect(n, func(node ast.Node) bool {
				switch node := node.(type) {
				case *ast.AssignStmt:
					for _, lh := range node.Lhs {
						v, ok := lh.(*ast.SelectorExpr)
						if !ok {
							continue
						}

						selection, ok := pass.TypesInfo.Selections[v]
						if !ok {
							continue
						}

						if _, ok := under(selection.Recv()).(*types.Pointer); ok {
							continue
						}

						if types.Identical(selection.Recv(), t.Recv().Type()) {
							pass.Reportf(lh.Pos(), "field update in value receiver method")
						}
					}
				}
				return true
			})
		}
	})

	return nil, nil
}

func under(t types.Type) types.Type {
	if named, _ := t.(*types.Named); named != nil {
		return under(named.Underlying())
	}
	return t
}
