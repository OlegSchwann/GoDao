package testdata

// The absence of an "err error" as the last parameter is incorrect.
// входные параметры: 0 шт
// выходные параметры: ошибка: 0 аргументов, нет err в конце

// GoDao: generate
type GoDao1 struct {
	// language=PostgreSQL
	DropTestDatabase func() `
        drop database if exists "test";`
}
