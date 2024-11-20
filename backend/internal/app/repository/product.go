package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"

	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
)

func (r *Repository) CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	query := `
	INSERT INTO products 
    (website_alias, name, description, price, image_ids, active, tags) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)
    RETURNING id, website_alias, name, description, price, image_ids, active, tags`

	row := r.conn.QueryRowContext(
		ctx, query,
		product.WebsiteAlias,
		product.Name,
		product.Description,
		product.Price,
		pq.Array(product.ImageIds),
		product.Active,
		pq.Array(product.Tags),
	)

	created := &model.Product{}

	if err := row.Scan(&created.Id, &created.WebsiteAlias, &created.Name, &created.Description, &created.Price,
		pq.Array(&created.ImageIds), &created.Active, pq.Array(&created.Tags)); err != nil {
		return nil, err
	}

	return created, nil
}

func (r *Repository) GetProductById(ctx context.Context, id int) (*model.Product, error) {
	query := `
	SELECT id, website_alias, name, description, price, image_ids, active, tags 
	FROM products WHERE id = $1`

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

func (r *Repository) GetActiveProductsByAlias(ctx context.Context, alias string) (model.ProductList, error) {
	query := `
	SELECT id, website_alias, name, description, price, image_ids, active, tags 
	FROM products WHERE website_alias = $1 AND active
	ORDER BY id`

	rows, err := r.conn.QueryContext(ctx, query, alias)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	products := make(model.ProductList, 0)
	for rows.Next() {
		product := &model.Product{}
		if err = rows.Scan(&product.Id, &product.WebsiteAlias, &product.Name, &product.Description, &product.Price,
			pq.Array(&product.ImageIds), &product.Active, pq.Array(&product.Tags)); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, rows.Err()
}
