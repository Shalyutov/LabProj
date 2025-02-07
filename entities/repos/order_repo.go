package repos

import (
	"github.com/google/uuid"
	"labproj/entities/preanalytic"
)

type OrderRepo interface {
	Create(order preanalytic.Order) error
	FindById(id uuid.UUID) (preanalytic.Order, error)
	Delete(order preanalytic.Order) error
}
