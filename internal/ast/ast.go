package ast

import (
	"github.com/OlegSchwann/GoDao/internal/flag"
	"go/ast"
	"go/parser"
	"go/token"
)

func ParseFile(config flag.Config) (*ast.File, error) {
	var maybeTrace parser.Mode // no effect by default
	if config.Verbose {
		maybeTrace = parser.Trace
	}

	return parser.ParseFile(token.NewFileSet(), config.InputGoFilePath, nil, parser.ParseComments|maybeTrace|parser.AllErrors)
}
