package testdata

import "github.com/jackc/pgtype"

// 7 Checking multiple input parameters; checking the unpacking of the structure fields.
// входные параметры: 5 шт нативные и импортированные параметры
// выходные параметры: режим template.QueryStruct: 2 параметра: структура с неизвестными полями и ошибка. изменение номеров и распаковка * не поддерживается пока.

type Setting struct {
	Key   int64
	Value pgtype.JSON
}

// GoDao: generate
type GoDao7 struct {
	// language=PostgreSQL
	SelectUsers func(ascendingOrder bool, deleted pgtype.Bool) (settings []Setting, err error) `
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
            case when not $1::bool then "key" end asc;`
}
