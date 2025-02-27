package preanalytic

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	Id        uuid.UUID  `sql:"id" json:"Id" binding:"required"`
	CreatedAt time.Time  `sql:"created_at" json:"CreatedAt" binding:"required"`
	DeletedAt *time.Time `sql:"deleted_at" json:"DeletedAt"`
}
