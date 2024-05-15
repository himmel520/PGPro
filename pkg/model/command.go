package model

import (
	"errors"
	"time"
)

var ErrEmptyCommand = errors.New("the command does not contain any data")

// Command represents the commands table in the database.
type Command struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Script      string `json:"script"`
}

// CommandRun represents the commands_run table in the database.
type CommandRun struct {
	ID        string    `json:"id"`
	CommandID string    `json:"command_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time,omitempty"`
	ExitCode  int       `json:"exitcode"`
	Output    string    `json:"output"`
}

// CommandInfo combines Command and CommandRun.
type CommandInfo struct {
	Command    Command    `json:"command"`
	CommandRun CommandRun `json:"command_run"`
}

func (c *Command) CheckEmptyStruct() bool {
	return c.Name == "" && c.Description == "" && c.Script == ""
}
