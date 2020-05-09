package extract

import (
	"errors"
	"github.com/OlegSchwann/GoDao/internal/template"
	"go/ast"
	"go/token"
	"strconv"
	"strings"
)

func Extract(file *ast.File) (dot template.DotType, err error) {
	// copy package name
	dot.PackageName = file.Name.Name

	// copy all used packages directly
	// TODO: Copy all used packages from types, not all from file
	for _, i := range file.Imports {
		importedPackage, err := strconv.Unquote(i.Path.Value)
		if err != nil {
			return template.DotType{}, errors.New("")
		}
		dot.Packages = append(dot.Packages, importedPackage)
	}

	for _, daoStruct := range goDaoStructs(file) {
		for _, field := range daoStruct.Fields.List {

			// TODO: Fill in the structure values correctly.
			var function template.Function

			function.Name = field.Names[0].Name // TODO: check != ""

			if function.SQL, err = strconv.Unquote(field.Tag.Value); err != nil { /* TODO: check  != "" */
				return template.DotType{}, err
			}

			fieldFunc, ok := field.Type.(*ast.FuncType)
			if !ok {
				return template.DotType{}, errors.New("field Func should be func")
			}

			for _, param := range fieldFunc.Params.List {
				variable := template.Variable{}

				variable.Name = param.Names[0].Name

				switch ident := param.Type.(type) {
				case *ast.Ident:
					variable.Type = ident.Name
				case *ast.SelectorExpr:
					variable.Type = ident.X.(*ast.Ident).Name + "." + ident.Sel.Name
				default:
					return template.DotType{}, errors.New("unknown type")
				}
				function.InputArguments = append(function.InputArguments, variable)
			}

			function.ReturnValueType, err = selectMode(fieldFunc)
			if err != nil {
				return template.DotType{}, err
			}

			if function.ReturnValueType == template.Exec {
				continue
			}

			if function.ReturnValueType == template.QueryRow {
				for _, result := range fieldFunc.Results.List{
					template.Variable{
						Name: result.Names[0].Name, // TODO: check name exist.
						Type: result.Type // TODO: *ast.Ident and *ast.SelectorExpr
					}
					function.OutputArguments
				}
			}

			dot.Functions = append(dot.Functions, function)
			// TODO: for usability reasons, аll "." operators shoud be check for errors.
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
			if strings.Contains(comment.Text, "GoDao") && strings.Contains(comment.Text, "generate") {
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

// Функция выдаёт тип шаблонизации.
// если нет возвращаемых параметров - template.Exec
// если есть несколько возвращаемых параметров - template.QueryStruct
// если слайс структур - template.QueryStruct
func selectMode(fieldFunc *ast.FuncType) (returnType template.ReturnValueType, err error) {
	// Check (err error)
	if l := len(fieldFunc.Results.List); !(l > 0 && fieldFunc.Results.List[l-1].Names[0].Name == "err" && fieldFunc.Results.List[l-1].Type.(*ast.Ident).Name == "error") {
		return returnType, errors.New("function must have err err as last argument")
	}

	results := fieldFunc.Results.List[0 : len(fieldFunc.Results.List)-1]
	if len(results) == 0 {
		return template.Exec, nil
	}

	if /* first variable in results is a a slice of any struct */ func() (isSliceOfStruct bool) {
		defer func() {
			if recover() != nil {
				isSliceOfStruct = false
			}
		}()
		return results[0].Type.(*ast.ArrayType).Elt.(*ast.Ident).Obj.Decl.(*ast.TypeSpec).Type.(*ast.StructType) != nil
	}() {
		return template.QueryStruct, nil
	}

	return template.QueryRow, nil
}
