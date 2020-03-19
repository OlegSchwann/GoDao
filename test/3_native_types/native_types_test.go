package native_types

import (
	"context"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
	"reflect"
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
	dao := NewGoDao(pool)
	err := dao.Init()
	if err != nil {
		t.Fatal(err)
	}
	json1 := []byte(`{"answer": 42}`)
	uuid1, err := dao.InsertOne(pgtype.JSON{Bytes: json1, Status: pgtype.Present})
	if err != nil {
		t.Fatal(err)
	}
	json, err := dao.Find(uuid1)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(json.Bytes, json1) {
		t.Fatal("expected `{\"answer\": 42}`, got ", json.Bytes)
	}
	json2 := []byte(`{"answer": "gladiolus"}`)
	uuid2, err := dao.InsertOne(pgtype.JSON{Bytes: json2, Status: pgtype.Present})
	if err != nil {
		t.Fatal(err)
	}
	documents, err := dao.FindAll([]pgtype.UUID{uuid1, uuid2})
	if err != nil {
		t.Fatal(err)
	}
	for _, doc := range documents {
		if !reflect.DeepEqual(doc.Bytes, json1) && !reflect.DeepEqual(doc.Bytes, json2) {
			t.Fatal("expected array of ", json1, json2, " got ", documents)
		}
	}
}
