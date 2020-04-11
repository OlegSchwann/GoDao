package extract

import (
	"fmt"
	"github.com/OlegSchwann/GoDao/internal/template"
	"go/ast"
	"go/token"
	"runtime"
	"strings"
)

func Extract(file *ast.File) (dot template.DotType, err error) {
	// copy package name
	dot.PackageName = file.Name.Name

	// copy all used packages directly
	// TODO: Copy all used packages from types, not all from file
	for _, i := range file.Imports {
		dot.Packages = append(dot.Packages, i.Path.Value)
	}

	for _, daoStruct := range goDaoStructs(file) {
		fmt.Printf("%#v", daoStruct)

		for _, field := range daoStruct.Fields.List {
			// TODO: Fill in the structure values correctly.
			// template.Function{
			//	Name: field.Names[0].Name,
			//	SQL: field.Tag.Value,
			//  ??
			// }
			runtime.KeepAlive(field)
		}
	}

	return dot, err
}

// search for "godao" struct in file
func goDaoStructs(file *ast.File) (structs []ast.StructType) {
	for _, decl := range file.Decls {
		decl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		if decl.Tok != token.TYPE {
			continue
		}
		if decl.Doc == nil {
			continue
		}
		for _, comment := range decl.Doc.List {
			if strings.Contains(comment.Text, "goDao") && strings.Contains(comment.Text, "generate") {
				goto haveLabel
			}
		}
		continue
	haveLabel: // If there's the "goDao: generate" somewhere in the comments before the structure:
		for _, spec := range decl.Specs {
			spec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			structType, ok := spec.Type.(*ast.StructType)
			if !ok {
				continue
			}
			structs = append(structs, *structType)
		}
	}
	return structs
}
