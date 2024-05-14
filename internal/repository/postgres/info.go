package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/himmel520/pgPro/internal/repository"
	"github.com/himmel520/pgPro/pkg/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// GetCommandInfo retrieves info about a command by its ID from the db.
func (r *Repository) GetCommandInfo(ctx context.Context, id string) (*model.CommandInfo, error) {
	query := `SELECT c.id, c.name, c.description, c.script,
					ci.id , ci.commands_id, ci.start_time, ci.end_time, ci.exitcode, ci.output
			 FROM commands c
			 JOIN commands_info ci ON c.id = ci.commands_id
				 WHERE c.id=$1`

	ci := &model.CommandInfo{}
	var endTime pgtype.Timestamp
	err := r.DB.QueryRow(ctx, query, id).Scan(
		&ci.Command.ID, &ci.Command.Name, &ci.Command.Description, &ci.Command.Script,
		&ci.CommandRun.ID, &ci.CommandRun.CommandID, &ci.CommandRun.StartTime, &endTime, &ci.CommandRun.ExitCode, &ci.CommandRun.Output)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return ci, repository.ErrRecordNotExist
	}

	ci.CommandRun.EndTime = endTime.Time

	return ci, err
}

// UpdateCommandInfo updates info about a command run in the db.
func (r *Repository) UpdateCommandInfo(ctx context.Context, c *model.CommandRun) error {
	args := []interface{}{}

	var params string
	if !c.EndTime.IsZero() {
		params += fmt.Sprintf("end_time=$%v,", len(args)+1)
		args = append(args, c.EndTime)
	}

	if c.ExitCode != 0 {
		params += fmt.Sprintf("exitcode=$%v,", len(args)+1)
		args = append(args, c.ExitCode)
	}

	if c.Output != "" {
		params += fmt.Sprintf("output=output || $%v", len(args)+1)
		args = append(args, c.Output)
	}

	query := fmt.Sprintf("UPDATE commands_info SET %s WHERE commands_id=$%v", strings.TrimSuffix(params, ","), len(args)+1)
	args = append(args, c.CommandID)

	_, err := r.DB.Exec(ctx, query, args...)
	return err
}
