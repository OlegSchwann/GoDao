package extract

import (
	"github.com/OlegSchwann/GoDao/internal/ast"
	"github.com/OlegSchwann/GoDao/internal/flag"
	"github.com/OlegSchwann/GoDao/internal/template"
	"reflect"
	"testing"
)

type args struct {
	Path string
}

type test struct {
	name    string
	args    args
	wantDot template.DotType
	wantErr bool
}

func TestExtract1(t *testing.T) {
	tt := test{
		name:    "1 The absence of an \"err error\" as the last parameter is incorrect.",
		args:    args{Path: "./testdata/1.go"},
		wantDot: template.DotType{},
		wantErr: true,
	}
	t.Run(tt.name, func(t *testing.T) {
		file, err := ast.ParseFile(flag.Config{
			InputGoFilePath:  tt.args.Path,
			OutputGoFilePath: "/dev/null",
			Verbose:          false,
		})
		if err != nil {
			t.Fatal("ast.ParseFile(\""+tt.args.Path+"\"): ", err)
		}
		gotDot, err := Extract(file)
		if (err != nil) != tt.wantErr {
			t.Errorf("Extract() error = %#v, wantErr %#v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(gotDot, tt.wantDot) {
			t.Errorf("Extract() gotDot = %#v, want %#v", gotDot, tt.wantDot)
		}
	})
}
func TestExtract2(t *testing.T) {
	tt := test{
		name: "2 Checking the template.Exec mode when the function signature = \"func() (err error)\".",
		args: args{Path: "./testdata/2.go"},
		wantDot: template.DotType{
			PackageName: "testdata",
			Packages:    []string{},
			Functions: []template.Function{{
				Name: "DropTestDatabase",
				// language=PostgreSQL
				SQL: `
        drop database if exists "test";`,
				InputArguments:  []template.Variable{},
				ReturnValueType: template.Exec,
			}},
		},
		wantErr: false,
	}
	t.Run(tt.name, func(t *testing.T) {
		file, err := ast.ParseFile(flag.Config{
			InputGoFilePath:  tt.args.Path,
			OutputGoFilePath: "/dev/null",
			Verbose:          false,
		})
		if err != nil {
			t.Fatal("ast.ParseFile(\""+tt.args.Path+"\"): ", err)
		}
		gotDot, err := Extract(file)
		if (err != nil) != tt.wantErr {
			t.Errorf("Extract() error = %#v, wantErr %#v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(gotDot, tt.wantDot) {
			t.Errorf("Extract() gotDot = %#v, want %#v", gotDot, tt.wantDot)
		}
	})
}
func TestExtract3(t *testing.T) {
	tt := test{
		name: "3 Checking incoming native parameters; checking outgoing composite parameters.",
		args: args{Path: "./testdata/3.go"},
		wantDot: template.DotType{
			PackageName: "testdata",
			Packages:    []string{"github.com/jackc/pgtype"},
			Functions: []template.Function{{
				Name: "GetSettings",
				// language=PostgreSQL
				SQL: `
        with "tmp"("k", "v") as (values
            (0::int8, '{"dark_theme": true}'::json),
            (1::int8, '{"cookies": false}'::json)
        ) select "v"
        from "tmp"
        where "k" = $1
        limit 1;`,
				InputArguments:  []template.Variable{{"id", "int64"}},
				ReturnValueType: template.QueryRow,
				OutputArguments: []template.Variable{{"json", "pgtype.JSON"}},
			}},
		},
		wantErr: false,
	}
	t.Run(tt.name, func(t *testing.T) {
		file, err := ast.ParseFile(flag.Config{
			InputGoFilePath:  tt.args.Path,
			OutputGoFilePath: "/dev/null",
			Verbose:          false,
		})
		if err != nil {
			t.Fatal("ast.ParseFile(\""+tt.args.Path+"\"): ", err)
		}
		gotDot, err := Extract(file)
		if (err != nil) != tt.wantErr {
			t.Errorf("Extract() error = %#v, wantErr %#v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(gotDot, tt.wantDot) {
			t.Errorf("Extract() gotDot = %#v, want %#v", gotDot, tt.wantDot)
		}
	})
}

func TestExtract4(t *testing.T) {
	tt := test{
		name: "4 Check imported input parameters; check imported outgoing parameters.",
		args: args{Path: "./testdata/4.go"},
		wantDot: template.DotType{
			PackageName: "testdata",
			Packages:    []string{"github.com/jackc/pgtype"},
			Functions: []template.Function{{
				Name: "NamedDatabaseSize",
				// language=PostgreSQL
				SQL: `
        select pg_database_size($1::name);`,
				InputArguments:  []template.Variable{{"name", "pgtype.Name"}},
				ReturnValueType: template.QueryRow,
				OutputArguments: []template.Variable{{"bytes", "int64"}},
			}},
		},
		wantErr: false,
	}
	t.Run(tt.name, func(t *testing.T) {
		file, err := ast.ParseFile(flag.Config{
			InputGoFilePath:  tt.args.Path,
			OutputGoFilePath: "/dev/null",
			Verbose:          false,
		})
		if err != nil {
			t.Fatal("ast.ParseFile(\""+tt.args.Path+"\"): ", err)
		}
		gotDot, err := Extract(file)
		if (err != nil) != tt.wantErr {
			t.Errorf("Extract() error = %#v, wantErr %#v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(gotDot, tt.wantDot) {
			t.Errorf("Extract() gotDot = %#v, want %#v", gotDot, tt.wantDot)
		}
	})
}
func TestExtract5(t *testing.T) {
	tt := test{
		name: "5 Test arrays, input and output.",
		args: args{Path: "./testdata/5.go"},
		wantDot: template.DotType{
			PackageName: "testdata",
			Functions: []template.Function{{
				Name: "Delta",
				// language=PostgreSQL
				SQL: `
		select array(
		    select "tmp"."to" - "tmp"."from"
		    from unnest(0 || $1::int4[], $1::int4[] || 0) as "tmp"("from", "to")
		)::int4[];`,
				InputArguments:  []template.Variable{{"input", "[]int32"}},
				ReturnValueType: template.QueryRow,
				OutputArguments: []template.Variable{{"output", "[]int32"}},
			}},
		},
		wantErr: false,
	}
	t.Run(tt.name, func(t *testing.T) {
		file, err := ast.ParseFile(flag.Config{
			InputGoFilePath:  tt.args.Path,
			OutputGoFilePath: "/dev/null",
			Verbose:          false,
		})
		if err != nil {
			t.Fatal("ast.ParseFile(\""+tt.args.Path+"\"): ", err)
		}
		gotDot, err := Extract(file)
		if (err != nil) != tt.wantErr {
			t.Errorf("Extract() error = %#v, wantErr %#v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(gotDot, tt.wantDot) {
			t.Errorf("Extract() gotDot = %#v, want %#v", gotDot, tt.wantDot)
		}
	})
}

func TestExtract6(t *testing.T) {
	tt := test{
		name: "6 Checking several outgoing parameters.",
		args: args{Path: "./testdata/6.go"},
		wantDot: template.DotType{
			PackageName: "testdata",
			Packages:    []string{"net", "time"},
			Functions: []template.Function{{
				Name: "SelectBots",
				// language=PostgreSQL
				SQL: `
		        with "tmp"("ip", "connect_time", "is_bot") as (values (
		            '2a02:6b8:b081:502::1:a'::inet,
		            '{"2020-04-13T15:12:15+03:00","2020-04-13T14:12:15+03:00","2020-04-13T13:12:15+03:00"}'::timestamp with time zone[],
		            true
		        )) select "ip", "connect_time", "is_bot"
		        from "tmp"
				where "is_bot";`,
				InputArguments: []template.Variable{
					{"ip", "net.IPNet"},
					{"connectTime", "[]time.Time"},
					{"isBot", "bool"},
				},
			}},
		},
		wantErr: false,
	}
	t.Run(tt.name, func(t *testing.T) {
		file, err := ast.ParseFile(flag.Config{
			InputGoFilePath:  tt.args.Path,
			OutputGoFilePath: "/dev/null",
			Verbose:          false,
		})
		if err != nil {
			t.Fatal("ast.ParseFile(\""+tt.args.Path+"\"): ", err)
		}
		gotDot, err := Extract(file)
		if (err != nil) != tt.wantErr {
			t.Errorf("Extract() error = %#v, wantErr %#v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(gotDot, tt.wantDot) {
			t.Errorf("Extract() gotDot = %#v, want %#v", gotDot, tt.wantDot)
		}
	})
}

func TestExtract7(t *testing.T) {
	tt := test{
		name: "7 Checking multiple input parameters; checking the unpacking of the structure fields.",
		args: args{Path: "./testdata/7.go"},
		wantDot: template.DotType{
			PackageName: "testdata",
			Packages:    []string{"github.com/jackc/pgtype"},
			Functions: []template.Function{{
				Name: "SelectUsers",
				// language=PostgreSQL
				SQL: `
        with "tmp" ("key", "value") as (values -- отсылка https://ru.wikipedia.org/wiki/Код_Дурова
            (1, '{"name": "Павел Дуров"}'::json),
            (2, '{"name": "Александра Владимирова", "deleted": true}'::json),
            (3, '{"name": "Вячеслав Мирилашвили", "deleted": true}'::json),
            (4, '{"name": "Лев Левиев", "deleted": true}'::json)                               
        ) select "key", "value"
        from "tmp"
        where coalesce("value"->>'deleted', false)::bool = $2::bool
        order by
            case when $1::bool then "key" end desc,
            case when not $1::bool then "key" end asc;`,
				InputArguments: []template.Variable{
					{"ascendingOrder", "bool"},
					{"deleted", "pgtype.Bool"},
				},
				OutputArguments:    []template.Variable{{"settings", "[]Setting"}},
				ReturnValueType:    template.QueryStruct,
				UnderlyingTypeName: "Setting",
				RowFieldsNames:     []string{"Key", "Value"},
			}},
		},
		wantErr: false,
	}
	t.Run(tt.name, func(t *testing.T) {
		file, err := ast.ParseFile(flag.Config{
			InputGoFilePath:  tt.args.Path,
			OutputGoFilePath: "/dev/null",
			Verbose:          false,
		})
		if err != nil {
			t.Fatal("ast.ParseFile(\""+tt.args.Path+"\"): ", err)
		}
		gotDot, err := Extract(file)
		if (err != nil) != tt.wantErr {
			t.Errorf("Extract() error = %#v, wantErr %#v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(gotDot, tt.wantDot) {
			t.Errorf("Extract() gotDot = %#v, want %#v", gotDot, tt.wantDot)
		}
	})
}
func TestExtract8(t *testing.T) {
	tt := test{
		name:    "8 Check that the composite types are not valid.",
		args:    args{Path: "./testdata/8.go"},
		wantDot: template.DotType{},
		wantErr: true,
	}
	t.Run(tt.name, func(t *testing.T) {
		file, err := ast.ParseFile(flag.Config{
			InputGoFilePath:  tt.args.Path,
			OutputGoFilePath: "/dev/null",
			Verbose:          false,
		})
		if err != nil {
			t.Fatal("ast.ParseFile(\""+tt.args.Path+"\"): ", err)
		}
		gotDot, err := Extract(file)
		if (err != nil) != tt.wantErr {
			t.Errorf("Extract() error = %#v, wantErr %#v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(gotDot, tt.wantDot) {
			t.Errorf("Extract() gotDot = %#v, want %#v", gotDot, tt.wantDot)
		}
	})
}
