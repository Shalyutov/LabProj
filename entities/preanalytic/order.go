package preanalytic

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	Id        uuid.UUID  `sql:"id"`
	CreatedAt time.Time  `sql:"created_at"`
	DeletedAt *time.Time `sql:"deleted_at"`
}
