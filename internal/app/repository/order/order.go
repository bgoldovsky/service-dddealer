package order

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/bgoldovsky/dealer/service-dealer/internal/app/models/order"
	"github.com/bgoldovsky/dealer/service-dealer/internal/app/models/order/item"
	"github.com/jackc/pgx/v4"
)

var (
	ErrOrderNotFount = errors.New("order not found")
	ErrItemsNotFount = errors.New("items not found")
)

type queryer interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type repository struct {
	database queryer
}

type Repository interface {
	Find(ctx context.Context, id order.ID) (*order.Order, error)
	Save(ctx context.Context, order *order.Order) error
}

func New(database queryer) *repository {
	return &repository{
		database: database,
	}
}

func (r *repository) Find(ctx context.Context, orderID order.ID) (*order.Order, error) {
	sql := "select o.id, o.status, o.workflow, o.updated_at, o.created_at from orders o where id = $1"
	row := r.database.QueryRow(ctx, sql, orderID)
	var o order.Order
	err := row.Scan(
		&o.ID,
		&o.Status,
		&o.Workflow,
		&o.Created,
		&o.Updated,
	)
	if isEmpty(err) {
		return nil, ErrOrderNotFount
	}

	if err != nil {
		return nil, fmt.Errorf("query order error: %v", err)
	}

	items, err := r.findItems(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("query items error: %v", err)
	}

	o.Items = items
	return &o, nil
}

func (r *repository) findItems(ctx context.Context, orderID order.ID) ([]item.Item, error) {
	sql := "select i.id, i.seller_id, i.price, i.discount, i.name, i.updated_at, i.created_at from items i where i.order_id = $1"
	rows, err := r.database.Query(ctx, sql, orderID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var items []item.Item

	for rows.Next() {
		var i item.Item
		err := rows.Scan(
			&i.ID,
			&i.SellerID,
			&i.Price,
			&i.Discount,
			&i.Name,
			&i.Updated,
			&i.Created,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	if len(items) == 0 {
		return nil, ErrItemsNotFount
	}

	return items, nil
}

func (r *repository) Save(ctx context.Context, o *order.Order) error {
	tx, err := r.database.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	var attributes = map[string]interface{}{
		"status":     o.Status,
		"workflow":   o.Workflow,
		"created_at": o.Created,
		"updated_at": o.Updated,
	}

	if o.ID != 0 {
		attributes["id"] = o.ID
	}

	var (
		i           uint8
		values      []interface{}
		columns     []string
		placeholder []string
		update      []string
	)

	for k, v := range attributes {
		i += 1
		values = append(values, v)
		columns = append(columns, fmt.Sprintf("%q", k))
		placeholder = append(placeholder, fmt.Sprintf("$%d", i))
		update = append(update, fmt.Sprintf("%s=excluded.%s", k, k))
	}

	sql := fmt.Sprintf(`insert into "orders" (%s) values (%s) on conflict (id) do update set %s returning id`,
		strings.Join(columns, ","),
		strings.Join(placeholder, ","),
		strings.Join(update, ","))

	row := tx.QueryRow(ctx, sql, values...)
	err = row.Scan(&o.ID)
	if isEmpty(err) {
		return ErrOrderNotFount
	}

	if err != nil {
		return fmt.Errorf("put order error: %v", err)
	}

	for _, i := range o.Items {
		err := r.saveItem(ctx, tx, o.ID, i)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("put order transaction error: %v", err)
	}

	return nil
}

func (r *repository) saveItem(ctx context.Context, tx pgx.Tx, orderID order.ID, i item.Item) error {
	var attributes = map[string]interface{}{
		"id":         i.ID,
		"order_id":   orderID,
		"seller_id":  i.SellerID,
		"price":      i.Price,
		"discount":   i.Discount,
		"name":       i.Name,
		"created_at": i.Created,
		"updated_at": i.Updated,
	}

	var (
		j           uint8
		values      []interface{}
		columns     []string
		placeholder []string
		update      []string
	)

	for k, v := range attributes {
		j += 1
		values = append(values, v)
		columns = append(columns, fmt.Sprintf("%q", k))
		placeholder = append(placeholder, fmt.Sprintf("$%d", j))
		update = append(update, fmt.Sprintf("%s=excluded.%s", k, k))
	}

	sql := fmt.Sprintf(`insert into "items" (%s) values (%s) on conflict (id) do update set %s returning id`,
		strings.Join(columns, ","),
		strings.Join(placeholder, ","),
		strings.Join(update, ","))

	row := tx.QueryRow(ctx, sql, values...)
	return row.Scan(&i.ID)
}

func isEmpty(err error) bool {
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			return true
		}
	}
	return false
}
