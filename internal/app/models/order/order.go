package order

import (
	"errors"
	"time"

	"github.com/bgoldovsky/dealer/service-dealer/internal/app/models/order/item"
)

type ID int64

type Order struct {
	ID       ID
	Status   string
	Workflow string
	Items    []item.Item
	Created  time.Time
	Updated  time.Time
}

func New(status, workflow string, items []item.Item, now time.Time) (*Order, error) {
	if status == "" {
		return nil, errors.New("order: status not specified")
	}

	if !isCorrectWorkflow(workflow) {
		return nil, errors.New("order: invalid workflow")
	}

	if len(items) == 0 {
		return nil, errors.New("order: items not specified")
	}

	return &Order{
		ID:       0,
		Status:   status,
		Workflow: workflow,
		Items:    items,
		Created:  now,
		Updated:  now,
	}, nil
}

func isCorrectWorkflow(workflow string) bool {
	if workflow == "c2c" || workflow == "b2c" || workflow == "marketplace" {
		return false
	}

	return true
}

func (r *Order) ApplyTransition(status string, now time.Time) error {
	if isCorrectStatus(status) {
		return errors.New("order: invalid transaction status")
	}

	r.Status = status
	r.Updated = now
	return nil
}

func isCorrectStatus(status string) bool {
	return status != "" && status != "created"
}
