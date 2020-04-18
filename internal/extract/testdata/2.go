package testdata

// Checking the template.Exec mode when the function signature = "func() (err error)".
// входные параметры: 0 шт
// выходные параметры: режим template.Exec: (err error)

// GoDao: generate
type GoDao2 struct {
	// language=PostgreSQL
	DropTestDatabase func() (err error) `
        drop database if exists "test";`
}
