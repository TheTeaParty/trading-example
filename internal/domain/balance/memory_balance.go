package balance

import (
	"context"

	"github.com/google/uuid"
	"sources.witchery.io/simba/trading/internal/domain"
)

type memoryRepository struct {
	balances []*domain.Balance
}

func (r *memoryRepository) UpdateByCurrencyExchangeAndName(ctx context.Context,
	currency, exchange, name string, balance *domain.Balance) error {
	for i, b := range r.balances {
		if b.Currency == currency && b.Exchange == exchange && b.Name == name {
			r.balances[i] = balance
			return nil
		}
	}

	return domain.ErrBalanceNotFound
}

func (r *memoryRepository) GetByCurrencyExchangeAndName(ctx context.Context, currency,
	exchange, name string) (*domain.Balance, error) {
	for _, b := range r.balances {
		if b.Currency == currency && b.Exchange == exchange && b.Name == name {
			return b, nil
		}
	}

	return nil, domain.ErrBalanceNotFound
}

func (r *memoryRepository) Create(ctx context.Context, balance *domain.Balance) error {
	balance.ID = uuid.New().String()
	r.balances = append(r.balances, balance)

	return nil
}

func (r *memoryRepository) Update(ctx context.Context, id string, balance *domain.Balance) error {
	for i, b := range r.balances {
		if b.ID == id {
			balance.ID = id
			r.balances[i] = balance
			return nil
		}
	}

	return domain.ErrBalanceNotFound
}

func (r *memoryRepository) Delete(ctx context.Context, id string) error {
	for i, b := range r.balances {
		if b.ID == id {
			r.balances = append(r.balances[:i], r.balances[i+1:]...)
			return nil
		}
	}

	return domain.ErrBalanceNotFound
}

func (r *memoryRepository) GetMatching(ctx context.Context,
	criteria domain.BalanceCriteria) ([]*domain.Balance, error) {
	var balances []*domain.Balance
	for _, b := range r.balances {
		for _, id := range criteria.AccountIDs {
			if b.AccountID == id {
				balances = append(balances, b)
				continue
			}
		}

		for _, name := range criteria.Names {
			if b.Name == name {
				balances = append(balances, b)
				continue
			}
		}

		for _, exchange := range criteria.Exchanges {
			if b.Exchange == exchange {
				balances = append(balances, b)
				continue
			}
		}
	}

	if len(criteria.Exchanges) == 0 && len(criteria.Names) == 0 && len(criteria.AccountIDs) == 0 {
		return r.balances, nil
	}

	return balances, nil
}

func NewBalanceMemoryRepository() domain.BalanceRepository {
	return &memoryRepository{balances: []*domain.Balance{}}
}
