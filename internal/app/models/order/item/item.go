package item

import (
	"errors"
	"time"
)

type ID int64

type Item struct {
	ID       ID
	SellerID int64
	Price    int64
	Discount *int64
	Name     string
	Created  time.Time
	Updated  time.Time
	Deleted  *time.Time
}

func New(itemID ID, sellerID int64, price int64, discount *int64, name string) (*Item, error) {
	if itemID <= 0 {
		return nil, errors.New("item: itemID must be positive integer")
	}

	if sellerID <= 0 {
		return nil, errors.New("item: sellerID must be positive integer")
	}

	if price <= 0 {
		return nil, errors.New("item: price must be positive integer")
	}

	if name == "" {
		return nil, errors.New("item: name not specified")
	}

	return &Item{
		ID:       itemID,
		SellerID: sellerID,
		Price:    price,
		Discount: discount,
		Name:     name,
	}, nil
}

func (r *Item) TotalCost() int64 {
	price := r.Price

	if r.Discount != nil && *r.Discount < price {
		price = price - *r.Discount
	}

	return price
}
