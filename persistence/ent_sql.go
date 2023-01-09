package persistence

import (
	"context"
	"strconv"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"github.com/maestre3d/coinlog/configuration"
	"github.com/maestre3d/coinlog/domainutil"
	"github.com/maestre3d/coinlog/ent"
	"github.com/maestre3d/coinlog/ent/migrate"
	"github.com/maestre3d/coinlog/entity"
	"github.com/maestre3d/coinlog/valueobject"
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

type querySQLFunc[E any] func(ctx context.Context, limit, offset int) ([]E, error)

func paginateSQLFunc[E, T any](ctx context.Context, c valueobject.Criteria, convFunc domainutil.ConvertFunc[E, T],
	querySQL querySQLFunc[E]) (items []T, nextPage valueobject.PageToken, err error) {
	pageOffset, _ := strconv.Atoi(valueobject.DecodePageToken(c.PageToken))
	defer func() {
		pageOffset += len(items)
		nextPage = valueobject.NewPageToken(strconv.Itoa(pageOffset))
	}()

	outList, err := querySQL(ctx, c.Limit, pageOffset)
	if err != nil {
		return nil, nil, err
	} else if len(outList) == 0 {
		return nil, nil, entity.ErrUserNotFound
	}

	items = domainutil.NewCollection[E, T](outList, convFunc)
	return
}
