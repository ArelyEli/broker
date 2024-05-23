package server

import (
	"arely.dev/controllers"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	r.GET("/health", controllers.Health)

	v1 := r.Group("/v1")

	businessGroup := v1.Group("/business")
	businessGroup.POST("/", controllers.CreateBusinessController)
	businessGroup.PATCH("/:id", controllers.UpdateBusinessController)
	businessGroup.GET("/:id/earnings", controllers.GetEarningsByBusinessController)

	transactionGroup := v1.Group("/transactions")
	transactionGroup.POST("/", controllers.CreateTransactionController)
	transactionGroup.GET("/earnings", controllers.GetEarningsController)

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
