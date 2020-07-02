package handler

import (
	"context"

	"sources.witchery.io/simba/trading/internal/domain"
	tradingAPI "sources.witchery.io/simba/trading/pkg/api/proto"
)

func (h *handler) GetPositions(ctx context.Context,
	req *tradingAPI.PositionRequest, rsp *tradingAPI.PositionResponse) error {
	criteria := domain.PositionCriteria{AccountIDs: req.AccountIds}
	positions, err := h.app.PositionRepository.GetMatching(ctx, criteria)
	if err != nil {
		rsp.Status = h.proceedError(err, "Error getting positions", map[string]interface{}{
			"criteria": criteria,
		})
		return err
	}

	rsp.Status = tradingAPI.Status_OK
	for _, p := range positions {
		rsp.Positions = append(rsp.Positions, &tradingAPI.Position{
			Id:                   p.ID,
			ExternalId:           p.ExternalID,
			AccountId:            p.AccountID,
			Exchange:             p.Exchange,
			Pair:                 p.Pair,
			Status:               p.Status,
			Amount:               p.Amount,
			BasePrice:            p.BasePrice,
			MarginFunding:        p.MarginFunding,
			MarginFundingType:    p.MarginFundingType,
			ProfitLoss:           p.ProfitLoss,
			ProfitLossPercentage: p.ProfitLossPercentage,
			CreatedAt:            p.CreatedAt.Unix(),
			UpdatedAt:            p.UpdatedAt.Unix(),
		})
	}

	return nil
}
