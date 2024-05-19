package http

import (
	"context"

	"github.com/gin-gonic/gin"
	_ "github.com/himmel520/pgPro/docs"
	"github.com/himmel520/pgPro/pkg/model"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// Service represents the interface for managing commands.
type Service interface {
	GetCommands(ctx context.Context) ([]*model.Command, error)
	GetCommandByID(ctx context.Context, id string) (*model.Command, error)
	UpdateCommand(ctx context.Context, c *model.Command, id string) error
	DeleteCommand(ctx context.Context, id string) error

	CreateCommand(ctx context.Context, c *model.Command) (string, error)
	RunCommand(ctx context.Context, id string) error
	StopCommand(ctx context.Context, id string) error

	GetCommandInfoByID(ctx context.Context, id string) (*model.CommandInfo, error)
}

// Handler represents the HTTP handler for managing commands.
type Handler struct {
	srv Service
}

// New creates a new instance of Handler.
func New(srv Service) *Handler {
	return &Handler{
		srv: srv,
	}
}

// InitRoutes initializes the HTTP routes.
func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// TODO: id validation
	api := r.Group("/api/v1")
	{
		api.GET("/ping", h.Ping)
		commands := api.Group("/commands")
		{
			commands.GET("/", h.GetCommands)                // Get all commands
			commands.GET("/:id", h.GetCommandByID)          // Get command by ID
			commands.GET("/:id/info", h.GetCommandInfoByID) // Get command output by ID
			commands.POST("/", h.CreateCommand)             // Create a new command
			commands.POST("/:id/stop", h.StopCommand)       // Stop a command by ID
			commands.PUT("/:id", h.UpdateCommand)           // Update a command by ID
			commands.DELETE("/:id", h.DeleteCommand)        // Delete a command by ID
		}
	}

	return r
}
