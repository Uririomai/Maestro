package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func (r *Repository) CreateSavedProduct(ctx context.Context, tx *sqlx.Tx, product *model.Product) (int, error) {
	query := `
	INSERT INTO saved_products (website_alias, name, description, price, image_ids, active, tags)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id`

	var id int
	err := tx.GetContext(ctx, &id, query, product.WebsiteAlias, product.Name, product.Description,
		product.Price, product.ImageIds, product.Active, product.Tags)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) GetSavedProductById(ctx context.Context, id int) (*model.Product, error) {
	query := `
	SELECT id, website_alias, name, description, price, image_ids, active, tags 
	FROM saved_products WHERE id = $1`

	row := r.conn.QueryRowContext(ctx, query, id)

	product := &model.Product{}
	if err := row.Scan(
		&product.Id,
		&product.WebsiteAlias,
		&product.Name,
		&product.Description,
		&product.Price,
		pq.Array(&product.ImageIds),
		&product.Active,
		pq.Array(&product.Tags),
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrNotFound
		}
		return nil, err
	}

	return product, nil
}
