package insert

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
	"testing"
)

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
		log.Fatal("Unable to connect to database:", err)
	}
	os.Exit(m.Run())
}

func TestNewGoDao(t *testing.T) {
	goDao := NewGoDao(pool)
	err := goDao.Init()
	if err != nil {
		t.Fatal("goDao.Init()", err)
	}
	err = goDao.Insert(101000, "Москва")
	if err != nil {
		t.Fatal("goDao.Insert(‎101000, \"Москва\")", err)
	}
	city, err := goDao.Select(101000)
	if err != nil {
		t.Fatal("goDao.Select(101000)", err)
	}
	if city != "Москва" {
		t.Fatal("expected \"Москва\", got ", city)
	}
	err = goDao.DeleteAll()
	if err != nil {
		t.Fatal("goDao.DeleteAll()", err)
	}
}
