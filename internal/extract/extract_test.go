package extract

import (
	"runtime"
	"testing"

	"github.com/OlegSchwann/GoDao/internal/ast"
	"github.com/OlegSchwann/GoDao/internal/flag"
)

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
