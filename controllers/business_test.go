package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
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

func TestMerchantController_CreateBusiness(t *testing.T) {
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

	assert.Equal(t, http.StatusOK, w.Code)

	var createdBusiness models.Business
	db.DB.First(&createdBusiness, "name = ?", "Test Business")
	assert.Equal(t, requestBody.Name, createdBusiness.Name)
	assert.Equal(t, createdBusiness.Commission, 10.0)
}

// func TestBusinessController_BadRequestCreateBusiness(t *testing.T) {
// 	db.ClearDB()

// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	requestBody := schemas.BusinessCreateRequest{
// 		Name:       "Test Business",
// 		Commission: 10.0,
// 	}
// 	requestBodyBytes, _ := json.Marshal(requestBody)

// 	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestBodyBytes))
// 	req.Header.Set("Content-Type", "application/json")
// 	c.Request = req

// 	CreateBusinessController(c)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var createdBusiness models.Business
// 	db.DB.First(&createdBusiness, "name = ?", "Test Business")
// 	assert.Equal(t, requestBody.Name, createdBusiness.Name)
// 	assert.Equal(t, createdBusiness.Commission, 10.0)
// }
