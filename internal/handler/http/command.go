package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/himmel520/pgPro/internal/repository"
	"github.com/himmel520/pgPro/pkg/model"
)

// @Summary Ping endpoint
// @Description Ping endpoint for health check
// @Tags ping
// @Produce json
// @Accept  json
// @Success 200 {object} response
// @Router /ping [get]
func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, response{"pong"})
}

// GetCommands retrieves all commands.
// @Summary Get all commands
// @Description Retrieve all commands
// @Tags commands
// @Produce json
// @Accept  json
// @Success 200 {array} model.Command
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/commands [get]
func (h *Handler) GetCommands(c *gin.Context) {
	commands, err := h.srv.GetCommands(c.Request.Context())
	if err != nil {
		if errors.Is(err, repository.ErrRecordsNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, commands)
}

// GetCommandByID retrieves a command by its ID.
// @Summary Get a command by ID
// @Description Retrieve a command by its ID
// @Tags commands
// @Produce json
// @Accept  json
// @Param id path string true "Command ID"
// @Success 200 {object} model.Command
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/commands/{id} [get]
func (h *Handler) GetCommandByID(c *gin.Context) {
	id := c.Param("id")
	command, err := h.srv.GetCommandByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotExist) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, command)
}

// CreateCommand creates a new command.
// @Summary Create a new command
// @Description Create a new command
// @Tags commands
// @Accept json
// @Produce json
// @Param command body model.Command true "Command object"
// @Success 201 {object} idResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/commands [post]
func (h *Handler) CreateCommand(c *gin.Context) {
	command := &model.Command{}
	if err := c.ShouldBindJSON(command); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.srv.CreateCommand(c.Request.Context(), command)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, idResponse{id})

}

// UpdateCommand updates a command by its ID.
// @Summary Update a command by ID
// @Description Update a command by its ID
// @Tags commands
// @Accept json
// @Produce json
// @Param id path string true "Command ID"
// @Param command body model.Command true "Command object"
// @Success 200 {object} response
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v1/commands/{id} [put]
func (h *Handler) UpdateCommand(c *gin.Context) {
	id := c.Param("id")
	command := &model.Command{}
	if err := c.ShouldBindJSON(command); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.srv.UpdateCommand(c.Request.Context(), command, id)
	if err != nil {
		if errors.Is(err, model.ErrEmptyCommand) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{"ok"})
}

// DeleteCommand deletes a command by its ID.
// DELETE /api/v1/commands/:id
// func (h *Handler) DeleteCommand(c *gin.Context) {
// 	id := c.Param("id")

// 	err := h.srv.DeleteCommand(c.Request.Context(), id)
// 	if err != nil {
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "ok"})
// }
