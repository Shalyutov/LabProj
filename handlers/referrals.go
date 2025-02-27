package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"labproj/entities"
	"labproj/entities/dictionary"
	"labproj/entities/preanalytic"
	"labproj/internal"
	"net/http"
	"slices"
)

func ConfigureReferralsEndpoints(router *gin.Engine, repo *internal.ReferralRepo, dict *entities.MedicalDictionary) {
	router.GET("/referrals/:id", func(c *gin.Context) {
		GetReferral(c, *repo)
	})
	router.GET("/referrals/:id/calculate", func(c *gin.Context) {
		GetCalculation(c, *repo, *dict)
	})
	router.GET("/referrals", func(c *gin.Context) {
		GetReferrals(c, *repo)
	})
	router.POST("/referrals", func(c *gin.Context) {
		SaveReferral(c, *repo)
	})
	router.POST("/referrals/:id/tests", func(c *gin.Context) {
		SaveReferralTests(c, *repo)
	})
	router.DELETE("/referrals/:id", func(c *gin.Context) {
		DeleteReferral(c, *repo)
	})
	router.DELETE("/referrals/:id/tests", func(c *gin.Context) {
		DeleteReferralTests(c, *repo)
	})
}

func GetCalculation(c *gin.Context, repo internal.ReferralRepo, dict entities.MedicalDictionary) {
	referralIdParam := c.Param("id")
	if referralIdParam == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	referralId, err := uuid.Parse(referralIdParam)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var referral *preanalytic.Referral
	referral, err = repo.FindById(referralId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if referral == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	tests := make([]dictionary.Test, len(referral.Tests))
	for current, test := range referral.Tests {
		index := slices.IndexFunc(dict.Tests, func(t dictionary.Test) bool { return int32(t.Id) == test.TestId })
		tests[current] = dict.Tests[index]
	}

	sum := dict.Calculate(tests)
	c.JSON(http.StatusOK, gin.H{
		"sum": sum,
	})
}

func GetReferral(c *gin.Context, repo internal.ReferralRepo) {
	referralIdParam := c.Param("id")
	if referralIdParam == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	referralId, err := uuid.Parse(referralIdParam)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var referral *preanalytic.Referral
	referral, err = repo.FindById(referralId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if referral == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, referral)
}

func GetReferrals(c *gin.Context, repo internal.ReferralRepo) {
	referrals, err := repo.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if len(referrals) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, referrals)
}

func SaveReferral(c *gin.Context, repo internal.ReferralRepo) {
	var referral preanalytic.Referral
	err := c.ShouldBindJSON(&referral)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = repo.Save(referral)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, referral)
}

func SaveReferralTests(c *gin.Context, repo internal.ReferralRepo) {
	referralIdParam := c.Param("id")
	if referralIdParam == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	referralId, err := uuid.Parse(referralIdParam)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var referralTests []int
	err = c.ShouldBindJSON(&referralTests)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = repo.AddTests(referralId, referralTests)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"referralId":    referralId,
		"referralTests": referralTests,
	})
}

func DeleteReferral(c *gin.Context, repo internal.ReferralRepo) {
	referralIdParam := c.Param("id")
	if referralIdParam == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	referralId, err := uuid.Parse(referralIdParam)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = repo.Delete(referralId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"referralId": referralId,
	})
}

func DeleteReferralTests(c *gin.Context, repo internal.ReferralRepo) {
	referralIdParam := c.Param("id")
	if referralIdParam == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	referralId, err := uuid.Parse(referralIdParam)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var referralTests []int
	err = c.ShouldBindJSON(&referralTests)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = repo.DeleteTests(referralId, referralTests)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"referralId": referralId,
	})
}
