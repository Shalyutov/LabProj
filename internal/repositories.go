package internal

import (
	"github.com/google/uuid"
	"labproj/entities/preanalytic"
)

type OrderRepo interface {
	Save(order preanalytic.Order) error
	FindById(id uuid.UUID) (preanalytic.Order, error)
	Delete(id uuid.UUID) error
}

type ReferralRepo interface {
	Save(referral preanalytic.Referral) error
	FindById(id uuid.UUID) (preanalytic.Referral, error)
	AddTests(id uuid.UUID, testId int) error
	DeleteTests(id uuid.UUID, testId int) error
	Delete(id uuid.UUID) error
}
