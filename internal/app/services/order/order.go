package order

import (
	"context"
	"fmt"
	"time"

	"github.com/bgoldovsky/dealer/service-dealer/internal/app/logger"
	"github.com/bgoldovsky/dealer/service-dealer/internal/app/models/order"
	"github.com/bgoldovsky/dealer/service-dealer/internal/app/models/order/item"
	orderRepo "github.com/bgoldovsky/dealer/service-dealer/internal/app/repository/order"
)

type service struct {
	repo orderRepo.Repository
}

func New(repo orderRepo.Repository) *service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, workflow string, items []item.Item) (*order.Order, error) {
	o, err := order.New("new", workflow, items, time.Now())
	if err != nil {
		logger.Log.
			WithError(err).
			WithField("workflow", workflow).
			Error("create order: validation error")
		return nil, fmt.Errorf("create order. validation error: %v", err)
	}

	err = s.repo.Save(ctx, o)
	if err != nil {
		logger.Log.
			WithError(err).
			WithField("order", o).
			Error("create order: save error")
		return nil, fmt.Errorf("create order. save error: %v", err)
	}

	return o, nil
}

func (s *service) ApplyTransition(ctx context.Context, id order.ID, status string) error {
	o, err := s.repo.Find(ctx, id)
	if err != nil {
		logger.Log.
			WithError(err).
			WithField("id", id).
			Error("apply transaction: find error")
		return fmt.Errorf("apply transaction. find error: %v", err)
	}

	err = o.ApplyTransition(status, time.Now())
	if err != nil {
		logger.Log.
			WithError(err).
			WithField("status", status).
			Error("apply transaction: change state error")
		return fmt.Errorf("apply transaction. change state error: %v", err)
	}

	return nil
}
