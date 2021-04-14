package order

import (
	"context"

	"github.com/bgoldovsky/dealer/service-dealer/internal/app/models/order/item"
)

type Service interface {
	Create(ctx context.Context, workflow string, items []item.Item) (*Order, error)
	ApplyTransition(ctx context.Context, id ID, transition string) error
}
