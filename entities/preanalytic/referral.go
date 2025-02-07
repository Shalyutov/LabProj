package preanalytic

import (
	"github.com/google/uuid"
	dict "labproj/entities/dictionary"
)

type Referral struct {
	Patient       *Patient
	Id            uuid.UUID
	Order         *Order
	Tests         []dict.Test
	Samples       []Sample
	Height        int
	Weight        int
	TickBite      bool
	HIVStatus     int
	PregnancyWeek int
}
