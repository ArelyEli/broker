package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"arely.dev/db"
	"arely.dev/models"
	"arely.dev/schemas"
)

func TestMain(m *testing.M) {
	db.Init()

	exitVal := m.Run()

	db.ClearDB()
	os.Exit(exitVal)
}

func TestCreateMerchantController(t *testing.T) {
	db.ClearDB()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	requestBody := schemas.CreateMerchantRequest{
		Name:       "Test Merchant",
		Commission: 10.0,
	}
	requestBodyBytes, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	CreateMerchantController(c)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdMerchant models.Merchant
	db.DB.First(&createdMerchant, "name = ?", "Test Merchant")
	assert.Equal(t, requestBody.Name, createdMerchant.Name)
	assert.Equal(t, requestBody.Commission, createdMerchant.Commission)
}

func TestCreateMerchantNoValidValueController(t *testing.T) {
	db.ClearDB()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	requestBody := schemas.CreateMerchantRequest{
		Name:       "Test Merchant",
		Commission: 1000,
	}
	requestBodyBytes, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	CreateMerchantController(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateMerchantController(t *testing.T) {
	db.ClearDB()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	merchant := models.Merchant{
		Name:       "Test Merchant",
		Commission: 10.0,
	}
	merchant.CreateMerchant()

	requestBody := schemas.UpdateMerchantRequest{
		Name:       "Updated Merchant",
		Commission: 20.0,
	}
	requestBodyBytes, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("PATCH", "/", bytes.NewBuffer(requestBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req
	c.Params = append(c.Params, gin.Param{Key: "id", Value: fmt.Sprintf("%d", merchant.ID)})

	UpdateMerchantController(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var updatedMerchant models.Merchant
	updatedMerchant.GetMerchant(strconv.Itoa(int(merchant.ID)))
	assert.Equal(t, requestBody.Name, updatedMerchant.Name)
	assert.Equal(t, requestBody.Commission, updatedMerchant.Commission)
}

func TestGetEarningsByMerchantController(t *testing.T) {
	db.ClearDB()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	merchant := models.Merchant{
		Name:       "Test Merchant",
		Commission: 10.0,
	}
	merchant.CreateMerchant()

	first_transaction := models.Transaction{
		Amount:     123.0,
		Fee:        123.0 * merchant.Commission / 100,
		MerchantID: merchant.ID,
	}
	first_transaction.CreateTransaction()

	second_transaction := models.Transaction{
		Amount:     234.0,
		Fee:        234.0 * merchant.Commission / 100,
		MerchantID: merchant.ID,
	}
	second_transaction.CreateTransaction()

	req, _ := http.NewRequest("GET", "/", nil)
	c.Request = req
	c.Params = append(c.Params, gin.Param{Key: "id", Value: fmt.Sprintf("%d", merchant.ID)})

	GetEarningsByMerchantController(c)

	assert.Equal(t, http.StatusOK, w.Code)

	type testResponse struct {
		Earnings float64 `json:"earnings"`
	}

	var response testResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, first_transaction.Fee+second_transaction.Fee, response.Earnings)
}
