package pbcheck

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"honnef.co/go/tools/analysis/code"
	"honnef.co/go/tools/analysis/report"
)

var Analyzer = &analysis.Analyzer{
	Name:     "pbcheck",
	Doc:      "pbcheck",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

var pkgPrefix string // -name flag

func init() {
	Analyzer.Flags.StringVar(&pkgPrefix, "prefix", pkgPrefix, "the proto pacakge prefix to check")
}

// const packagePrefix = "(*github.com/lyqscmy/pbcheck/proto"

func run(pass *analysis.Pass) (interface{}, error) {
	safe := make(map[token.Pos]struct{})
	pkgPrefix = "(*" + pkgPrefix

	fn := func(node ast.Node) {
		switch e := node.(type) {
		case *ast.AssignStmt:
			for _, v := range e.Lhs {
				if sel, ok := v.(*ast.SelectorExpr); ok {
					name := code.SelectorName(pass, sel)
					if strings.HasPrefix(name, pkgPrefix) {
						safe[sel.Pos()] = struct{}{}
					}
				}
			}
		case *ast.SelectorExpr:
			if _, ok := safe[e.Pos()]; ok {
				return
			}
			name := code.SelectorName(pass, e)
			if strings.HasPrefix(name, pkgPrefix) && !strings.HasPrefix(e.Sel.Name, "Get") {
				report.Report(pass, e, fmt.Sprintf("Use: %s.Get%s()", e.X, e.Sel.Name))
			}
		}
	}
	code.Preorder(pass, fn, (*ast.AssignStmt)(nil), (*ast.SelectorExpr)(nil))
	return nil, nil
}
