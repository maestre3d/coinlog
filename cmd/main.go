package main

import (
	"context"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"github.com/maestre3d/coinlog/ent"
	"github.com/maestre3d/coinlog/ent/migrate"
)

func main() {
	client, err := ent.Open(dialect.Postgres, "host=localhost port=6432 user=postgres dbname=coinlog password=root")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	err = client.Schema.Create(context.TODO(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true))
	if err != nil {
		panic(err)
	}
}
