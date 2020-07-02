package order

import (
	"context"
	"time"

	"github.com/google/uuid"
	"sources.witchery.io/simba/trading/internal/domain"
)

type memoryRepository struct {
	orders []*domain.Order
}

func (r *memoryRepository) GetWithExternalID(ctx context.Context,
	externalID string, exchange string) (*domain.Order, error) {
	for i, o := range r.orders {
		if o.ExternalID == externalID && o.Exchange == exchange {
			return r.orders[i], nil
		}
	}

	return nil, domain.ErrOrderNotFound
}

func (r *memoryRepository) Create(ctx context.Context, order *domain.Order) error {
	order.ID = uuid.New().String()

	r.orders = append(r.orders, order)

	return nil
}

func (r *memoryRepository) UpdateWithExternalID(ctx context.Context,
	externalID string, exchange string, order *domain.Order) error {
	for i, o := range r.orders {
		if o.ExternalID == externalID && o.Exchange == exchange {
			order.UpdatedAt = time.Now()
			r.orders[i] = order
			return nil
		}
	}

	return domain.ErrOrderNotFound
}

func (r *memoryRepository) DeleteWithExternalID(ctx context.Context, externalID string, exchange string) error {
	for i, o := range r.orders {
		if o.ExternalID == externalID && o.Exchange == exchange {
			r.orders[len(r.orders)-1], r.orders[i] = r.orders[i], r.orders[len(r.orders)-1]
			r.orders = r.orders[:len(r.orders)-1]
			return nil
		}
	}

	return domain.ErrPositionNotFound
}

func (r *memoryRepository) GetMatching(ctx context.Context, criteria domain.OrderCriteria) ([]*domain.Order, error) {
	orders := make([]*domain.Order, 0)
	for _, o := range r.orders {
		isFound := false
		for _, a := range criteria.AccountIds {
			if a == o.AccountID {
				isFound = true
				break
			}
		}

		if !isFound && len(criteria.AccountIds) > 0 {
			continue
		}

		orders = append(orders, o)
	}

	return orders, nil
}

func NewOrderMemoryRepository() domain.OrderRepository {
	return &memoryRepository{orders: []*domain.Order{}}
}
