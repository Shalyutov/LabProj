package ydb

import (
	"labproj/entities/preanalytic"
)

type Annotated interface {
	preanalytic.Order | preanalytic.BaseReferral | preanalytic.Patient | preanalytic.ReferralTest | preanalytic.ReferralSample | preanalytic.Sample
}
