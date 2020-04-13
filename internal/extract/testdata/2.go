package testdata

// входные параметры: 0 шт
// выходные параметры: режим template.Exec: (err error)

// GoDao: generate
type GoDao2 struct {
	// language=PostgreSQL
	DropTestDatabase func() (err error) `
        drop database if exists "test";`
}
