package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response model of response
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Respond to client in predefined format
func respondJSON(c *gin.Context, data interface{}, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: data.(string) + " - " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, Response{
			Code:    http.StatusOK,
			Message: "ok",
			Data:    data,
		})
	}
}
