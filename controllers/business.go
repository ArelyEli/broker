package controllers

import (
	"math/rand"
	"net/http"

	"arely.dev/models"
	"arely.dev/schemas"
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator/v2"
)

var (
	g          = galidator.New()
	customizer = g.Validator(schemas.CreateBusinessRequest{})
)

func CreateBusinessController(c *gin.Context) {
	var input schemas.CreateBusinessRequest
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": customizer.DecryptErrors(err)})
		return
	}

	if input.Commission == 0 {
		input.Commission = rand.Float64()*99 + 1
	}

	business := models.Business{
		Name:       input.Name,
		Commission: input.Commission,
	}

	if err := business.CreateBusiness(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create business"})
		return
	}

	// TODO: return the business in the correct format
	c.JSON(http.StatusCreated, business)
}
