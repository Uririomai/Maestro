package repository

import (
	"context"
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
	"github.com/jmoiron/sqlx"
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
	query := `SELECT id, cart_id, product_id, count FROM cart_items WHERE cart_id = $1`

	items := make([]*model.CartItem, 0)
	err := r.conn.SelectContext(ctx, &items, query, id)
	if err != nil {
		return nil, err
	}

	return &model.Cart{Items: items}, nil
}

func (r *Repository) UpsertCartItem(ctx context.Context, cartId, productId, count int) error {
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
