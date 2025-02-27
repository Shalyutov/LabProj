package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"labproj/entities/preanalytic"
	"labproj/internal"
	"net/http"
)

func ConfigureSamplesEndpoints(router *gin.Engine, repo *internal.SampleRepo) {
	router.GET("/samples/:id", func(c *gin.Context) {
		GetSample(c, *repo)
	})
	router.GET("/samples", func(c *gin.Context) {
		GetSamples(c, *repo)
	})
	router.POST("/samples", func(c *gin.Context) {
		SaveSample(c, *repo)
	})
	router.DELETE("/samples/:id", func(c *gin.Context) {
		DeleteSample(c, *repo)
	})
}

func GetSample(c *gin.Context, repo internal.SampleRepo) {
	sampleIdParam := c.Param("id")
	if sampleIdParam == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	sampleId, err := uuid.Parse(sampleIdParam)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var sample *preanalytic.Sample
	sample, err = repo.FindById(sampleId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if sample == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, sample)
}

func GetSamples(c *gin.Context, repo internal.SampleRepo) {
	samples, err := repo.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if len(samples) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, samples)
}

func SaveSample(c *gin.Context, repo internal.SampleRepo) {
	var sample preanalytic.Sample
	err := c.ShouldBindJSON(&sample)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = repo.Save(sample)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, sample)
}

func DeleteSample(c *gin.Context, repo internal.SampleRepo) {
	sampleIdParam := c.Param("id")
	if sampleIdParam == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	sampleId, err := uuid.Parse(sampleIdParam)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = repo.DeleteById(sampleId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, nil)
}
