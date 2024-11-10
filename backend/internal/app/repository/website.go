package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
)

func (r *Repository) CreateWebsite(ctx context.Context, alias string, adminId int) (*model.Website, error) {
	query := `INSERT INTO websites (alias, admin_id) VALUES ($1, $2)`

	website := &model.Website{}
	err := r.conn.GetContext(ctx, query, alias, adminId)
	if isSQLError(err, model.UniqueConstraintViolationCode) {
		return nil, model.ErrAliasTaken
	}
	if err != nil {
		return nil, err
	}

	return website, nil
}

func (r *Repository) AdminHaveWebsite(ctx context.Context, adminId int) (bool, error) {
	query := `SELECT COUNT(*) FROM websites WHERE admin_id=$1`

	var count int

	err := r.conn.GetContext(ctx, &count, query, adminId)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *Repository) GetWebsiteByAlias(ctx context.Context, alias string) (*model.Website, error) {
	query := `SELECT id, admin_id, alias FROM websites WHERE alias = $1`

	website := &model.Website{}
	err := r.conn.GetContext(ctx, website, query, alias)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, model.ErrWebsiteNotFound
	}
	if err != nil {
		return nil, err
	}

	return website, nil
}

func (r *Repository) GetWebsiteByAdminId(ctx context.Context, adminId int) (*model.Website, error) {
	query := `SELECT id, admin_id, alias FROM websites WHERE admin_id = $1 LIMIT 1`

	website := &model.Website{}
	err := r.conn.GetContext(ctx, website, query, adminId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, model.ErrWebsiteNotFound
	}
	if err != nil {
		return nil, err
	}

	return website, nil
}
