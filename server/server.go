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
	merchantGroup := v1.Group("/merchant")
	merchant := controllers.MerchantController{}
	merchantGroup.POST("/", merchant.CreateMerchantController)

	r.Run()
}
