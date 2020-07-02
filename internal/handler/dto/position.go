package dto

import (
	"time"

	"sources.witchery.io/simba/trading/internal/domain"
	tradingAPI "sources.witchery.io/simba/trading/pkg/api/proto"
)

func PositionFromDTO(position *tradingAPI.Position) domain.Position {
	return domain.Position{
		ID:                   position.Id,
		ExternalID:           position.ExternalId,
		AccountID:            position.AccountId,
		Exchange:             position.Exchange,
		Pair:                 position.Pair,
		Status:               position.Status,
		Amount:               position.Amount,
		BasePrice:            position.BasePrice,
		MarginFunding:        position.MarginFunding,
		MarginFundingType:    position.MarginFundingType,
		ProfitLoss:           position.ProfitLoss,
		ProfitLossPercentage: position.ProfitLossPercentage,
		CreatedAt:            time.Unix(position.CreatedAt, 0),
		UpdatedAt:            time.Unix(position.UpdatedAt, 0),
	}
}

func PositionToDTO(position domain.Position) *tradingAPI.Position {
	return &tradingAPI.Position{
		Id:                   position.ID,
		ExternalId:           position.ExternalID,
		AccountId:            position.AccountID,
		Exchange:             position.Exchange,
		Pair:                 position.Pair,
		Status:               position.Status,
		Amount:               position.Amount,
		BasePrice:            position.BasePrice,
		MarginFunding:        position.MarginFunding,
		MarginFundingType:    position.MarginFundingType,
		ProfitLoss:           position.ProfitLoss,
		ProfitLossPercentage: position.ProfitLossPercentage,
		CreatedAt:            position.CreatedAt.Unix(),
		UpdatedAt:            position.UpdatedAt.Unix(),
	}
}
