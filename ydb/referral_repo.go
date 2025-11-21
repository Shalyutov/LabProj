package ydb

import (
	"labproj/entities/preanalytic"
	"time"

	"github.com/google/uuid"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

type ReferralRepo struct {
	DB *Orm
}

func NewReferral(referral preanalytic.BaseReferral, tests []preanalytic.ReferralTest, samples []preanalytic.ReferralSample) preanalytic.Referral {
	return preanalytic.Referral{
		Base:    referral,
		Tests:   tests,
		Samples: samples,
	}
}

func (r ReferralRepo) Save(referral preanalytic.Referral) error {
	err := r.SaveReferral(referral.Base)
	if err != nil {
		return err
	}
	tests := make([]int, len(referral.Tests))
	for _, value := range referral.Tests {
		tests = append(tests, int(value.TestId))
	}
	err = r.AddTests(referral.Base.Id, tests)
	if err != nil {
		return err
	}
	return nil
}

func (r ReferralRepo) SaveReferral(referral preanalytic.BaseReferral) error {
	q := `
		DECLARE $id AS Uuid;
		DECLARE $order_id AS Uuid?;
		DECLARE $patient_id AS Uuid?;
		DECLARE $issued_at AS Datetime;
		DECLARE $send_at AS Datetime?;
		DECLARE $deleted_at AS Datetime?;
		DECLARE $height AS Float?;
		DECLARE $weight AS Float?;
		DECLARE $tick_bite AS Bool?;
		DECLARE $hiv_status AS Int?;
		DECLARE $pregnancy_week AS Int?;
		DECLARE $accepted_at AS Datetime?;
		UPSERT INTO referrals ( id, order_id, patient_id, issued_at, send_at, deleted_at, 
			height, weight, tick_bite, hiv_status, pregnancy_week, accepted_at )
		VALUES ( $id, $order_id, $patient_id, $issued_at, $send_at, $deleted_at, 
			$height, $weight, $tick_bite, $hiv_status, $pregnancy_week, $accepted_at );
	`
	params := table.NewQueryParameters(
		table.ValueParam("$id", types.UuidValue(referral.Id)),
		table.ValueParam("$order_id", types.NullableUUIDTypedValue(referral.Order)),
		table.ValueParam("$patient_id", types.NullableUUIDTypedValue(referral.Patient)),
		table.ValueParam("$issued_at", types.DatetimeValueFromTime(referral.IssuedAt)),
		table.ValueParam("$send_at", types.NullableDatetimeValueFromTime(referral.SendAt)),
		table.ValueParam("$deleted_at", types.NullableDatetimeValueFromTime(referral.DeletedAt)),
		table.ValueParam("$height", types.NullableFloatValue(referral.Height)),
		table.ValueParam("$weight", types.NullableFloatValue(referral.Weight)),
		table.ValueParam("$tick_bite", types.NullableBoolValue(referral.TickBite)),
		table.ValueParam("$hiv_status", types.NullableInt32Value(referral.HIVStatus)),
		table.ValueParam("$pregnancy_week", types.NullableInt32Value(referral.PregnancyWeek)),
		table.ValueParam("$accepted_at", types.NullableDatetimeValueFromTime(referral.AcceptedAt)),
	)
	return r.DB.Execute(q, params)
}

func (r ReferralRepo) FindById(id uuid.UUID) (*preanalytic.Referral, error) {
	q := `
		DECLARE $id AS Uuid;
		SELECT
			id, issued_at, order_id, hiv_status, patient_id, 
			deleted_at, send_at, height, weight, tick_bite, pregnancy_week,
			accepted_at
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
	referrals, err := Query[preanalytic.BaseReferral](r.DB, q, params)
	if err != nil {
		panic(err)
	}
	if len(referrals) == 0 {
		return nil, nil
	}

	q = `
		DECLARE $id AS Uuid;
		SELECT
			referral_id, test_id
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
	referralTests, err := Query[preanalytic.ReferralTest](r.DB, q, params)
	if err != nil {
		panic(err)
	}

	q = `
		DECLARE $id AS Uuid;
		SELECT
			id, referral_id, issued_at, is_valid, case_id
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

	referral := NewReferral(referrals[0], referralTests, referralSamples)
	return &referral, nil
}

func (r ReferralRepo) AddTests(id uuid.UUID, tests []int) error {
	rows := make([]types.Value, 0, len(tests))
	for _, test := range tests {
		rows = append(rows, types.StructValue(
			types.StructFieldValue("referral_id", types.UuidValue(id)),
			types.StructFieldValue("test_id", types.Int32Value(int32(test))),
		))
	}
	return r.DB.BulkUpsert("referral_tests", types.ListValue(rows...))
}

func (r ReferralRepo) DeleteTests(id uuid.UUID, tests []int) error {
	q := `
		DECLARE $id AS Uuid;
		DECLARE $tests AS List<Int32>;
		DELETE FROM referral_tests 
		WHERE referral_id = $id AND test_id in $tests;
	`
	testList := make([]types.Value, 0, len(tests))
	for _, test := range tests {
		testList = append(testList, types.Int32Value(int32(test)))
	}
	params := table.NewQueryParameters(
		table.ValueParam("$id", types.UuidValue(id)),
		table.ValueParam("$tests", types.ListValue(testList...)),
	)
	return r.DB.Execute(q, params)
}

func (r ReferralRepo) Delete(id uuid.UUID) error {
	q := `
		DECLARE $id AS Uuid;
		DELETE FROM referrals 
		WHERE id = $id;
	`
	params := table.NewQueryParameters(
		table.ValueParam("$id", types.UuidValue(id)),
	)
	return r.DB.Execute(q, params)
}

func (r ReferralRepo) GetAll() ([]preanalytic.Referral, error) {
	q := `
		SELECT
			id, issued_at, order_id, hiv_status, patient_id, 
			deleted_at, send_at, height, weight, tick_bite, pregnancy_week, 
			accepted_at
		FROM
			referrals
	`
	params := query.WithParameters(
		ydb.ParamsBuilder().
			Build(),
	)
	referrals, err := Query[preanalytic.BaseReferral](r.DB, q, params)
	if err != nil {
		panic(err)
	}

	q = `
		SELECT
			referral_id, test_id
		FROM
			referral_tests
	`
	params = query.WithParameters(
		ydb.ParamsBuilder().
			Build(),
	)
	referralTests, err := Query[preanalytic.ReferralTest](r.DB, q, params)
	if err != nil {
		panic(err)
	}

	tests := make(map[uuid.UUID][]preanalytic.ReferralTest, len(referrals))
	for _, referralTest := range referralTests {
		if _, ok := tests[referralTest.ReferralId]; !ok {
			tests[referralTest.ReferralId] = make([]preanalytic.ReferralTest, 0)
		}
		tests[referralTest.ReferralId] = append(tests[referralTest.ReferralId], referralTest)
	}

	q = `
		SELECT
			id, referral_id, issued_at, is_valid, case_id
		FROM
			samples
	`
	params = query.WithParameters(
		ydb.ParamsBuilder().
			Build(),
	)
	referralSamples, err := Query[preanalytic.ReferralSample](r.DB, q, params)
	if err != nil {
		panic(err)
	}
	samples := make(map[uuid.UUID][]preanalytic.ReferralSample, len(referrals))
	for _, referralSample := range referralSamples {
		if _, ok := samples[referralSample.ReferralId]; !ok {
			samples[referralSample.ReferralId] = make([]preanalytic.ReferralSample, 0)
		}
		samples[referralSample.ReferralId] = append(samples[referralSample.ReferralId], referralSample)
	}

	allReferrals := make([]preanalytic.Referral, 0, len(referrals))
	for _, base := range referrals {
		referral := NewReferral(base, tests[base.Id], samples[base.Id])
		allReferrals = append(allReferrals, referral)
	}

	return allReferrals, nil
}

func (r ReferralRepo) SendToLab(sendAt time.Time, referrals []uuid.UUID) error {
	q := `
		DECLARE $send_at AS Datetime;
		DECLARE $referrals AS List<Uuid>;
		UPDATE referrals
		SET send_at = $send_at
		WHERE id in $referrals;
	`
	referralsList := make([]types.Value, 0, len(referrals))
	for _, referral := range referrals {
		referralsList = append(referralsList, types.UuidValue(referral))
	}
	params := table.NewQueryParameters(
		table.ValueParam("$send_at", types.DatetimeValueFromTime(sendAt)),
		table.ValueParam("$referrals", types.ListValue(referralsList...)),
	)
	return r.DB.Execute(q, params)
}
