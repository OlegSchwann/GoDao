package flag

import (
	"reflect"
	"testing"
)

func TestShellParsing1(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name     string
		args     args
		wantConf Config
		wantErr  bool
	}{{
		name: "1 file",
		args: args{
			args: []string{"godao", "./file.go"},
		},
		wantConf: Config{
			InputGoFilePath:  "./file.go",
			OutputGoFilePath: "./file_dao.go",
			Verbose:          false,
		},
		wantErr: false,
	}, {
		name: "debug",
		args: args{
			args: []string{"godao", "-verbose=true", "./package/file.go"},
		},
		wantConf: Config{
			InputGoFilePath:  "./package/file.go",
			OutputGoFilePath: "./package/file_dao.go",
			Verbose:          true,
		},
		wantErr: false,
	}, {
		name: "to many arguments",
		args: args{
			args: []string{"godao", "./sql.go", "./package/file.go"},
		},
		wantErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotConf, err := ShellParsing(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShellParsing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotConf, tt.wantConf) {
				t.Errorf("ShellParsing() gotConf = %v, want %v", gotConf, tt.wantConf)
			}
		})
	}
}
