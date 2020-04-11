package native_types

import (
	"github.com/jackc/pgtype"
)

// goDao: generate
// language=PostgreSQL
type GoDao struct {
	Init func() (err error) `
		create extension if not exists pgcrypto;
        create table if not exists "equivalent_of_mongodb" (
            "key" uuid primary key default gen_random_uuid(),
            "value" json not null
        );`

	InsertOne func(json pgtype.JSON) (uuid pgtype.UUID, err error) `
        insert into "equivalent_of_mongodb"("value")
        values ($1::json)
        returning "key";`

	Find func(uuid pgtype.UUID) (json pgtype.JSON, err error) `
		select "value"
		from "equivalent_of_mongodb"
		where key = $1::uuid;`

	FindAll func(uuids []pgtype.UUID) (documents []pgtype.JSON, err error) `
        select "value"
        from "equivalent_of_mongodb"
        where key = any ($1::uuid[])`
}
