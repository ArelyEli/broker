package server

import (
	"arely.dev/controllers"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	health := controllers.HealthController{}

	r.GET("/health", health.Health)

	r.Run()
}
