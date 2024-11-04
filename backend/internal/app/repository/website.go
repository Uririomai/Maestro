package repository

import (
	"context"
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
)

func (r *Repository) CreateWebsite(ctx context.Context, alias string, adminId int) error {
	query := `INSERT INTO websites (alias, admin_id) VALUES ($1, $2)`

	_, err := r.conn.ExecContext(ctx, query, alias, adminId)
	if isSQLError(err, model.UniqueConstraintViolationCode) {
		return model.ErrAliasTaken
	}
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) AdminHaveWebsite(ctx context.Context, adminID int) (bool, error) {
	query := `SELECT COUNT(*) FROM websites WHERE admin_id=$1`

	var count int

	err := r.conn.GetContext(ctx, &count, query, adminID)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
