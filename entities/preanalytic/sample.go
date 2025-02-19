package preanalytic

import (
	"github.com/google/uuid"
	"labproj/entities/dictionary"
	"time"
)

type Sample struct {
	Id       uuid.UUID         `sql:"id"`
	Referral uuid.UUID         `sql:"referral_id"`
	IssuedAt time.Time         `sql:"issued_at"`
	IsValid  *bool             `sql:"is_valid"`
	Case     dictionary.Supply `sql:"case_id"`
}
