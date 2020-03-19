package insert

// goDao:generate
// language=PostgreSQL
type GoDao struct {
	Init func() error `
		create table if not exists "postal_code" (
		    "code" int4 primary key check ("code" >= 100000 and "code" < 1000000), -- 6 разрядов
			"city" text not null
		);`

	Insert func(code int32, city string) error `
        insert into "postal_code"(
            "code", "city"
        ) values (
            $1::int4, $2::text
        ) on conflict ("code") do update set
            "city" = EXCLUDED."city";`

	Select func(code int32) (city string, err error) `
		select "city" from "postal_code"
		where "code" = $1::int4;`

	DeleteAll func() (err error) `
        delete from "postal_code";`
}

//go:generate go run ../../main.go
