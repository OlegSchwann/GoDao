package testdata

import "github.com/jackc/pgtype"

// 3 Checking incoming native parameters; checking outgoing composite parameters.
// входные параметры: 1 шт нативный параметр
// выходные параметры: режим template.QueryRow: 1 параметр составной, кроме ошибки

// GoDao: generate
type GoDao3 struct {
	// language=PostgreSQL
	GetSettings func(id int64) (json pgtype.JSON, err error) `
        with "tmp"("k", "v") as (values
            (0::int8, '{"dark_theme": true}'::json),
            (1::int8, '{"cookies": false}'::json)
        ) select "v"
        from "tmp"
        where "k" = $1
        limit 1;`
}
