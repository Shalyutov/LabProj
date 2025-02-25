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
	Patient       *uuid.UUID `sql:"patient_id" json:"patient"`
	Id            uuid.UUID  `sql:"id" json:"id"`
	Order         *uuid.UUID `sql:"order_id" json:"order"`
	IssuedAt      time.Time  `sql:"issued_at" json:"issuedAt"`
	DeletedAt     *time.Time `sql:"deleted_at" json:"deletedAt"`
	SendAt        *time.Time `sql:"send_at" json:"sendAt"`
	Height        *float32   `sql:"height" json:"height"`
	Weight        *float32   `sql:"weight" json:"weight"`
	TickBite      *bool      `sql:"tick_bite" json:"tickBite"`
	HIVStatus     *int8      `sql:"hiv_status" json:"HIVStatus"`
	PregnancyWeek *int8      `sql:"pregnancy_week" json:"pregnancyWeek"`
}

type ReferralSample struct {
	Id       uuid.UUID `sql:"id" json:"id"`
	IssuedAt time.Time `sql:"issued_at" json:"issuedAt"`
	IsValid  *bool     `sql:"is_valid" json:"isValid"`
	Case     int       `sql:"case_id" json:"caseId"`
}

type ReferralTest struct {
	TestId int `sql:"test_id" json:"testId"`
}
