package ydb

import (
	"github.com/google/uuid"
	"labproj/entities/preanalytic"
	"time"
)

type Annotated interface {
	preanalytic.Order | Referral | preanalytic.Patient | preanalytic.ReferralTest | preanalytic.ReferralSample
}

type Referral struct {
	Patient       *uuid.UUID `sql:"patient_id"`
	Id            uuid.UUID  `sql:"id"`
	Order         *uuid.UUID `sql:"order_id"`
	IssuedAt      time.Time  `sql:"issued_at"`
	DeletedAt     *time.Time `sql:"deleted_at"`
	SendAt        *time.Time `sql:"send_at"`
	Height        *float32   `sql:"height"`
	Weight        *float32   `sql:"weight"`
	TickBite      *bool      `sql:"tick_bite"`
	HIVStatus     *int8      `sql:"hiv_status"`
	PregnancyWeek *int8      `sql:"pregnancy_week"`
}
