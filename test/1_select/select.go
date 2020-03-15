// noinspection GoStructTag
package test_select

// goDao: goDao
// language=PostgreSQL
type GoDao struct {
	Add func(a, b int64) (int64, error) `
		select ($1::int8 + $2::int8)::int8 as sum;`
}
