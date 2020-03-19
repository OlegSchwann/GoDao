//
package test_select

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewGoDao(pool *pgxpool.Pool, ctx context.Context) GoDao {
	// language=PostgreSQL
	return GoDao{
		Add: func(a, b int64) (sum int64, err error) {
			sql := `
		select ($1::int8 + $2::int8)::int8 as sum;`
			err = pool.QueryRow(ctx, sql, a, b).Scan(&sum)
			return
		},
	}
}
