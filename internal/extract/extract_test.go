package extract

import (
	"github.com/OlegSchwann/GoDao/internal/template"
	"runtime"
	"testing"

	"github.com/OlegSchwann/GoDao/internal/ast"
	"github.com/OlegSchwann/GoDao/internal/flag"
)

var GoDao1 = template.DotType{PackageName: "testdata"}
var GoDao2 = template.DotType{PackageName: "testdata"}
var GoDao3 = template.DotType{PackageName: "testdata"}
var GoDao4 = template.DotType{PackageName: "testdata"}
var GoDao5 = template.DotType{PackageName: "testdata"}
var GoDao6 = template.DotType{PackageName: "testdata"}
var GoDao7 = template.DotType{PackageName: "testdata"}
var GoDao8 = template.DotType{PackageName: "testdata"}

func TestExtract(t *testing.T) {
	file, err := ast.ParseFile(flag.Config{
		InputGoFilePath:  "./test/database_size.go",
		OutputGoFilePath: "./test/database_size_dao.go",
		Verbose:          true,
	})
	if err != nil {
		t.Fatal(err)
	}

	gotDot, err := Extract(file)
	if err != nil {
		t.Fatal(err)
	}

	runtime.KeepAlive(gotDot)
}
