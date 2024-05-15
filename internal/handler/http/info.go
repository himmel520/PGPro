package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/himmel520/pgPro/internal/repository"
	"github.com/himmel520/pgPro/internal/service/runner"
)

// GetCommandInfoByID retrieves information about a command by its ID.
// @Summary Get information about a command by its ID
// @Description Retrieve information about a command by its ID
// @Tags commands info
// @Produce json
// @Accept  json
// @Param id path string true "Command ID"
// @Success 200 {object} model.CommandInfo
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/commands/{id}/info [get]
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
// @Summary Stop a command by its ID
// @Description Stop a command by its ID
// @Tags commands info
// @Produce json
// @Accept  json
// @Param id path string true "Command ID"
// @Success 200 {object} response
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/commands/{id}/stop [post]
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

	c.JSON(http.StatusOK, response{"ok"})
}
