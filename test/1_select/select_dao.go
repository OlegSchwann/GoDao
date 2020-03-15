package test_select

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewGoDao(pool *pgxpool.Pool, ctx context.Context) GoDao {
	return GoDao{
		Add: func(a, b int64) (sum int64, err error) {
			err = pool.QueryRow(ctx, `select ($1::int8 + $2::int8)::int8 as sum;`, a, b).Scan(&sum)
			return
		},
	}
}
