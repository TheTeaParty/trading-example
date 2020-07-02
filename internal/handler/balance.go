package handler

import (
	"context"

	"sources.witchery.io/simba/trading/internal/domain"
	tradingAPI "sources.witchery.io/simba/trading/pkg/api/proto"
)

func (h *handler) GetBalances(ctx context.Context,
	req *tradingAPI.BalanceRequest, rsp *tradingAPI.BalanceResponse) error {
	criteria := domain.BalanceCriteria{
		AccountIDs: req.AccountIds,
		Names:      req.Names,
		Exchanges:  req.Exchanges,
	}
	balances, err := h.app.BalanceRepository.GetMatching(ctx, criteria)
	if err != nil {
		rsp.Status = h.proceedError(err, "Error getting balances", map[string]interface{}{
			"criteria": criteria,
		})
		return err
	}

	for _, b := range balances {
		rsp.Balances = append(rsp.Balances, &tradingAPI.Balance{
			Currency:  b.Currency,
			Name:      b.Name,
			AccountId: b.AccountID,
			Total:     b.Total,
			Exchange:  b.Exchange,
		})
	}

	return nil
}
