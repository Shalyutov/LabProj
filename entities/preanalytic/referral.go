package preanalytic

import (
	"github.com/google/uuid"
	"time"
)

type Referral struct {
	Patient       *uuid.UUID
	Id            uuid.UUID
	Order         *uuid.UUID
	Tests         []int
	Samples       []ReferralSample
	IssuedAt      time.Time
	DeletedAt     *time.Time
	SendAt        *time.Time
	Height        *float64
	Weight        *float64
	TickBite      *bool
	HIVStatus     *int
	PregnancyWeek *int
}

type ReferralSample struct {
	Id       uuid.UUID `sql:"id"`
	IssuedAt time.Time `sql:"issued_at"`
	IsValid  *bool     `sql:"is_valid"`
	Case     int       `sql:"case_id"`
}
