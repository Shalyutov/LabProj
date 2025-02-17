package internal

import (
	"github.com/google/uuid"
	"labproj/entities/dictionary"
	"labproj/entities/preanalytic"
)

type OrderRepo interface {
	Create(order preanalytic.Order) error
	FindById(id uuid.UUID) (preanalytic.Order, error)
	Delete(order preanalytic.Order) error
}

type ReferralRepo interface {
	Create(referral preanalytic.Referral) error
	FindById(id uuid.UUID) (preanalytic.Referral, error)
	GetSamples(referral preanalytic.Referral) ([]preanalytic.Sample, error)
	GetTests(referral preanalytic.Referral) ([]dictionary.Test, error)
	GetPatient(referral preanalytic.Referral) (preanalytic.Patient, error)
	SetPatient(referral preanalytic.Referral, patient preanalytic.Patient) error
	GetOrder(referral preanalytic.Referral) (preanalytic.Order, error)
	SetOrder(referral preanalytic.Referral, order preanalytic.Order) error
	AddTest(referral preanalytic.Referral, test dictionary.Test) error
	AddSample(referral preanalytic.Referral, sample preanalytic.Sample) error
	DeleteTest(referral preanalytic.Referral, test dictionary.Test) error
	DeleteSample(referral preanalytic.Referral, sample preanalytic.Sample) error
	Delete(referral preanalytic.Referral) error
}
