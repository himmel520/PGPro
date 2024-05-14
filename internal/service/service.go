package service

import (
	"context"

	"github.com/himmel520/pgPro/internal/service/runner"
	"github.com/himmel520/pgPro/pkg/model"
)

// Repository defines the methods that should be implemented by a repository.
type Repository interface {
	// CRUD operations for managing commands
	GetCommands(ctx context.Context) ([]*model.Command, error)
	GetCommandByID(ctx context.Context, id string) (*model.Command, error)
	CreateCommand(ctx context.Context, c *model.Command) (string, error)
	UpdateCommand(ctx context.Context, c *model.Command, id string) error
	DeleteCommand(ctx context.Context, id string) error

	// Additional methods for managing command info
	GetCommandInfo(ctx context.Context, id string) (*model.CommandInfo, error)
	UpdateCommandInfo(ctx context.Context, c *model.CommandRun) error
}

// Service represents the business logic layer of the application.
type Service struct {
	repo   Repository
	runner *runner.Runner // Runner for executing commands
}

// New creates a new instance of the service with the provided repository and runner.
func New(repo Repository) *Service {
	return &Service{repo: repo, runner: runner.NewRunner()}
}
