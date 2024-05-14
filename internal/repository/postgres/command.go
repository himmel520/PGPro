package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/himmel520/pgPro/internal/repository"
	"github.com/himmel520/pgPro/pkg/model"
	"github.com/jackc/pgx/v5"
)

// GetCommands retrieves all commands from the db.
func (r *Repository) GetCommands(ctx context.Context) ([]*model.Command, error) {
	rows, err := r.DB.Query(ctx, "SELECT * FROM commands")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commands := []*model.Command{}
	for rows.Next() {
		c := &model.Command{}
		if err := rows.Scan(&c.ID, &c.Name, &c.Description, &c.Script); err != nil {
			return nil, err
		}
		commands = append(commands, c)
	}

	if len(commands) == 0 {
		return nil, repository.ErrRecordsNotFound
	}

	return commands, nil
}

// GetCommandByID retrieves a command from the db by its ID.
func (r *Repository) GetCommandByID(ctx context.Context, id string) (*model.Command, error) {
	c := &model.Command{}
	err := r.DB.QueryRow(ctx, "SELECT id, name, description, script FROM commands where id=$1", id).Scan(&c.ID, &c.Name, &c.Description, &c.Script)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, repository.ErrRecordNotExist
	}

	return c, err
}

// CreateCommand creates a new command and command_info in the db.
func (r *Repository) CreateCommand(ctx context.Context, c *model.Command) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	var id string
	err = tx.QueryRow(ctx, "INSERT INTO commands (name, description, script) VALUES ($1, $2, $3) returning id",
		c.Name, c.Description, c.Script).Scan(&id)
	if err != nil {
		return "", err
	}

	_, err = tx.Exec(ctx, "INSERT INTO commands_info (commands_id, start_time) VALUES ($1, $2)", id, time.Now())
	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}

	return id, err
}

// UpdateCommand updates an existing command in the db.
func (r *Repository) UpdateCommand(ctx context.Context, c *model.Command, id string) error {
	args := []interface{}{}

	var params string
	if c.Name != "" {
		params += fmt.Sprintf("name=$%v,", len(args)+1)
		args = append(args, c.Name)
	}

	if c.Description != "" {
		params += fmt.Sprintf("description=$%v,", len(args)+1)
		args = append(args, c.Description)
	}

	if c.Script != "" {
		params += fmt.Sprintf("script=$%v", len(args)+1)
		args = append(args, c.Script)
	}

	query := fmt.Sprintf("UPDATE commands SET %s where id=$%v", strings.TrimSuffix(params, ","), len(args)+1)
	args = append(args, id)

	_, err := r.DB.Exec(ctx, query, args...)
	return err
}

// DeleteCommand deletes a command from the db by its ID.
func (r *Repository) DeleteCommand(ctx context.Context, id string) error {
	// _, err := r.DB.Exec(ctx, "DELETE FROM commands WHERE id=$1", id)
	// return err
	return nil
}
