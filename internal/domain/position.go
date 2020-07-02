package domain

import (
	"context"
	"errors"
	"time"

	exDomain "github.com/witchery-io/go-exchanges/pkg/domain"
)

var (
	ErrPositionNotFound = errors.New("position not found")
)

type Position struct {
	ID                   string    `json:"id" bson:"_id"`
	ExternalID           string    `json:"externalId" bson:"externalId"`
	AccountID            string    `json:"accountId" bson:"accountId"`
	Exchange             string    `json:"exchange" bson:"exchange"`
	Pair                 string    `json:"pair" bson:"pair"`
	Status               string    `json:"status" bson:"status"`
	Amount               int64     `json:"amount" bson:"amount"`
	BasePrice            int64     `json:"basePrice" bson:"basePrice"`
	MarginFunding        float64   `json:"marginFunding" bson:"marginFunding"`
	MarginFundingType    int64     `json:"marginFundingType" bson:"marginFundingType"`
	ProfitLoss           int64     `json:"profitLoss" bson:"profitLoss"`
	ProfitLossPercentage float64   `json:"profitLossPercentage" bson:"profitLossPercentage"`
	CreatedAt            time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt" bson:"updatedAt"`
}

type PositionCriteria struct {
	AccountIDs []string
	Exchange   string
	ExternalID string
}

type PositionRepository interface {
	Create(ctx context.Context, position *Position) error
	GetWithExternalID(ctx context.Context, positionID string, exchange string) (*Position, error)
	UpdateWithExternalID(ctx context.Context, positionID string, exchange string, position *Position) error
	DeleteWithExternalID(ctx context.Context, positionID string, exchange string) error
	GetMatching(ctx context.Context, criteria PositionCriteria) ([]*Position, error)
}

func PositionToDomain(position exDomain.Position) Position {
	return Position{
		ExternalID:           position.ID,
		AccountID:            position.AccountID,
		Exchange:             position.Exchange,
		Pair:                 position.Pair.String(),
		Status:               position.Status.String(),
		Amount:               position.Amount,
		BasePrice:            position.BasePrice,
		MarginFunding:        position.MarginFunding,
		MarginFundingType:    position.MarginFundingType,
		ProfitLoss:           position.ProfitLoss,
		ProfitLossPercentage: position.ProfitLossPercentage,
	}
}

func PositionFromDomain(position Position) (*exDomain.Position, error) {
	// @todo change
	pair := exDomain.NewCurrencyPairFrom2Currencies(
		exDomain.Currency(position.Pair[0:3]), exDomain.Currency(position.Pair[3:]))

	status, err := exDomain.ParsePositionStatus(position.Status)
	if err != nil {
		return nil, err
	}

	return &exDomain.Position{
		ID:                   position.ExternalID,
		AccountID:            position.AccountID,
		Exchange:             position.Exchange,
		Pair:                 pair,
		Status:               status,
		Amount:               position.Amount,
		BasePrice:            position.BasePrice,
		MarginFunding:        position.MarginFunding,
		MarginFundingType:    position.MarginFundingType,
		ProfitLoss:           position.ProfitLoss,
		ProfitLossPercentage: position.ProfitLossPercentage,
	}, nil
}
