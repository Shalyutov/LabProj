package preanalytic

import (
	"github.com/google/uuid"
	"time"
)

type Referral struct {
	Patient       *uuid.UUID
	Id            uuid.UUID
	Order         *uuid.UUID
	Tests         []ReferralTest
	Samples       []ReferralSample
	IssuedAt      time.Time
	DeletedAt     *time.Time
	SendAt        *time.Time
	Height        *float32
	Weight        *float32
	TickBite      *bool
	HIVStatus     *int8
	PregnancyWeek *int8
}

type ReferralSample struct {
	Id       uuid.UUID `sql:"id"`
	IssuedAt time.Time `sql:"issued_at"`
	IsValid  *bool     `sql:"is_valid"`
	Case     int       `sql:"case_id"`
}

type ReferralTest struct {
	TestId int `sql:"test_id"`
}
