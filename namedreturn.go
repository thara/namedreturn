package namedreturn

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "namedreturn finds named returns"

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "namedreturn",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	ps := make(map[types.Object]*ast.Ident)
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			correct(pass, n.Type, ps)
		}
	})

	suggest(pass, ps)

	return nil, nil
}

func correct(pass *analysis.Pass, typ *ast.FuncType, ps map[types.Object]*ast.Ident) {
	if typ.Results == nil {
		return
	}
	for _, f := range typ.Results.List {
		for _, n := range f.Names {
			if n == nil {
				return
			}
			obj := pass.TypesInfo.Defs[n]
			if obj != nil {
				ps[obj] = n
			}
		}
	}
}

func suggest(pass *analysis.Pass, ps map[types.Object]*ast.Ident) {
	for _, id := range ps {
		if id == nil {
			continue
		}

		fix := analysis.SuggestedFix{
			Message: fmt.Sprintf("%s is named return value, remove it", id.Name),
			TextEdits: []analysis.TextEdit{{
				Pos:     id.Pos(),
				End:     id.End(),
				NewText: []byte(""),
			}},
		}

		pass.Report(analysis.Diagnostic{
			Pos:            id.Pos(),
			End:            id.End(),
			Message:        fmt.Sprintf("%s is named return value", id.Name),
			SuggestedFixes: []analysis.SuggestedFix{fix},
		})
	}
}
