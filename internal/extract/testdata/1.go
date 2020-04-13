package testdata

// входные параметры: 0 шт
// выходные параметры: ошибка: 0 аргументов, нет err в конце

// GoDao: generate
type GoDao1 struct {
	// language=PostgreSQL
	DropTestDatabase func() `
        drop database if exists "test";`
}
