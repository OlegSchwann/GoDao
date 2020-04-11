package test

// goDao: generate
type goDao struct {
	// language=PostgreSQL
	DatabaseSize func() (bytes int64) `
		select pg_database_size(current_database());`
}
