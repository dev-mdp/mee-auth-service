package responses

import (
	"github.com/gin-gonic/gin"
)

type JSONResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"status"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

func Success(c *gin.Context, data interface{}, StatusCode int) {
	c.JSON(StatusCode, JSONResponse{
		StatusCode: StatusCode,
		Message:    "success",
		Data:       data,
	})
}

func Error(c *gin.Context, err interface{}, StatusCode int, message string) {
	c.JSON(StatusCode, JSONResponse{
		StatusCode: StatusCode,
		Message:    message,
		Error:      err,
	})
}
