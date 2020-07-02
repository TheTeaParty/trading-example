package dto

import (
	"sources.witchery.io/simba/trading/internal/domain"
	tradingAPI "sources.witchery.io/simba/trading/pkg/api/proto"
)

func OrderFromNewDTO(order *tradingAPI.NewOrder, account domain.Account) domain.Order {
	return domain.Order{
		Direction:      order.Direction,
		Context:        order.Context,
		Type:           order.Type,
		Pair:           order.Pair,
		OriginalAmount: order.Amount,
		Price:          order.Price,
		AccountID:      account.ID,
		Exchange:       account.Exchange,
	}
}

func OrderToDTO(o domain.Order) *tradingAPI.Order {
	return &tradingAPI.Order{
		Id:                    o.ID,
		ExternalId:            o.ExternalID,
		Direction:             o.Direction,
		Context:               o.Context,
		Type:                  o.Type,
		Pair:                  o.Pair,
		OriginalAmount:        o.OriginalAmount,
		RemainingAmount:       o.RemainingAmount,
		Price:                 o.Price,
		AverageExecutionPrice: o.AverageExecutionPrice,
		OpenedAt:              o.OpenedAt.Unix(),
		UpdatedAt:             o.UpdatedAt.Unix(),
		CanceledAt:            o.CanceledAt.Unix(),
		AccountId:             o.AccountID,
		Status:                o.Status,
		Exchange:              o.Exchange,
	}
}
