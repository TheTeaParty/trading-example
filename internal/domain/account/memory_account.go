package account

import (
	"context"
	"time"

	"sources.witchery.io/simba/trading/internal/domain"
)

type memoryRepository struct {
	accounts []*domain.Account
}

func (r *memoryRepository) GetMatching(ctx context.Context,
	criteria domain.AccountCriteria) ([]*domain.Account, error) {
	if len(criteria.Ids) == 0 {
		return r.accounts, nil
	}

	var accounts []*domain.Account
	for _, a := range r.accounts {
		for _, id := range criteria.Ids {
			if id == a.ID {
				accounts = append(accounts, a)
			}
		}
	}

	return accounts, nil
}

func NewAccountMemoryRepository() domain.AccountRepository {
	return &memoryRepository{accounts: []*domain.Account{
	}}
}
