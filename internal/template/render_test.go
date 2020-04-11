package template

import (
	"os"
	"testing"
)

func TestRender(t *testing.T) {
	tt := DotType{
		PackageName: "goDao",
		Packages:    []string{"github.com/jackc/pgtype"},
		Functions: []Function{{
			Name: "Add",
			// language=PostgreSQL
			SQL: `
-- https://en.wikipedia.org/wiki/List_of_most_popular_websites
create table if not exists "popular_websites" (
	"site" text,
	"domain" text,
	"alexa_top" int2,
	"similar_web_top" int2,
	"description" text,
	"territory" text
);`,
			ReturnValueType: Exec,
		}, {
			Name: "Sub",
			// language=PostgreSQL
			SQL: `
select ($1::float8 - $2::float8)::float8;`,
			ReturnValueType: QueryRow,
			InputArguments: []Variable{{
				Name: "a", Type: "float64",
			}, {
				Name: "b", Type: "float64",
			}},
			OutputArguments: []Variable{{
				Name: "sum",
				Type: "float64",
			}},
		}, {
			Name: "GetPopularIn",
			// language=PostgreSQL
			SQL: `
select "domain" from "popular_websites" order by $1::name limit 5;`,
			InputArguments:     []Variable{{Name: "sortColumn", Type: "pgtype.Name"}},
			ReturnValueType:    QueryPlain,
			OutputArguments:    []Variable{{Name: "domains", Type: "[]string"}},
			UnderlyingTypeName: "string",
		}, {
			Name: "FullGetPopularIn",
			// language=PostgreSQL
			SQL: `
select "site", "domain", "alexa_top", "similar_web_top", "description", "territory"
from "popular_websites"
order by $1::name
limit 5;`,
			InputArguments:     []Variable{{Name: "sortColumn", Type: "pgtype.Name"}},
			OutputArguments:    []Variable{{Name: "popularWebsites", Type: "[]PopularWebsite"}},
			ReturnValueType:    QueryStruct,
			UnderlyingTypeName: "PopularWebsite",
			RowFieldsNames:     []string{"Site", "Domain", "Alexa_top", "SimilarWebTop", "Description", "Territory"},
		}},
	}
	err := Render(tt, os.Stdout) // TODO: add test for output
	if err != nil {
		t.Fatal(err)
	}
}
