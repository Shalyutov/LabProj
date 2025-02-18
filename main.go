package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	yc "github.com/ydb-platform/ydb-go-yc"
	"labproj/database"
	"labproj/entities"
	dict "labproj/entities/dictionary"
	"labproj/entities/preanalytic"
	"net/http"
	"strconv"
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

	ydbOrm := database.NewYdbOrm(db, &ctx)
	ordersRepo := database.NewYdbOrderRepo(ydbOrm)

	r := gin.Default()
	r.GET("/tests/:id", func(c *gin.Context) {
		GetTest(c, template.Tests)
	})
	r.GET("/orders/:id", func(c *gin.Context) {
		orderIdParam := c.Param("id")
		if orderIdParam == "" {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		orderId, err := uuid.Parse(orderIdParam)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		var order *preanalytic.Order
		order, err = ordersRepo.FindById(orderId)
		if err != nil {
			c.AbortWithStatus(http.StatusBadGateway)
		}
		if order == nil {
			c.AbortWithStatus(http.StatusNotFound)
		}
		c.JSON(200, order)
	})
	err = r.Run()
	if err != nil {
		return
	}
}

func GetTest(c *gin.Context, tests []dict.Test) {
	index, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if index >= int64(len(tests)) || index < 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	test := tests[index]
	c.JSON(200, test)
}
