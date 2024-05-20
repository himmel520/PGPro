package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// validateID is a middleware that validates the format of the UUID in the request parameter
func(h *Handler) validateID() gin.HandlerFunc {
	return func (c *gin.Context)  {
		id := c.Param("id")
		if err:= uuid.Validate(id); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		c.Next()
	}
}