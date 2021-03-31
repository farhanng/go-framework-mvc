package router

import (
	"github.com/gin-gonic/gin"

	"github.com/farhanng/go-framework-mvc/controller"
)

// Router for routing through app
func Router() {

	route := gin.Default()

	route.GET("/welcome", controller.Get())

	route.Run(":8080")
}
