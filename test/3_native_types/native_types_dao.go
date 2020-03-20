package native_types

import (
	"context"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewGoDao(pool *pgxpool.Pool) GoDao {
	return GoDao{
		Init: func() (err error) {
			// language=PostgreSQL
			sql := `
		create extension if not exists pgcrypto;
        create table if not exists "equivalent_of_mongodb" (
            "key" uuid primary key default gen_random_uuid(),
            "value" json not null
        );`
			_, err = pool.Exec(context.Background(), sql)
			return
		},
		InsertOne: func(json pgtype.JSON) (uuid pgtype.UUID, err error) {
			// language=PostgreSQL
			sql := `
        insert into "equivalent_of_mongodb"("value") values ($1::json) returning "key";`
			err = pool.QueryRow(context.Background(), sql, json).Scan(&uuid)
			return
		},
		Find: func(uuid pgtype.UUID) (json pgtype.JSON, err error) {
			// language=PostgreSQL
			sql := `
		select "value" from "equivalent_of_mongodb" where key = $1::uuid;`
			err = pool.QueryRow(context.Background(), sql, uuid).Scan(&json)
			return
		},
		FindAll: func(uuids []pgtype.UUID) (documents []pgtype.JSON, err error) {
			// language=PostgreSQL
			sql := `
        select "value"
        from "equivalent_of_mongodb"
        where key = any ($1::uuid[])`
			rows, err := pool.Query(context.Background(), sql, uuids)
			if err != nil {
				return
			}
			defer rows.Close()
			for rows.Next() {
				var tmp pgtype.JSON
				err = rows.Scan(&tmp)
				if err != nil {
					return
				}
				documents = append(documents, tmp)
			}
			return
		},
	}
}
