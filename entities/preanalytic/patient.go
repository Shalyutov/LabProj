package preanalytic

import (
	"time"

	"github.com/google/uuid"
)

type Patient struct {
	Id             uuid.UUID  `sql:"id" json:"Id" binding:"required"`
	Surname        *string    `sql:"surname" json:"Surname"`
	Name           *string    `sql:"name" json:"Name"`
	LastName       *string    `sql:"lastname" json:"LastName"`
	Gender         *string    `sql:"gender" json:"Gender"`
	Email          *string    `sql:"email" json:"Email"`
	Representative *string    `sql:"representative" json:"Representative"`
	Document       *uint64    `sql:"document" json:"Document"`
	Phone          *uint64    `sql:"phone" json:"Phone"`
	BirthDate      *time.Time `sql:"birth_date" json:"BirthDate"`
}
