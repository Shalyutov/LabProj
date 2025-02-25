package ydb

import (
	"context"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

type Orm struct {
	DB  *ydb.Driver
	Ctx *context.Context
}

func NewYdbOrm(driver *ydb.Driver, ctx *context.Context) *Orm {
	return &Orm{driver, ctx}
}

func (y Orm) Execute(query string, parameters *table.QueryParameters) error {
	err := y.DB.Table().DoTx(
		*y.Ctx,
		func(ctx context.Context, tx table.TransactionActor) (err error) {
			res, err := tx.Execute(ctx, query, parameters)
			if err != nil {
				return err
			}
			if err = res.Err(); err != nil {
				return err
			}
			return res.Close()
		}, table.WithIdempotent(),
	)
	return err
}

func (y Orm) BulkUpsert(ydbTable string, rows types.Value) error {
	err := y.DB.Table().Do(
		*y.Ctx,
		func(ctx context.Context, s table.Session) (err error) {
			return s.BulkUpsert(ctx, y.DB.Scheme().Database()+"/"+ydbTable, rows)
		},
	)
	return err
}

func Query[T Annotated](y *Orm, q string, parameters query.ExecuteOption) ([]T, error) {
	res, err := y.DB.Query().Query(*y.Ctx, q, parameters, query.WithIdempotent())
	if err != nil {
		panic(err)
	}
	defer func() { _ = res.Close(*y.Ctx) }()

	items := make([]T, 0)
	for rs, err := range res.ResultSets(*y.Ctx) {
		if err != nil {
			return nil, err
		}
		for row, err := range rs.Rows(*y.Ctx) {
			if err != nil {
				return nil, err
			}
			var item T
			if err = row.ScanStruct(&item); err != nil {
				return nil, err
			}
			items = append(items, item)
		}
	}
	return items, nil
}
