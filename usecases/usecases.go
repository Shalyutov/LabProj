package usecases

import (
	"github.com/google/uuid"
	dict "labproj/entities/dictionary"
	"labproj/entities/preanalytic"
)

type OrderUseCase interface {
	CreateOrder(order preanalytic.Order) error
	GetOrderById(uuid uuid.UUID) (*preanalytic.Order, error)
	DeleteOrder(order preanalytic.Order) error
}

type ReferralUseCase interface {
	CreateReferral(referral preanalytic.Referral) error
	AddTest(referral preanalytic.Referral, test *dict.Test) error
	DeleteTest(referral preanalytic.Referral, test *dict.Test) error
	GetReferralById(uuid uuid.UUID) (*preanalytic.Referral, error)
	SetOrder(referral preanalytic.Referral, order preanalytic.Order) error
	DeleteReferral(referral preanalytic.Referral) error
}
