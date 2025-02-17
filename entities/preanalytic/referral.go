package preanalytic

import (
	"github.com/google/uuid"
	dict "labproj/entities/dictionary"
	"time"
)

type Referral struct {
	Patient       *Patient
	Id            uuid.UUID
	Order         *Order
	Tests         []dict.Test
	Samples       []Sample
	IssuedAt      time.Time
	SendAt        time.Time
	Height        int
	Weight        int
	TickBite      bool
	HIVStatus     int
	PregnancyWeek int
}
