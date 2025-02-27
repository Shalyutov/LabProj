package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"labproj/entities/preanalytic"
	"labproj/internal"
	"net/http"
)

func ConfigurePatientsEndpoints(router *gin.Engine, repo *internal.PatientRepo) {
	router.GET("/patients/:id", func(c *gin.Context) {
		GetPatient(c, *repo)
	})
	router.GET("/patients", func(c *gin.Context) {
		GetPatients(c, *repo)
	})
	router.POST("/patients", func(c *gin.Context) {
		SavePatient(c, *repo)
	})
	router.DELETE("/patients/:id", func(c *gin.Context) {
		DeletePatient(c, *repo)
	})
}

func GetPatient(c *gin.Context, repo internal.PatientRepo) {
	patientIdParam := c.Param("id")
	if patientIdParam == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	patientId, err := uuid.Parse(patientIdParam)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var patient *preanalytic.Patient
	patient, err = repo.FindById(patientId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if patient == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, patient)
}

func GetPatients(c *gin.Context, repo internal.PatientRepo) {
	patients, err := repo.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if len(patients) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, patients)
}

func SavePatient(c *gin.Context, repo internal.PatientRepo) {
	var patient preanalytic.Patient
	err := c.ShouldBindJSON(&patient)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = repo.Save(patient)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, patient)
}

func DeletePatient(c *gin.Context, repo internal.PatientRepo) {
	patientIdParam := c.Param("id")
	if patientIdParam == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	patientId, err := uuid.Parse(patientIdParam)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = repo.DeleteById(patientId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, nil)
}
