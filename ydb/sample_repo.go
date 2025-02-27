package ydb

import (
	"github.com/google/uuid"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
	"labproj/entities/preanalytic"
)

type SampleRepo struct {
	DB *Orm
}

func (s SampleRepo) Save(sample preanalytic.Sample) error {
	q := `
		DECLARE $id AS Uuid;
		DECLARE $referral_id AS Uuid;
		DECLARE $issued_at AS Datetime;
		DECLARE $is_valid AS Bool?;
		DECLARE $case_id AS int;
		UPSERT INTO samples ( id, referral_id, issued_at, is_valid, case_id )
		VALUES ( $id, $referral_id, $issued_at, $is_valid, $case_id );
	`
	params := table.NewQueryParameters(
		table.ValueParam("$id", types.UuidValue(sample.Id)),
		table.ValueParam("$referral_id", types.UuidValue(sample.Referral)),
		table.ValueParam("$issued_at", types.DatetimeValueFromTime(sample.IssuedAt)),
		table.ValueParam("$is_valid", types.NullableBoolValue(sample.IsValid)),
		table.ValueParam("$case_id", types.Int32Value(sample.Case)),
	)
	return s.DB.Execute(q, params)
}

func (s SampleRepo) FindById(id uuid.UUID) (*preanalytic.Sample, error) {
	q := `
		DECLARE $id AS Uuid;
		SELECT
			id, referral_id, issued_at, is_valid, case_id
		FROM
			samples
		WHERE 
			id = $id;
	`
	params := query.WithParameters(
		ydb.ParamsBuilder().
			Param("$id").Uuid(id).
			Build(),
	)
	samples, err := Query[preanalytic.Sample](s.DB, q, params)
	if err != nil {
		return nil, err
	}
	if len(samples) == 0 {
		return nil, nil
	}
	return &samples[0], err
}

func (s SampleRepo) FindAllByReferralId(id uuid.UUID) ([]preanalytic.Sample, error) {
	q := `
		DECLARE $id AS Uuid;
		SELECT
			id, referral_id, issued_at, is_valid, case_id
		FROM
			samples
		WHERE 
			referral_id = $id;
	`
	params := query.WithParameters(
		ydb.ParamsBuilder().
			Param("$id").Uuid(id).
			Build(),
	)
	samples, err := Query[preanalytic.Sample](s.DB, q, params)
	if err != nil {
		return nil, err
	}
	return samples, err
}

func (s SampleRepo) GetAll() ([]preanalytic.Sample, error) {
	q := `
		SELECT
			id, referral_id, issued_at, is_valid, case_id
		FROM
			samples
	`
	params := query.WithParameters(
		ydb.ParamsBuilder().
			Build(),
	)
	samples, err := Query[preanalytic.Sample](s.DB, q, params)
	if err != nil {
		return nil, err
	}
	return samples, err
}

func (s SampleRepo) DeleteById(id uuid.UUID) error {
	q := `
	  DECLARE $id AS Uuid;
	  DELETE FROM samples
	  WHERE id = $id;
	`
	params := table.NewQueryParameters(
		table.ValueParam("$id", types.UuidValue(id)),
	)
	return s.DB.Execute(q, params)
}
