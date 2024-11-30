package repository

import (
	"context"
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func (r *Repository) CreateOrder(ctx context.Context, customerId int, comment string) (int, error) {
	tx, err := r.conn.Beginx()
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	// Получение корзины
	cart, err := r.GetCart(ctx, customerId)
	if err != nil {
		return 0, err
	}
	if len(cart.Items) == 0 {
		return 0, model.ErrEmptyOrder
	}

	// Подсчет общей суммы
	totalSum := 0
	// Запоминаем товары, чтобы позже сохранить из в отдельную таблицу
	products := make([]*model.Product, 0, len(cart.Items))
	// Запоминаем кол-во товаров по айди товара
	counts := make(map[int]int, len(cart.Items))
	for _, item := range cart.Items {
		totalSum += item.Count * item.Price
		products = append(products, &item.Product)
		counts[item.Product.Id] = item.Count

		// Удаление товаров из корзины
		err = r.DeleteCartItemTX(ctx, tx, item.CartId, item.Product.Id)
		if err != nil {
			return 0, err
		}
	}

	query := `
	INSERT INTO orders (customer_id, total_sum, date_time, status, comment) 
	VALUES ($1, $2, CURRENT_TIMESTAMP, $3, $4)
	RETURNING id`

	var orderId int
	err = tx.GetContext(ctx, &orderId, query, customerId, totalSum, model.OrderStatusWaitingPayment, comment)
	if err != nil {
		return 0, err
	}

	// Создание сохраненных товаров и элементов заказа
	for _, product := range products {
		id, err := r.CreateSavedProduct(ctx, tx, product)
		if err != nil {
			return 0, err
		}
		err = r.CreateOrderItem(ctx, tx, orderId, id, counts[product.Id])
		if err != nil {
			return 0, err
		}
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return orderId, nil
}

func (r *Repository) CreateOrderItem(ctx context.Context, tx *sqlx.Tx, orderId, productId, count int) error {
	query := `INSERT INTO order_items (order_id, saved_product_id, count) VALUES ($1, $2, $3)`

	_, err := tx.ExecContext(ctx, query, orderId, productId, count)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetOrderIdsByCustomerId(ctx context.Context, customerId int) ([]int, error) {
	query := `SELECT id FROM orders WHERE customer_id = $1`

	ids := make([]int, 0)

	err := r.conn.SelectContext(ctx, &ids, query, customerId)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (r *Repository) GetOrderById(ctx context.Context, orderId int) (*model.Order, error) {
	queryOrderInfo := `SELECT id, customer_id, total_sum, date_time, status, comment FROM orders WHERE id = $1`

	queryOrderItems := `
	SELECT oi.id, oi.order_id, oi.count, 
	       sp.id, sp.website_alias, sp.name, sp.description, sp.price, sp.image_ids, sp.active, sp.tags
	FROM order_items oi
	JOIN saved_products sp ON oi.saved_product_id = sp.id 
	WHERE order_id = $1
	ORDER BY oi.id`

	order := &model.Order{}
	err := r.conn.GetContext(ctx, order, queryOrderInfo, orderId)
	if err != nil {
		return nil, err
	}

	items := make([]*model.OrderItem, 0)
	rows, err := r.conn.QueryContext(ctx, queryOrderItems, orderId)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		item := &model.OrderItem{}
		if err = rows.Scan(
			&item.Id, &item.OrderId, &item.Count,
			&item.Product.Id, &item.WebsiteAlias, &item.Name, &item.Description, &item.Price,
			pq.Array(&item.ImageIds), &item.Active, pq.Array(&item.Tags),
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if rows.Err() != nil {
		return nil, err
	}

	order.Items = items
	return order, nil
}
