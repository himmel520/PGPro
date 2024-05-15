package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type response struct {
	Message string `json:"message"`
}

type idResponse struct {
	ID string `json:"id"`
}

type errorResponse struct {
	Message string `json:"message"`
}

// newErrorResponse creates a new error response and aborts the request
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Errorf("Response Error: %s", message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
