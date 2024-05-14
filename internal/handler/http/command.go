package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/himmel520/pgPro/internal/repository"
	"github.com/himmel520/pgPro/pkg/model"
)

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// GetCommands retrieves all commands.
// GET /api/v1/commands/
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
// GET /api/v1/commands/:id
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
// POST /api/v1/commands/
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

	c.JSON(http.StatusCreated, gin.H{"id": id})

}

// UpdateCommand updates a command by its ID.
// PUT /api/v1/commands/:id
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

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
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
