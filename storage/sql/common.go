package sql

import (
	"context"
	"strconv"

	"github.com/maestre3d/coinlog/parser"
	"github.com/maestre3d/coinlog/storage"
)

// querySQLFunc represents a SQL query as callback. Used by paginateSQLFunc to hold actual queries executions.
type querySQLFunc[E any] func(ctx context.Context, limit, offset int) ([]E, error)

// paginateSQLFunc executes pagination for SQL-related queries.
//
// Steps followed:
//
// - It parses domain.PageToken and uses it as offset.
//
// - Executes the actual queries (querySQLFunc).
//
// - If items found (no errors nor empty lists), it creates a collection from E (sql model) of T (entity).
// Uses convFunc (parser.ParseFunc).
//
// - Builds and encodes next page token (domain.PageToken).
func paginateSQLFunc[E, T any](ctx context.Context, c storage.Criteria,
	convFunc parser.ParseFunc[E, T], querySQL querySQLFunc[E]) (items []T, nextPage storage.PageToken, err error) {

	pageOffset, _ := strconv.Atoi(storage.DecodePageToken(c.PageToken))
	defer func() {
		if pageOffset == -1 {
			return
		}
		pageOffset += len(items)
		nextPage = storage.NewPageToken(strconv.Itoa(pageOffset))
	}()

	outList, err := querySQL(ctx, c.Limit, pageOffset)
	if err != nil {
		return nil, nil, err
	} else if len(outList) == 0 {
		pageOffset = -1 // no pages left
	}

	items = parser.NewCollection[E, T](outList, convFunc)
	return
}
