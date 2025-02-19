package ydb

import (
	"github.com/google/uuid"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
	dict "labproj/entities/dictionary"
	"labproj/entities/preanalytic"
)

type ReferralRepo struct {
	DB *Orm
}

func NewReferral(referral Referral, tests []int, samples []preanalytic.ReferralSample) *preanalytic.Referral {
	return &preanalytic.Referral{
		Id:            referral.Id,
		IssuedAt:      referral.IssuedAt,
		DeletedAt:     referral.DeletedAt,
		SendAt:        referral.SendAt,
		Height:        referral.Height,
		Weight:        referral.Weight,
		TickBite:      referral.TickBite,
		HIVStatus:     referral.HIVStatus,
		PregnancyWeek: referral.PregnancyWeek,
		Patient:       referral.Patient,
		Order:         referral.Order,
		Tests:         tests,
		Samples:       samples,
	}
}

func (r ReferralRepo) Create(referral preanalytic.Referral) error {
	//TODO implement me
	panic("implement me")
}

func (r ReferralRepo) FindById(id uuid.UUID) (*preanalytic.Referral, error) {
	q := `
		DECLARE $id AS Uuid;
		SELECT
			id, issued_at, order_id, hiv_status, patient_id, 
			deleted_at, send_at, height, weight, tick_bite, pregnancy_week
		FROM
			referrals
		WHERE 
			id = $id;
	`
	params := query.WithParameters(
		ydb.ParamsBuilder().
			Param("$id").Uuid(id).
			Build(),
	)
	referrals, err := Query[Referral](r.DB, q, params)
	if err != nil {
		panic(err)
	}
	if len(referrals) == 0 {
		return nil, nil
	}

	q = `
		DECLARE $id AS Uuid;
		SELECT
			test_id
		FROM
			referral_tests
		WHERE 
			referral_id = $id;
	`
	params = query.WithParameters(
		ydb.ParamsBuilder().
			Param("$id").Uuid(referrals[0].Id).
			Build(),
	)
	referralTests, err := Query[ReferralTest](r.DB, q, params)
	if err != nil {
		panic(err)
	}
	tests := make([]int, 0)
	for _, referralTest := range referralTests {
		tests = append(tests, referralTest.TestId)
	}

	q = `
		DECLARE $id AS Uuid;
		SELECT
			id, issued_at, is_valid, case_id
		FROM
			samples
		WHERE 
			referral_id = $id;
	`
	params = query.WithParameters(
		ydb.ParamsBuilder().
			Param("$id").Uuid(referrals[0].Id).
			Build(),
	)
	referralSamples, err := Query[preanalytic.ReferralSample](r.DB, q, params)
	if err != nil {
		panic(err)
	}

	referral := NewReferral(referrals[0], tests, referralSamples)
	return referral, nil
}

func (r ReferralRepo) GetSamples(referral preanalytic.Referral) ([]preanalytic.Sample, error) {
	//TODO implement me
	panic("implement me")
}

func (r ReferralRepo) GetTests(referral preanalytic.Referral) ([]dict.Test, error) {
	//TODO implement me
	panic("implement me")
}

func (r ReferralRepo) GetPatient(referral preanalytic.Referral) (preanalytic.Patient, error) {
	//TODO implement me
	panic("implement me")
}

func (r ReferralRepo) SetPatient(referral preanalytic.Referral, patient preanalytic.Patient) error {
	//TODO implement me
	panic("implement me")
}

func (r ReferralRepo) GetOrder(referral preanalytic.Referral) (preanalytic.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (r ReferralRepo) SetOrder(referral preanalytic.Referral, order preanalytic.Order) error {
	//TODO implement me
	panic("implement me")
}

func (r ReferralRepo) AddTest(referral preanalytic.Referral, test dict.Test) error {
	//TODO implement me
	panic("implement me")
}

func (r ReferralRepo) AddSample(referral preanalytic.Referral, sample preanalytic.Sample) error {
	//TODO implement me
	panic("implement me")
}

func (r ReferralRepo) DeleteTest(referral preanalytic.Referral, test dict.Test) error {
	//TODO implement me
	panic("implement me")
}

func (r ReferralRepo) DeleteSample(referral preanalytic.Referral, sample preanalytic.Sample) error {
	//TODO implement me
	panic("implement me")
}

func (r ReferralRepo) Delete(referral preanalytic.Referral) error {
	//TODO implement me
	panic("implement me")
}
