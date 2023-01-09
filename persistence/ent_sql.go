package persistence

import (
	"context"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"github.com/maestre3d/coinlog/configuration"
	"github.com/maestre3d/coinlog/ent"
	"github.com/maestre3d/coinlog/ent/migrate"
)

func NewEntClient(cfg configuration.DatabaseSQL) (*ent.Client, func(), error) {
	c, err := ent.Open(dialect.Postgres, cfg.ConnectionString)
	if err != nil {
		return nil, nil, err
	}
	err = c.Schema.Create(context.TODO(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		return nil, nil, err
	}
	return c, func() {
		_ = c.Close()
	}, nil
}
