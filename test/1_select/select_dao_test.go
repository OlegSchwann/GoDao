package test_select

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
	"testing"
)

//go:generate true

var pool *pgxpool.Pool
func TestMain(m *testing.M) {
	var err error
	databaseUrl, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Fatal("Environment variable DATABASE_URL is required to connect to Postgres. " +
			"Example: DATABASE_URL='postgres://postgres:@localhost:5432/postgres?sslmode=disable&pool_max_conns=8';")
	}
	pool, err = pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		log.Fatal( "Unable to connect to database:", err)
	}
	os.Exit(m.Run())
}

func TestNewGoDao(t *testing.T) {
	goDao := NewGoDao(pool, context.Background())

	type args struct {
		a int64
		b int64
	}

	tests := []struct {
		name string
		args args
		want int64
	}{{
		name: "2+3=5",
		args: args{2, 3},
		want: 5,
	}, {
		name: "-3+2=-1",
		args: args{-3, 2},
		want: -1,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum, err := goDao.Add(tt.args.a, tt.args.b)
			if err != nil {
				t.Error(err)
			}
			if sum != tt.want {
				t.Fatal("expected ", tt.want, " got ", sum)
			}
		})
	}
}
