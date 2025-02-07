package preanalytic

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	Id        uuid.UUID
	CreatedAt time.Time
	DeletedAt time.Time
}
