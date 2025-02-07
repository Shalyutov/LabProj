package database

import (
	"github.com/google/uuid"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"labproj/entities/preanalytic"
)

type YdbOrderRepo struct {
	DB *ydb.Driver
}

func NewYdbOrderRepo(driver *ydb.Driver) *YdbOrderRepo {
	return &YdbOrderRepo{driver}
}

func (y YdbOrderRepo) Create(order preanalytic.Order) error {
	//TODO implement me
	panic("implement me")
}

func (y YdbOrderRepo) FindById(id uuid.UUID) (preanalytic.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (y YdbOrderRepo) Delete(order preanalytic.Order) error {
	//TODO implement me
	panic("implement me")
}
