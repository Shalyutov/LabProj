package ydb

import (
	"labproj/entities/preanalytic"
)

type User struct {
	FullName  *string `sql:"full_name"`
	IsBlocked *bool   `sql:"is_blocked"`
}

type UserScope struct {
	Scope string `sql:"scope"`
}

type Annotated interface {
	preanalytic.Order | preanalytic.BaseReferral | preanalytic.Patient | preanalytic.ReferralTest | preanalytic.ReferralSample | preanalytic.Sample | User | UserScope
}
