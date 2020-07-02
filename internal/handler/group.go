package handler

import (
	"context"

	"sources.witchery.io/simba/trading/internal/domain"
	tradingAPI "sources.witchery.io/simba/trading/pkg/api/proto"
)

func (h *handler) GetGroups(ctx context.Context, req *tradingAPI.GroupRequest, rsp *tradingAPI.GroupResponse) error {
	criteria := domain.GroupCriteria{}
	groups, err := h.app.GroupRepository.GetMatching(ctx, criteria)
	if err != nil {
		rsp.Status = h.proceedError(err, "Error getting groups", map[string]interface{}{
			"criteria": criteria,
		})
		return err
	}

	for _, g := range groups {
		accounts, err := h.app.AccountRepository.GetMatching(ctx, domain.AccountCriteria{Ids: g.AccountIDs})
		if err != nil {
			rsp.Status = h.proceedError(err, "Error getting group accounts", map[string]interface{}{
				"ids": g.AccountIDs,
			})
			return err
		}

		group := &tradingAPI.Group{
			Id:         g.ID,
			Name:       g.Name,
			AccountIds: g.AccountIDs,
		}

		for _, a := range accounts {
			group.Accounts = append(group.Accounts, &tradingAPI.Account{
				Id:          a.ID,
				Name:        a.Name,
				Exchange:    a.Exchange,
				Credentials: a.Credentials,
			})
		}

		rsp.Groups = append(rsp.Groups, group)
	}

	return nil
}
