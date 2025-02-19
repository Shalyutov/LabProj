package preanalytic

import "github.com/google/uuid"

type Patient struct {
	Id             uuid.UUID `sql:"id"`
	Surname        *string   `sql:"surname"`
	Name           *string   `sql:"name"`
	LastName       *string   `sql:"last_name"`
	Gender         *string   `sql:"gender"`
	Email          *string   `sql:"email"`
	Representative *string   `sql:"representative"`
	Document       *uint64   `sql:"document"`
	Phone          *uint64   `sql:"phone"`
	BirthYear      *int      `sql:"birth_year"`
	BirthMonth     *int      `sql:"birth_month"`
	BirthDay       *int      `sql:"birth_day"`
}
