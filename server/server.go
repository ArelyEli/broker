package server

import (
	"arely.dev/controllers"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	r.GET("/health", controllers.Health)

	v1 := r.Group("/v1")

	merchantGroup := v1.Group("/merchants")
	merchantGroup.POST("/", controllers.CreateMerchantController)
	merchantGroup.PATCH("/:id", controllers.UpdateMerchantController)
	merchantGroup.GET("/:id/earnings", controllers.GetEarningsByMerchantController)

	transactionGroup := v1.Group("/transactions")
	transactionGroup.POST("/", controllers.CreateTransactionController)
	transactionGroup.GET("/earnings", controllers.GetEarningsController)

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
