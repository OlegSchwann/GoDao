package insert

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewGoDao(pool *pgxpool.Pool) GoDao {
	// language=PostgreSQL
	return GoDao{
		Init: func() (err error) {
			sql := `
		create table if not exists "postal_code" (
		    "code" int4 primary key check ("code" >= 100000 and "code" < 1000000), -- 6 разрядов
			"city" text not null
		);`
			_, err = pool.Exec(context.Background(), sql)
			return
		},
		Insert: func(code int32, city string) (err error) {
			sql := `
        insert into "postal_code"(
            "code", "city"
        ) values (
            $1::int4, $2::text
        ) on conflict ("code") do update set
            "city" = EXCLUDED."city";`
			_, err = pool.Exec(context.Background(), sql, code, city)
			return
		},
		Select: func(code int32) (city string, err error) {
			sql := `
		select "city" from "postal_code"
		where "code" = $1::int4;`
			err = pool.QueryRow(context.Background(), sql, code).Scan(&city)
			return
		},
		DeleteAll: func() (err error){
			sql := `
        delete from "postal_code";`
			_, err = pool.Exec(context.Background(), sql)
			return
		},
	}
}
