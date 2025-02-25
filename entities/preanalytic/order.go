package preanalytic

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	Id        uuid.UUID  `sql:"id" json:"id" binding:"required"`
	CreatedAt time.Time  `sql:"created_at" json:"createdAt" binding:"required"`
	DeletedAt *time.Time `sql:"deleted_at" json:"deletedAt"`
}
