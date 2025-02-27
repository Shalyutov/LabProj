package preanalytic

import (
	"github.com/google/uuid"
	"time"
)

type Referral struct {
	Base    BaseReferral
	Tests   []ReferralTest
	Samples []ReferralSample
}

type BaseReferral struct {
	Patient       *uuid.UUID `sql:"patient_id" json:"Patient"`
	Id            uuid.UUID  `sql:"id" json:"Id" binding:"required"`
	Order         *uuid.UUID `sql:"order_id" json:"Order"`
	IssuedAt      time.Time  `sql:"issued_at" json:"IssuedAt" binding:"required"`
	DeletedAt     *time.Time `sql:"deleted_at" json:"DeletedAt"`
	SendAt        *time.Time `sql:"send_at" json:"SendAt"`
	Height        *float32   `sql:"height" json:"Height"`
	Weight        *float32   `sql:"weight" json:"Weight"`
	TickBite      *bool      `sql:"tick_bite" json:"TickBite"`
	HIVStatus     *int32     `sql:"hiv_status" json:"HIVStatus"`
	PregnancyWeek *int32     `sql:"pregnancy_week" json:"PregnancyWeek"`
}

type ReferralSample struct {
	Id         uuid.UUID `sql:"id" json:"Id" binding:"required"`
	ReferralId uuid.UUID `sql:"referral_id" json:"ReferralId" binding:"required"`
	IssuedAt   time.Time `sql:"issued_at" json:"IssuedAt" binding:"required"`
	IsValid    *bool     `sql:"is_valid" json:"IsValid"`
	Case       int32     `sql:"case_id" json:"Case" binding:"required"`
}

type ReferralTest struct {
	ReferralId uuid.UUID `sql:"referral_id" json:"ReferralId" binding:"required"`
	TestId     int32     `sql:"test_id" json:"TestId" binding:"required"`
}
