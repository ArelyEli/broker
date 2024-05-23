package controllers

import (
	"net/http"
	"strconv"

	"arely.dev/models"
	"arely.dev/schemas"
	"github.com/gin-gonic/gin"
)

var (
	customizerCreateTransaction = g.Validator(schemas.CreateTransactionRequest{})
)

func CreateTransactionController(c *gin.Context) {
	var request schemas.CreateTransactionRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": customizerCreateTransaction.DecryptErrors(err)})
		return
	}

	merchant := models.Merchant{}
	if err := merchant.GetMerchant(strconv.Itoa(int(request.MerchantID))); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Merchant not found"})
		return
	}

	transaction := models.Transaction{
		Amount:     request.Amount,
		MerchantID: request.MerchantID,
		Fee:        request.Amount * merchant.Commission / 100,
	}

	if err := transaction.CreateTransaction(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}
	transactionResponse := schemas.TransactionResponse{
		ID:         transaction.ID,
		MerchantID: strconv.Itoa(int(transaction.MerchantID)),
		Amount:     transaction.Amount,
		Commision:  merchant.Commission,
		Fee:        transaction.Fee,
		CreatedAt:  transaction.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, transactionResponse)
}

func GetEarningsController(c *gin.Context) {
	transaction := models.Transaction{}
	earnings := transaction.GetEarnings()
	c.JSON(http.StatusOK, gin.H{"earnings": earnings})
}
