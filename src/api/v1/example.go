package v1

import (
	"authservice/src/responses"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func ExampleRoutes(group *gin.RouterGroup) {
	example := group.Group("/example")
	example.GET("/generate_uuid", func(c *gin.Context) {
		id := uuid.New()

		responses.Success(c, id, http.StatusOK)
	})
}
