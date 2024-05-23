package controllers

import (
	"net/http"

	"arely.dev/models"
	"arely.dev/schemas"
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator/v2"
)

var (
	// TODO DRY: move this to a shared package
	g          = galidator.New()
	customizer = g.Validator(schemas.CreateMerchantRequest{})
)

func CreateMerchantController(c *gin.Context) {
	var input schemas.CreateMerchantRequest
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": customizer.DecryptErrors(err)})
		return
	}

	merchant := models.Merchant{
		Name:       input.Name,
		Commission: input.Commission,
	}

	if err := merchant.CreateMerchant(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create merchant"})
		return
	}

	merchantResponse := schemas.MerchantResponse{
		ID:         merchant.ID,
		Name:       merchant.Name,
		Commission: merchant.Commission,
		CreatedAt:  merchant.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAt:   merchant.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, merchantResponse)
}

func UpdateMerchantController(c *gin.Context) {
	var input schemas.UpdateMerchantRequest
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": customizer.DecryptErrors(err)})
		return
	}

	var merchant models.Merchant
	merchantID := c.Param("id")
	if err := merchant.GetMerchant(merchantID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Merchant not found"})
		return
	}

	err := merchant.UpdateMerchant(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update merchant"})
		return
	}

	merchantResponse := schemas.MerchantResponse{
		ID:         merchant.ID,
		Name:       merchant.Name,
		Commission: merchant.Commission,
		CreatedAt:  merchant.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAt:   merchant.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusOK, merchantResponse)
}

func GetEarningsByMerchantController(c *gin.Context) {
	var merchant models.Merchant
	merchantID := c.Param("id")
	if err := merchant.GetMerchant(merchantID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Merchant not found"})
		return
	}

	earnings := merchant.GetEarningsByMerchant(merchantID)

	c.JSON(http.StatusOK, gin.H{"earnings": earnings})
}
