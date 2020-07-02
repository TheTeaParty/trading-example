package domain

import (
	"context"
	"errors"

	exDomain "github.com/witchery-io/go-exchanges/pkg/domain"
)

var (
	ErrBalanceNotFound = errors.New("balance not found")
)

type Balance struct {
	ID        string
	Currency  string
	Name      string
	AccountID string
	Total     int64
	Exchange  string
}

type BalanceCriteria struct {
	AccountIDs []string
	Names      []string
	Exchanges  []string
}

type BalanceRepository interface {
	GetByCurrencyExchangeAndName(ctx context.Context, currency, exchange, name string) (*Balance, error)
	Create(ctx context.Context, balance *Balance) error
	UpdateByCurrencyExchangeAndName(ctx context.Context, currency, exchange, name string, balance *Balance) error
	Update(ctx context.Context, id string, balance *Balance) error
	Delete(ctx context.Context, id string) error
	GetMatching(ctx context.Context, criteria BalanceCriteria) ([]*Balance, error)
}

func BalanceToDomain(balance exDomain.Balance) Balance {
	return Balance{
		Currency:  balance.Currency.String(),
		Name:      balance.Name,
		AccountID: balance.AccountID,
		Total:     balance.Total,
		Exchange:  balance.Exchange,
	}
}
