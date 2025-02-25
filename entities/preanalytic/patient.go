package preanalytic

import "github.com/google/uuid"

type Patient struct {
	Id             uuid.UUID `sql:"id" json:"id"`
	Surname        *string   `sql:"surname" json:"surname"`
	Name           *string   `sql:"name" json:"name"`
	LastName       *string   `sql:"lastname" json:"lastname"`
	Gender         *string   `sql:"gender" json:"gender"`
	Email          *string   `sql:"email" json:"email"`
	Representative *string   `sql:"representative" json:"representative"`
	Document       *uint64   `sql:"document" json:"document"`
	Phone          *uint64   `sql:"phone" json:"phone"`
	BirthYear      *int32    `sql:"birth_year" json:"birthYear"`
	BirthMonth     *int32    `sql:"birth_month" json:"birthMonth"`
	BirthDay       *int32    `sql:"birth_day" json:"birthDay"`
}
