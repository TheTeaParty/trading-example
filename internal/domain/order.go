package domain

import (
	"context"
	"errors"
	"time"

	exDomain "github.com/witchery-io/go-exchanges/pkg/domain"
)

var (
	ErrOrderNotFound = errors.New("order not found")
)

type Order struct {
	ID                    string `json:"id" bson:"_id"`
	ExternalID            string `json:"externalId" bson:"externalId"`
	Direction             string `json:"direction" bson:"direction"`
	Context               string `json:"context" bson:"context"`
	Type                  string `json:"type" bson:"type"`
	Pair                  string `json:"pair" bson:"pair"`
	OriginalAmount        int64  `json:"originalAmount" bson:"originalAmount"`
	RemainingAmount       int64  `json:"remainingAmount" bson:"remainingAmount"`
	Price                 int64  `json:"price" bson:"price"`
	AverageExecutionPrice int64
	OpenedAt              time.Time
	UpdatedAt             time.Time
	CanceledAt            time.Time
	AccountID             string
	Status                string
	Exchange              string
}

type OrderCriteria struct {
	AccountIds []string
}

type OrderRepository interface {
	Create(ctx context.Context, order *Order) error
	GetWithExternalID(ctx context.Context, externalID string, exchange string) (*Order, error)
	UpdateWithExternalID(ctx context.Context, externalID string, exchange string, order *Order) error
	DeleteWithExternalID(ctx context.Context, externalID string, exchange string) error
	GetMatching(ctx context.Context, criteria OrderCriteria) ([]*Order, error)
}

func OrderFromDomain(order Order) (*exDomain.Order, error) {
	oDirection, err := exDomain.ParseOrderDirection(order.Direction)
	if err != nil {
		return nil, err
	}
	oContext, err := exDomain.ParseOrderContext(order.Context)
	if err != nil {
		return nil, err
	}
	oType, err := exDomain.ParseOrderType(order.Type)
	if err != nil {
		return nil, err
	}

	// @todo fix
	pair := exDomain.NewCurrencyPairFrom2Currencies("BTC", "USD")

	return &exDomain.Order{
		Direction:      oDirection,
		Context:        oContext,
		Type:           oType,
		Pair:           pair,
		OriginalAmount: order.OriginalAmount,
		Price:          order.Price,
		AccountID:      order.AccountID,
		Exchange:       order.Exchange,
	}, nil
}

func OrderToDomain(order *exDomain.Order) (Order, error) {
	return Order{
		ID:                    order.OrderNumber,
		ExternalID:            order.OrderNumber,
		Direction:             order.Direction.String(),
		Context:               order.Context.String(),
		Type:                  order.Type.String(),
		Pair:                  order.Pair.String(),
		OriginalAmount:        order.OriginalAmount,
		RemainingAmount:       order.RemainingAmount,
		Price:                 order.Price,
		AverageExecutionPrice: order.AverageExecutionPrice,
		OpenedAt:              order.OpenedAt,
		UpdatedAt:             order.UpdatedAt,
		CanceledAt:            order.CanceledAt,
		AccountID:             order.AccountID,
		Status:                order.Status.String(),
		Exchange:              order.Exchange,
	}, nil
}
