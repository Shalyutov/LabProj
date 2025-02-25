package preanalytic

import (
	"github.com/google/uuid"
	"time"
)

type Sample struct {
	Id       uuid.UUID `sql:"id" json:"id"`
	Referral uuid.UUID `sql:"referral_id" json:"referralId"`
	IssuedAt time.Time `sql:"issued_at" json:"issuedAt"`
	IsValid  *bool     `sql:"is_valid" json:"isValid"`
	Case     int32     `sql:"case_id" json:"caseId"`
}
