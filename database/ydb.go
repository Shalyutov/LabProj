package database

import (
	"context"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"labproj/entities/preanalytic"
)

type YdbOrm struct {
	DB  *ydb.Driver
	Ctx *context.Context
}

type YdbAnnotated interface {
	preanalytic.Order
}

func NewYdbOrm(driver *ydb.Driver, ctx *context.Context) *YdbOrm {
	return &YdbOrm{driver, ctx}
}

func (y YdbOrm) Execute(query string, parameters *table.QueryParameters) error {
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

func Query[T YdbAnnotated](y *YdbOrm, q string, parameters query.ExecuteOption) ([]T, error) {
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
