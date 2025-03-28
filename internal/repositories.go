package internal

import (
	"github.com/google/uuid"
	"labproj/entities/preanalytic"
	"time"
)

type OrderRepo interface {
	Save(order preanalytic.Order) error
	FindById(id uuid.UUID) (*preanalytic.Order, error)
	Delete(id uuid.UUID) error
	GetAll() ([]preanalytic.Order, error)
}

type ReferralRepo interface {
	Save(referral preanalytic.Referral) error
	FindById(id uuid.UUID) (*preanalytic.Referral, error)
	AddTests(id uuid.UUID, testId []int) error
	DeleteTests(id uuid.UUID, testId []int) error
	Delete(id uuid.UUID) error
	GetAll() ([]preanalytic.Referral, error)
	SendToLab(sendAt time.Time, referrals []uuid.UUID) error
}

type PatientRepo interface {
	Save(patient preanalytic.Patient) error
	FindById(id uuid.UUID) (*preanalytic.Patient, error)
	DeleteById(id uuid.UUID) error
	GetAll() ([]preanalytic.Patient, error)
}

type SampleRepo interface {
	Save(sample preanalytic.Sample) error
	FindById(id uuid.UUID) (*preanalytic.Sample, error)
	FindAllByReferralId(id uuid.UUID) ([]preanalytic.Sample, error)
	DeleteById(id uuid.UUID) error
	GetAll() ([]preanalytic.Sample, error)
}
