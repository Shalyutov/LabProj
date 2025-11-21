package lab

import (
	"labproj/entities/dictionary"
	"labproj/entities/preanalytic"
	"time"

	"github.com/google/uuid"
)

type WorkUnit struct {
	Base      BaseWorkUnit
	Referral  preanalytic.BaseReferral
	Eqiupment dictionary.Eqiupment
	Test      dictionary.Test
}

type BaseWorkUnit struct {
	Id          uuid.UUID  `sql:"id" json:"Id" binding:"required"`
	ReferralId  uuid.UUID  `sql:"referral_id" json:"ReferralId" binding:"required"`
	EquipmentId int        `sql:"equipment_id" json:"EquipmentId" binding:"required"`
	TestId      int        `sql:"test_id" json:"TestId" binding:"required"`
	QueuedAt    *time.Time `sql:"queued_at" json:"QueuedAt"`
	ProcessedAt *time.Time `sql:"processed_at" json:"ProcessedAt"`
	UnitResult  int        `sql:"unit_result" json:"UnitResult"`
}
