package preanalytic

import (
	"github.com/google/uuid"
	"time"
)

type Sample struct {
	Id       uuid.UUID `sql:"id" json:"Id" binding:"required"`
	Referral uuid.UUID `sql:"referral_id" json:"Referral" binding:"required"`
	IssuedAt time.Time `sql:"issued_at" json:"IssuedAt" binding:"required"`
	IsValid  *bool     `sql:"is_valid" json:"IsValid"`
	Case     int32     `sql:"case_id" json:"Case" binding:"required"`
}
