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
	cmd, ok := r.GetCmd(id)
	if !ok {
		return ErrCommandNotExists
	}

	// Execute the command
	return cmd.run()
}

func(r *Runner) DeleteCommand(id string) {
	r.mu.Lock()
	delete(r.commands, id)
	r.mu.Unlock()
}

// StopCommand stops a command with the given ID.
func (r *Runner) StopCommand(id string) error {
	cmd, ok := r.GetCmd(id)
	if !ok {
		return ErrCommandNotExists
	}

	// Stop the command
	cmd.stop()
	return nil
}

// GetOutputChan retrieves the output channel for a command with the given ID.
func (r *Runner) GetOutputChan(id string) (chan string, error) {
	cmd, ok := r.GetCmd(id)
	if !ok {
		return nil, ErrCommandNotExists
	}

	return cmd.output, nil
}

// GetExitCode retrieves the exit code of a command by its ID.
func (r *Runner) GetExitCode(id string) (int, error) {
	cmd, ok := r.GetCmd(id)
	if !ok {
		return -1, ErrCommandNotExists
	}

	return cmd.getStatus(), nil
}

// GetCmd retrieves a command by its ID and indicates whether it exists.
func (r *Runner) GetCmd(id string) (*commandRun, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	cmd, ok := r.commands[id]
	return cmd, ok
}

// IsCmdExist checks if a command exists by its ID.
func (r *Runner) IsCmdExist(id string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	_, ok := r.commands[id]
	return ok
}