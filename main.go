package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

// Пишем кодогенератор.
// Пишем, потом рефакторим, проверяем гипотезы.

// загружаем код
// парсим пакет
// находим нужные структуры
// находим в структурах нужные структуры
// формируем правило в виде структуры
// подбираем шаблон по этой структуре
// записываем в файлик с похожим расширением

// начнём с написания тестов.
// пример использования и результата, из которого можно потом наделать шаблонов.

func main() {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseFile(fset, "/home/max-dev/go/src/github.com/OlegSchwann/GoDao/test/2_insert/insert.go", nil, parser.ParseComments|parser.Trace|parser.AllErrors)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, decl := range pkgs.Decls {
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
			if strings.Contains(comment.Text, "goDao:generate") {
				goto haveLabel
			}
		}
		continue
	haveLabel:
		for _, spec := range decl.Specs {
			spec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			structType, ok := spec.Type.(*ast.StructType)
			if !ok {
				continue
			}
			if structType.Fields == nil {
				continue
			}
			for _, field := range structType.Fields.List {
				fmt.Printf("%#v", field)
			}
		}
	}
}

// задача:
// объяснить, как использовать нативные типы postgres через библиотеку.
// задача 2
// используя регулярку \b\$\d+::(\w+)\b для проверки типов во время шаблонизации.
// задача - исчерпывающее соответствие типов go и postgres.

// типы go ---> pgtype типы -+-> типы postges

// замечания
// в библиотеке pgx потерян тип time with timezone
// в библиотеке потерян тип char и character, который строка с фиксированной длинной, обрезающаяся при превышении без предупреждения.
// в библиотеке потерян тип xml
// в документации https://www.postgresql.org/docs/current/datatype.html потерян тип "char", который byte, есть только char, который n симполов.
// numeric и decimal не имеют конструктора из math.big,  не смотря на то, что используют его под капотом.
// в библиотеке отсутствует тип macaddr8
// в библиотеке отсутствует тип macaddr8[]
// в библиотеке отсутствует тип macaddr[]
