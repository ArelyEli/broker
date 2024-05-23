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

func TestCreateBusinessController(t *testing.T) {
	db.ClearDB()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	requestBody := schemas.CreateBusinessRequest{
		Name:       "Test Business",
		Commission: 10.0,
	}
	requestBodyBytes, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	CreateBusinessController(c)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdBusiness models.Business
	db.DB.First(&createdBusiness, "name = ?", "Test Business")
	assert.Equal(t, requestBody.Name, createdBusiness.Name)
	assert.Equal(t, requestBody.Commission, createdBusiness.Commission)
}

func TestCreateBusinessNoValidValueController(t *testing.T) {
	db.ClearDB()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	requestBody := schemas.CreateBusinessRequest{
		Name:       "Test Business",
		Commission: 1000,
	}
	requestBodyBytes, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	CreateBusinessController(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateBusinessController(t *testing.T) {
	db.ClearDB()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	business := models.Business{
		Name:       "Test Business",
		Commission: 10.0,
	}
	business.CreateBusiness()

	requestBody := schemas.UpdateBusinessRequest{
		Name:       "Updated Business",
		Commission: 20.0,
	}
	requestBodyBytes, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("PATCH", "/", bytes.NewBuffer(requestBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req
	c.Params = append(c.Params, gin.Param{Key: "id", Value: fmt.Sprintf("%d", business.ID)})

	UpdateBusinessController(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var updatedBusiness models.Business
	updatedBusiness.GetBusiness(strconv.Itoa(int(business.ID)))
	assert.Equal(t, requestBody.Name, updatedBusiness.Name)
	assert.Equal(t, requestBody.Commission, updatedBusiness.Commission)
}

func TestGetEarningsByBusinessController(t *testing.T) {
	db.ClearDB()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	business := models.Business{
		Name:       "Test Business",
		Commission: 10.0,
	}
	business.CreateBusiness()

	first_transaction := models.Transaction{
		Amount:     123.0,
		Fee:        123.0 * business.Commission / 100,
		BusinessID: business.ID,
	}
	first_transaction.CreateTransaction()

	second_transaction := models.Transaction{
		Amount:     234.0,
		Fee:        234.0 * business.Commission / 100,
		BusinessID: business.ID,
	}
	second_transaction.CreateTransaction()

	req, _ := http.NewRequest("GET", "/", nil)
	c.Request = req
	c.Params = append(c.Params, gin.Param{Key: "id", Value: fmt.Sprintf("%d", business.ID)})

	GetEarningsByBusinessController(c)

	assert.Equal(t, http.StatusOK, w.Code)

	type testResponse struct {
		Earnings float64 `json:"earnings"`
	}

	var response testResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, first_transaction.Fee+second_transaction.Fee, response.Earnings)
}
