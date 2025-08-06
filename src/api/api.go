package api

import (
	v1 "authservice/src/api/v1"
	"github.com/gin-gonic/gin"
)

func Routes() {
	setApi := gin.Default()

	api := setApi.Group("/api/v1")
	v1.ExampleRoutes(api)

	setApi.Run(":8080")
}
