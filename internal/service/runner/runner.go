package runner

import (
	"errors"
	"sync"
)

var (
	// ErrCommandExists is returned when attempting to add a command that already exists in the runner.
	ErrCommandExists = errors.New("command already exists in the runner")
	// ErrCommandNotExists is returned when attempting to perform an operation on a command that does not exist in the runner.
	ErrCommandNotExists = errors.New("command does not exist in the runner")
)

// Runner represents a command runner.
type Runner struct {
	commands map[string]*commandRun // Map to store commands by ID
	mu       sync.RWMutex
}

// NewRunner creates a new instance of Runner.
func NewRunner() *Runner {
	return &Runner{
		commands: map[string]*commandRun{},
	}
}

// AddCommand adds a new command to the runner.
func (r *Runner) AddCommand(id string, script string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.commands[id]
	if ok {
		return ErrCommandExists
	}

	r.commands[id] = newCommandRun(script)
	return nil
}

// RunCommand executes a command with the given ID.
func (r *Runner) RunCommand(id string) error {
	r.mu.Lock()
	cmd, ok := r.commands[id]
	if !ok {
		return ErrCommandNotExists
	}
	r.mu.Unlock()

	// Execute the command
	return cmd.run()
}

// StopCommand stops a command with the given ID.
func (r *Runner) StopCommand(id string) error {
	r.mu.Lock()
	cmd, ok := r.commands[id]
	if !ok {
		return ErrCommandNotExists
	}
	r.mu.Unlock()

	// Stop the command
	cmd.stop()
	return nil
}

// GetOutputChan retrieves the output channel for a command with the given ID.
func (r *Runner) GetOutputChan(id string) (chan string, error) {
	r.mu.RLock()
	cmd, ok := r.commands[id]
	if !ok {
		return nil, ErrCommandNotExists
	}
	r.mu.RUnlock()

	return cmd.output, nil
}

func (r *Runner) GetExitCode(id string) (int, error) {
	r.mu.RLock()
	cmd, ok := r.commands[id]
	if !ok {
		return -1, ErrCommandNotExists
	}
	r.mu.RUnlock()

	return cmd.getStatus(), nil
}
