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

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
