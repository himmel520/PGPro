package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/himmel520/pgPro/pkg/model"
	"github.com/sirupsen/logrus"
)

var ErrCommandIsRunning = errors.New("the command is running")

// GetCommands retrieves all commands from the repository.
func (s *Service) GetCommands(ctx context.Context) ([]*model.Command, error) {
	return s.repo.GetCommands(ctx)
}

// GetCommandByID retrieves a command by its ID from the repository.
func (s *Service) GetCommandByID(ctx context.Context, id string) (*model.Command, error) {
	return s.repo.GetCommandByID(ctx, id)
}

// UpdateCommand updates a command in the repository.
func (s *Service) UpdateCommand(ctx context.Context, c *model.Command, id string) error {
	if c.CheckEmptyStruct() {
		return model.ErrEmptyCommand
	}

	return s.repo.UpdateCommand(ctx, c, id)
}

// DeleteCommand deletes a command by its ID from the repository.
func (s *Service) DeleteCommand(ctx context.Context, id string) error {
	if s.runner.IsCmdExist(id){
		return ErrCommandIsRunning
	}
	return s.repo.DeleteCommand(ctx, id)
}

// GetCommandInfoByID retrieves info about a command by its ID from the repository.
func (s *Service) GetCommandInfoByID(ctx context.Context, id string) (*model.CommandInfo, error) {
	return s.repo.GetCommandInfo(ctx, id)
}

func (s *Service) CreateCommand(ctx context.Context, c *model.Command) (string, error) {
	id, err := s.repo.CreateCommand(ctx, c)
	if err != nil {
		return "", err
	}

	go func() {
		if err := s.runner.AddCommand(id, c.Script); err != nil {
			logrus.Error("failed to add command to the runner:", err)
			return
		}

		if err := s.RunCommand(context.Background(), id); err != nil {
			logrus.Error("failed to run command:", err)
		}

		s.runner.DeleteCommand(id)
	}()

	return id, nil
}

// RunCommand executes a command with the given ID and updates its output and status.
func (s *Service) RunCommand(ctx context.Context, id string) error {
	errChan := make(chan error, 1)
	defer close(errChan)

	go func() {
		errChan <- s.runner.RunCommand(id)
	}()

	logrus.Infof("the comand ID=%v is running", id)

	outputs, err := s.runner.GetOutputChan(id)
	if err != nil {
		return fmt.Errorf("failed to get output channel: %v", err)
	}

	for output := range outputs {
		cr := &model.CommandRun{CommandID: id, Output: output}
		if err := s.repo.UpdateCommandInfo(ctx, cr); err != nil {
			return fmt.Errorf("failed to update command info: %v", err)
		}
	}

	status, _ := s.runner.GetExitCode(id)
	cr := &model.CommandRun{CommandID: id, EndTime: time.Now(), ExitCode: status}
	if err := s.repo.UpdateCommandInfo(ctx, cr); err != nil {
		return fmt.Errorf("failed to update command info at end: %v", err)
	}

	return <-errChan
}

// StopCommand stops the execution of a command with the given ID.
func (s *Service) StopCommand(ctx context.Context, id string) error {
	if err := s.runner.StopCommand(id); err != nil {
		return fmt.Errorf("failed to stop command: %v", err)
	}
	logrus.Infof("the comand ID=%v is stopped", id)
	return nil
}
