package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"labproj/entities/preanalytic"
	"labproj/internal"
	"net/http"
)

func ConfigureOrderEndpoints(router *gin.Engine, repo *internal.OrderRepo) {
	router.GET("/orders/:id", func(c *gin.Context) {
		GetOrder(c, *repo)
	})
	router.POST("/orders", func(c *gin.Context) {
		SaveOrder(c, *repo)
	})
	router.DELETE("/orders/:id", func(c *gin.Context) {
		DeleteOrder(c, *repo)
	})
}

func GetOrder(c *gin.Context, repo internal.OrderRepo) {
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
	order, err = repo.FindById(orderId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadGateway)
	}
	if order == nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, order)
}

func SaveOrder(c *gin.Context, repo internal.OrderRepo) {
	var order preanalytic.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = repo.Save(order)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, order)
}

func DeleteOrder(c *gin.Context, repo internal.OrderRepo) {
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
	err = repo.Delete(orderId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, nil)
}
