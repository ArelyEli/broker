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
	customizer = g.Validator(schemas.CreateBusinessRequest{})
)

func CreateBusinessController(c *gin.Context) {
	var input schemas.CreateBusinessRequest
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": customizer.DecryptErrors(err)})
		return
	}

	business := models.Business{
		Name:       input.Name,
		Commission: input.Commission,
	}

	if err := business.CreateBusiness(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create business"})
		return
	}

	businessResponse := schemas.BusinessResponse{
		ID:         business.ID,
		Name:       business.Name,
		Commission: business.Commission,
		CreatedAt:  business.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAt:   business.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, businessResponse)
}

func UpdateBusinessController(c *gin.Context) {
	var input schemas.UpdateBusinessRequest
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": customizer.DecryptErrors(err)})
		return
	}

	var business models.Business
	businessID := c.Param("id")
	if err := business.GetBusiness(businessID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Business not found"})
		return
	}

	err := business.UpdateBusiness(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update business"})
		return
	}

	businessResponse := schemas.BusinessResponse{
		ID:         business.ID,
		Name:       business.Name,
		Commission: business.Commission,
		CreatedAt:  business.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAt:   business.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusOK, businessResponse)
}

func GetEarningsByBusinessController(c *gin.Context) {
	var business models.Business
	businessID := c.Param("id")
	if err := business.GetBusiness(businessID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Business not found"})
		return
	}

	earnings := business.GetEarningsByBusiness(businessID)

	c.JSON(http.StatusOK, gin.H{"earnings": earnings})
}
