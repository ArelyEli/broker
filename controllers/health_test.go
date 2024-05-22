package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHealthController_Health(t *testing.T) {
	healthController := HealthController{}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	healthController.Health(c)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	expectedBody := `{"message":"OK"}`
	if w.Body.String() != expectedBody {
		t.Errorf("Incorrect response body. Expected %s, got %s", expectedBody, w.Body.String())
	}
}
