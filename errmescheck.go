package errmescheck

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "errmescheck is check difference 'xxx failed()' and 'failed xxx'"

const targetPackage = "github.com/pkg/errors"
const targetType = targetPackage + "."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "errmescheck",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.ReturnStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		if returns, ok := n.(*ast.ReturnStmt); ok {
			for _, e := range returns.Results {
				if c, ok := e.(*ast.CallExpr); ok {
					// check package and func name
					if s, ok := c.Fun.(*ast.SelectorExpr); ok {
						if s.Sel.Name != "Wrap" {
							return
						}
						if i, ok := s.X.(*ast.Ident); ok {
							if i.Name != "errors" {
								return
							}
						}
					}

					// check error message
					if len(c.Args) == 0 {
						return
					}
					errExpr := c.Args[len(c.Args)-1]
					if b, ok := errExpr.(*ast.BasicLit); ok {
						if !strings.HasPrefix(strings.Trim(b.Value, "\""), "failed to") {
							pass.Reportf(b.Pos(), "The prefix of the error message should be 'failed to ...'")
						}
					}
				}
			}
		}
	})

	return nil, nil
}
