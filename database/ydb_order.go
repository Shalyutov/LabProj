package database

import (
	"github.com/google/uuid"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
	"labproj/entities/preanalytic"
)

type YdbOrderRepo struct {
	DB *YdbOrm
}

func NewYdbOrderRepo(orm *YdbOrm) *YdbOrderRepo {
	return &YdbOrderRepo{orm}
}

func (y YdbOrderRepo) Create(order preanalytic.Order) error {
	q := `
	  DECLARE $id AS Uuid;
	  DECLARE $created AS Datetime;
	  UPSERT INTO orders ( id, created_at )
	  VALUES ( $id, $created );
	`
	params := table.NewQueryParameters(
		table.ValueParam("$id", types.UuidValue(order.Id)),
		table.ValueParam("$created", types.DatetimeValueFromTime(order.CreatedAt)),
	)
	return y.DB.Execute(q, params)
}

func (y YdbOrderRepo) FindById(id uuid.UUID) (*preanalytic.Order, error) {
	q := `
		DECLARE $id AS Uuid;
		SELECT
			id, created_at, deleted_at
		FROM
			orders
		WHERE 
			id = $id;
	`
	params := query.WithParameters(
		ydb.ParamsBuilder().
			Param("$id").Uuid(id).
			Build(),
	)
	orders, err := Query[preanalytic.Order](y.DB, q, params)
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, nil
	}
	return &orders[0], err
}

func (y YdbOrderRepo) Delete(order preanalytic.Order) error {
	q := `
	  DECLARE $id AS Uuid;
	  DELETE FROM orders
	  WHERE id = $id;
	`
	params := table.NewQueryParameters(
		table.ValueParam("$id", types.UuidValue(order.Id)),
	)
	return y.DB.Execute(q, params)
}
