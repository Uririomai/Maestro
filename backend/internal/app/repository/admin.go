package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
)

func (r *Repository) CreateAdmin(ctx context.Context, email, passwordHash string) (int, error) {
	query := `
	INSERT INTO admins 
    (email, password_hash) VALUES ($1, $2)
    RETURNING id`

	var id int
	err := r.conn.GetContext(ctx, &id, query, email, passwordHash)
	if isSQLError(err, model.UniqueConstraintViolationCode) {
		return 0, model.ErrEmailRegistered
	}
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) GetAdminIdByEmailPassword(ctx context.Context, email, passwordHash string) (*model.Admin, error) {
	query := `
	SELECT id, email, first_name, last_name, father_name, city, 
       telegram, image_id, email_notification, telegram_notification
	FROM admins WHERE email=$1 AND password_hash=$2`

	admin := &model.Admin{}

	err := r.conn.GetContext(ctx, admin, query, email, passwordHash)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, model.ErrWrongEmailOrPassword
	}
	if err != nil {
		return nil, err
	}

	return admin, nil
}
