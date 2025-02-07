package preanalytic

import (
	"github.com/google/uuid"
	"labproj/entities/dictionary"
	"time"
)

type Sample struct {
	Id       uuid.UUID
	Referral Referral
	IssuedAt time.Time
	IsValid  bool
	Case     dictionary.Supply
}
