package database

import (
	"context"
	"github.com/google/uuid"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/result"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/result/named"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
	"labproj/entities/preanalytic"
)

type YdbOrderRepo struct {
	DB  *ydb.Driver
	Ctx *context.Context
}

func NewYdbOrderRepo(driver *ydb.Driver, ctx *context.Context) *YdbOrderRepo {
	return &YdbOrderRepo{driver, ctx}
}

func (y YdbOrderRepo) Create(order preanalytic.Order) error {
	query := `
	  DECLARE $id AS Uuid;
	  DECLARE $created AS Datetime;
	  UPSERT INTO orders ( id, created_at )
	  VALUES ( $id, $created );
	`
	params := table.NewQueryParameters(
		table.ValueParam("$id", types.UuidValue(order.Id)),
		table.ValueParam("$created", types.DatetimeValueFromTime(order.CreatedAt)),
	)
	return y.Execute(query, params)
}

func (y YdbOrderRepo) FindById(id uuid.UUID) (*preanalytic.Order, error) {
	query := `
		DECLARE $id AS Uuid;
		SELECT
			id, created_at, deleted_at
		FROM
			orders
		WHERE 
			id = $id;
	`
	params := table.NewQueryParameters(
		table.ValueParam("$id", types.UuidValue(id)),
	)
	orders, err := y.Query(query, params)
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, nil
	}
	return &orders[0], err
}

func (y YdbOrderRepo) Delete(order preanalytic.Order) error {
	//TODO implement me
	panic("implement me")
}

func (y YdbOrderRepo) Execute(query string, parameters *table.QueryParameters) error {
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

func (y YdbOrderRepo) Query(query string, parameters *table.QueryParameters) ([]preanalytic.Order, error) {
	readTx := table.TxControl(
		table.BeginTx(
			table.WithOnlineReadOnly(),
		),
		table.CommitTx(),
	)

	orders := make([]preanalytic.Order, 0)
	err := y.DB.Table().Do(*y.Ctx,
		func(ctx context.Context, s table.Session) (err error) {
			var res result.Result
			_, res, err = s.Execute(ctx, readTx, query, parameters)
			if err != nil {
				return err
			}
			defer func(res result.Result) {
				_ = res.Close()
			}(res)

			for res.NextResultSet(ctx) {
				for res.NextRow() {
					order := preanalytic.Order{}
					err = res.ScanNamed(
						named.Required("id", &order.Id),
						named.Required("created_at", &order.CreatedAt),
						named.Optional("deleted_at", &order.DeletedAt),
					)
					if err != nil {
						return err
					}
					orders = append(orders, order)
				}
			}
			return res.Err()
		},
	)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
