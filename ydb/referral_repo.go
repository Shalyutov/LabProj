package ydb

import (
	"context"
	"github.com/google/uuid"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
	"labproj/entities/preanalytic"
)

type ReferralRepo struct {
	DB *Orm
}

func NewReferral(referral Referral, tests []preanalytic.ReferralTest, samples []preanalytic.ReferralSample) *preanalytic.Referral {
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

func (r ReferralRepo) Save(referral Referral) error {
	q := `
		DECLARE $id AS Uuid;
		DECLARE $order_id AS Uuid;
		DECLARE $patient_id AS Uuid;
		DECLARE $issued_at AS Datetime;
		DECLARE $send_at AS Datetime;
		DECLARE $deleted_at AS Datetime;
		DECLARE $height AS Float;
		DECLARE $weight AS Float;
		DECLARE $tick_bite AS Bool;
		DECLARE $hiv_status AS Int;
		DECLARE $pregnancy_week AS Int;
		UPSERT INTO referrals ( id, order_id, patient_id, issued_at, send_at, deleted_at, 
			height, weight, tick_bite, hiv_status, pregnancy_week )
		VALUES ( $id, $order_id, $patient_id, $issued_at, $send_at, $deleted_at, 
			$height, $weight, $tick_bite, $hiv_status, $pregnancy_week );
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
		table.ValueParam("$hiv_status", types.NullableInt8Value(referral.HIVStatus)),
		table.ValueParam("$pregnancy_week", types.NullableInt8Value(referral.PregnancyWeek)),
	)
	return r.DB.Execute(q, params)
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
	referralTests, err := Query[preanalytic.ReferralTest](r.DB, q, params)
	if err != nil {
		panic(err)
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

	referral := NewReferral(referrals[0], referralTests, referralSamples)
	return referral, nil
}

func (r ReferralRepo) AddTests(id uuid.UUID, tests []int) error {
	err := r.DB.DB.Table().Do(
		*r.DB.Ctx,
		func(ctx context.Context, s table.Session) (err error) {
			rows := make([]types.Value, 0, len(tests))
			for _, test := range tests {
				rows = append(rows, types.StructValue(
					types.StructFieldValue("referral_id", types.UuidValue(id)),
					types.StructFieldValue("test_id", types.Int32Value(int32(test))),
				))
			}
			return s.BulkUpsert(ctx, r.DB.DB.Scheme().Database()+"/referral_tests", types.ListValue(rows...))
		},
	)
	return err
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
