package preanalytic

import "github.com/google/uuid"

type Patient struct {
	Id         uuid.UUID
	Surname    string
	Name       string
	LastName   string
	Gender     string
	Email      string
	Document   uint64
	Phone      uint64
	BirthYear  int
	BirthMonth int
	BirthDay   int
}
