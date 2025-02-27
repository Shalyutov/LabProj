package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	yc "github.com/ydb-platform/ydb-go-yc"
	"labproj/entities"
	"labproj/handlers"
	"labproj/internal"
	orm "labproj/ydb"
	"net/http"
)

func main() {
	template := entities.TemplateMedicalDictionary()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, err := ydb.Open(ctx,
		"grpcs://ydb.serverless.yandexcloud.net:2135/ru-central1/b1gb8mvbo8og4g8184q8/etn026pjpjqev1v6fneq",
		yc.WithInternalCA(),
		yc.WithServiceAccountKeyFileCredentials("./key.json"),
	)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = db.Close(ctx)
	}()

	ydbOrm := orm.NewYdbOrm(db, &ctx)

	var ordersRepo internal.OrderRepo = orm.OrderRepo{DB: ydbOrm}
	var referralsRepo internal.ReferralRepo = orm.ReferralRepo{DB: ydbOrm}
	var patientsRepo internal.PatientRepo = orm.PatientRepo{DB: ydbOrm}
	var samplesRepo internal.SampleRepo = orm.SampleRepo{DB: ydbOrm}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"version": "0.0.1",
			"env":     "dev-alpha",
		})
	})

	r.GET("/preanalytic", func(c *gin.Context) {
		c.JSON(http.StatusOK, template)
	})

	handlers.ConfigureOrderEndpoints(r, &ordersRepo)
	handlers.ConfigureSamplesEndpoints(r, &samplesRepo)
	handlers.ConfigurePatientsEndpoints(r, &patientsRepo)
	handlers.ConfigureReferralsEndpoints(r, &referralsRepo, &template)

	err = r.Run()
	if err != nil {
		return
	}
}
