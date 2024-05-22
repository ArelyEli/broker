package controllers

import (
	"math/rand"
	"net/http"

	"arely.dev/models"
	"arely.dev/schemas"
	"github.com/gin-gonic/gin"
)

type MerchantController struct{}

func (m MerchantController) CreateMerchantController(c *gin.Context) {
	var input schemas.MerchantCreateRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Commission == 0 {
		input.Commission = rand.Float64()*99 + 1
	}

	merchant := models.Merchant{
		Name:       input.Name,
		Commission: input.Commission,
	}

	merchant.CreateMerchant()
}
