package testdata

import "github.com/jackc/pgtype"

// входные параметры: 1 шт импортированный
// выходные параметры: режим template.QueryRow: 1 параметр нативный, кроме ошибки

// GoDao: generate
type GoDao4 struct {
	// language=PostgreSQL
	NamedDatabaseSize func(name pgtype.Name) (bytes int64, err error) `
        select pg_database_size($1::name);`
}
