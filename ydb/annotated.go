package ydb

import (
	"github.com/google/uuid"
	"labproj/entities/preanalytic"
	"time"
)

type Annotated interface {
	preanalytic.Order | Referral | preanalytic.Patient | Sample | ReferralTest | preanalytic.ReferralSample
}

type ReferralTest struct {
	TestId int `sql:"test_id"`
}

type Referral struct {
	Patient       *uuid.UUID `sql:"patient_id"`
	Id            uuid.UUID  `sql:"id"`
	Order         *uuid.UUID `sql:"order_id"`
	IssuedAt      time.Time  `sql:"issued_at"`
	DeletedAt     *time.Time `sql:"deleted_at"`
	SendAt        *time.Time `sql:"send_at"`
	Height        *float64   `sql:"height"`
	Weight        *float64   `sql:"weight"`
	TickBite      *bool      `sql:"tick_bite"`
	HIVStatus     *int       `sql:"hiv_status"`
	PregnancyWeek *int       `sql:"pregnancy_week"`
}

type Sample struct {
	Id       uuid.UUID `sql:"id"`
	Referral uuid.UUID `sql:"referral_id"`
	IssuedAt time.Time `sql:"issued_at"`
	IsValid  *bool     `sql:"is_valid"`
	Case     int       `sql:"case_id"`
}
