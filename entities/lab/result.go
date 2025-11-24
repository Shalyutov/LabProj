package lab

import (
	"labproj/entities/dictionary"
	"time"

	"github.com/google/uuid"
)

type Result struct {
	Base      BaseResult
	Indicator dictionary.Indicator
}

type BaseResult struct {
	UnitId       uuid.UUID
	IndicatorId  int
	StringValue  string
	BinaryValue  bool
	IntegerValue float32
	IssuedAt     time.Time
	ConfirmedAt  time.Time
	IsValid      bool
}
