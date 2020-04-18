package testdata

// 5 Test arrays, input and output.
// входные параметры: 1 шт массив
// выходные параметры: режим template.QueryRow: 1 параметр массив, кроме ошибки

// GoDao: generate
type GoDao5 struct {
	// language=PostgreSQL
	Delta func(input []int32) (output []int32, err error) `
        select array(
            select "tmp"."to" - "tmp"."from"
            from unnest(0 || $1::int4[], $1::int4[] || 0) as "tmp"("from", "to")
        )::int4[];`
}
