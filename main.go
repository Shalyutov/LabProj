package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	yc "github.com/ydb-platform/ydb-go-yc"
	"labproj/entities"
	"labproj/handlers"
	"labproj/internal"
	"labproj/middleware"
	orm "labproj/ydb"
	"net/http"
	"time"
)

//goland:noinspection SpellCheckingInspection
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
			"version": "25C1",
			"env":     "dev",
		})
	})

	r.GET("/preanalytic", middleware.Authorize([]string{"admin", "general"}, ydbOrm), func(c *gin.Context) {
		c.JSON(http.StatusOK, template)
	})

	r.GET("/token", GetTestToken)

	handlers.ConfigureOrderEndpoints(r, &ordersRepo)
	handlers.ConfigureSamplesEndpoints(r, &samplesRepo)
	handlers.ConfigurePatientsEndpoints(r, &patientsRepo)
	handlers.ConfigureReferralsEndpoints(r, &referralsRepo, &template)

	err = r.Run(":8080")
	if err != nil {
		return
	}
}

func GetTestToken(c *gin.Context) {
	if c.GetHeader("API-KEY") != "TEST_KEY" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "AuthLIS",
		Subject:   "shalyutov",
		Audience:  []string{"LIS"},
	}
	secret := []byte("o9384u98vr8nfy93e8ur034u03h9458uy0469h56y0n9i6tpv394omd28d3y4rv9873b456b")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, tokenString)
}
