package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/himmel520/pgPro/internal/repository"
	"github.com/himmel520/pgPro/internal/service/runner"
)

// GetCommandInfoByID retrieves information about a command by its ID.
// GET /api/v1/commands/:id/info
func (h *Handler) GetCommandInfoByID(c *gin.Context) {
	id := c.Param("id")
	cmdInfo, err := h.srv.GetCommandInfoByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotExist) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cmdInfo)
}

// StopCommand stops a command by its ID.
// POST /api/v1/commands/:id/stop
func (h *Handler) StopCommand(c *gin.Context) {
	id := c.Param("id")

	if err := h.srv.StopCommand(c.Request.Context(), id); err != nil {
		if errors.Is(err, runner.ErrCommandNotExists) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
