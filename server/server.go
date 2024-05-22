package server

import (
	"arely.dev/controllers"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	health := controllers.HealthController{}
	r.GET("/health", health.Health)

	v1 := r.Group("/v1")
	businessGroup := v1.Group("/business")
	businessGroup.POST("/", controllers.CreateBusinessController)

	r.Run()
}
