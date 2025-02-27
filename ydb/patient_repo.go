package ydb

import (
	"github.com/google/uuid"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
	"labproj/entities/preanalytic"
)

type PatientRepo struct {
	DB *Orm
}

func (p PatientRepo) Save(patient preanalytic.Patient) error {
	q := `
		DECLARE $id AS Uuid;
		DECLARE $surname AS Utf8?;
		DECLARE $name AS Utf8?;
		DECLARE $lastname AS Utf8?;
		DECLARE $gender AS Utf8?;
		DECLARE $email AS Utf8?;
		DECLARE $representative AS Utf8?;
		DECLARE $document AS Uint64?;
		DECLARE $phone AS Uint64?;
		DECLARE $birth_year AS Int?;
		DECLARE $birth_month AS Int?;
		DECLARE $birth_day AS Int?;
		UPSERT INTO patients ( id, surname, name, lastname, gender, email, 
			representative, document, phone, birth_year, birth_month, birth_day )
		VALUES ( $id, $surname, $name, $lastname, $gender, $email, 
			$representative, $document, $phone, $birth_year, $birth_month, $birth_day );
	`
	params := table.NewQueryParameters(
		table.ValueParam("$id", types.UuidValue(patient.Id)),
		table.ValueParam("$surname", types.NullableUTF8Value(patient.Surname)),
		table.ValueParam("$name", types.NullableUTF8Value(patient.Name)),
		table.ValueParam("$lastname", types.NullableUTF8Value(patient.LastName)),
		table.ValueParam("$gender", types.NullableUTF8Value(patient.Gender)),
		table.ValueParam("$email", types.NullableUTF8Value(patient.Email)),
		table.ValueParam("$representative", types.NullableUTF8Value(patient.Representative)),
		table.ValueParam("$document", types.NullableUint64Value(patient.Document)),
		table.ValueParam("$phone", types.NullableUint64Value(patient.Phone)),
		table.ValueParam("$birth_year", types.NullableInt32Value(patient.BirthYear)),
		table.ValueParam("$birth_month", types.NullableInt32Value(patient.BirthMonth)),
		table.ValueParam("$birth_day", types.NullableInt32Value(patient.BirthDay)),
	)
	return p.DB.Execute(q, params)
}

func (p PatientRepo) FindById(id uuid.UUID) (*preanalytic.Patient, error) {
	q := `
		DECLARE $id AS Uuid;
		SELECT
			id, surname, name, lastname, gender, email, 
			representative, document, phone, birth_year, birth_month, birth_day
		FROM
			patients
		WHERE 
			id = $id;
	`
	params := query.WithParameters(
		ydb.ParamsBuilder().
			Param("$id").Uuid(id).
			Build(),
	)
	patients, err := Query[preanalytic.Patient](p.DB, q, params)
	if err != nil {
		return nil, err
	}
	if len(patients) == 0 {
		return nil, nil
	}
	return &patients[0], err
}

func (p PatientRepo) GetAll() ([]preanalytic.Patient, error) {
	q := `
		SELECT
			id, surname, name, lastname, gender, email, 
			representative, document, phone, birth_year, birth_month, birth_day
		FROM
			patients
	`
	params := query.WithParameters(
		ydb.ParamsBuilder().
			Build(),
	)
	patients, err := Query[preanalytic.Patient](p.DB, q, params)
	if err != nil {
		return nil, err
	}
	return patients, err
}

func (p PatientRepo) DeleteById(id uuid.UUID) error {
	q := `
	  DECLARE $id AS Uuid;
	  DELETE FROM patients
	  WHERE id = $id;
	`
	params := table.NewQueryParameters(
		table.ValueParam("$id", types.UuidValue(id)),
	)
	return p.DB.Execute(q, params)
}
