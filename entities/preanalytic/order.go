package preanalytic

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	Id      uuid.UUID
	Created time.Time
}
