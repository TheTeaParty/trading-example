package handler

import (
	"context"

	"sources.witchery.io/simba/trading/internal/domain"
	tradingAPI "sources.witchery.io/simba/trading/pkg/api/proto"
)

func (h *handler) GetAccounts(ctx context.Context,
	req *tradingAPI.AccountRequest, rsp *tradingAPI.AccountResponse) error {
	accounts, err := h.app.AccountRepository.GetMatching(ctx, domain.AccountCriteria{Ids: []string{}})
	if err != nil {
		rsp.Status = h.proceedError(err, "Error getting accounts", map[string]interface{}{})
		return err
	}

	for _, a := range accounts {
		rsp.Accounts = append(rsp.Accounts, &tradingAPI.Account{
			Id:          a.ID,
			Name:        a.Name,
			Exchange:    a.Exchange,
			Credentials: a.Credentials,
		})
	}

	return nil
}
