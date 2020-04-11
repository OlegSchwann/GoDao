package ast

import (
	"github.com/OlegSchwann/GoDao/internal/flag"
	"testing"
)

func TestParseFile(t *testing.T) {
	type args struct {
		config flag.Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{{
		name: "parse file",
		args: args{
			config: flag.Config{
				InputGoFilePath:  "./test/cube_root.go",
				OutputGoFilePath: "./test/cube_root_dao.go",
				Verbose:          false,
			},
		},
		wantErr: false,
	}, {
		name: "parse file",
		args: args{
			config: flag.Config{
				InputGoFilePath:  "./test/not_exist.go",
				OutputGoFilePath: "./test/not_exist_dao.go",
				Verbose:          false,
			},
		},
		wantErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFile(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got == nil || len(got.Decls) != 1 {
				t.Errorf("expected valid AST with 1 declaration, got error.")
			}
		})
	}
}
