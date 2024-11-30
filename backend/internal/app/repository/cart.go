package repository

import (
	"context"
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func (r *Repository) CreateCartTX(ctx context.Context, tx *sqlx.Tx, customerId int) error {
	query := `INSERT INTO carts (id) VALUES ($1);`

	_, err := tx.ExecContext(ctx, query, customerId)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetCart(ctx context.Context, id int) (*model.Cart, error) {
	query := `
	SELECT ci.id, ci.cart_id, ci.count, 
	       p.id, p.website_alias, p.name, p.description, p.price, p.image_ids, p.active, p.tags
	FROM cart_items ci
	JOIN products p ON ci.product_id = p.id 
	WHERE cart_id = $1
	ORDER BY ci.id`

	items := make([]*model.CartItem, 0)

	rows, err := r.conn.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		item := &model.CartItem{}
		if err = rows.Scan(
			&item.Id, &item.CartId, &item.Count,
			&item.Product.Id, &item.WebsiteAlias, &item.Name, &item.Description, &item.Price,
			pq.Array(&item.ImageIds), &item.Active, pq.Array(&item.Tags),
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return &model.Cart{Items: items}, rows.Err()
}

func (r *Repository) UpsertCartItem(ctx context.Context, cartId, productId, count int) error {
	// TODO: добавить поле is_select (типо выбрано для заказа) и чекать при оформлении заказа
	query := `
	INSERT INTO cart_items (cart_id, product_id, count) 
	VALUES ($1, $2, $3)
	ON CONFLICT (cart_id, product_id) DO UPDATE 
	SET count = excluded.count;`

	if count == 0 {
		err := r.DeleteCartItem(ctx, cartId, productId)
		if err != nil {
			return err
		}
		return nil
	}

	_, err := r.conn.ExecContext(ctx, query, cartId, productId, count)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteCartItem(ctx context.Context, cartId, productId int) error {
	query := `DELETE FROM cart_items WHERE cart_id = $1 AND product_id = $2`

	_, err := r.conn.ExecContext(ctx, query, cartId, productId)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteCartItemTX(ctx context.Context, tx *sqlx.Tx, cartId, productId int) error {
	query := `DELETE FROM cart_items WHERE cart_id = $1 AND product_id = $2`

	_, err := tx.ExecContext(ctx, query, cartId, productId)
	if err != nil {
		return err
	}

	return nil
}
