package handler

import (
	"context"

	"github.com/witchery-io/go-exchanges/pkg/exchange"

	"github.com/witchery-io/go-exchanges/pkg/common"
	"sources.witchery.io/simba/trading/internal/domain"
	"sources.witchery.io/simba/trading/internal/handler/dto"
	tradingAPI "sources.witchery.io/simba/trading/pkg/api/proto"
)

func (h *handler) CreateOrder(ctx context.Context, req *tradingAPI.NewOrder, rsp *tradingAPI.OrderResponse) error {
	//@todo check for balance before submit
	accounts, err := h.app.AccountRepository.GetMatching(ctx, domain.AccountCriteria{Ids: req.AccountIds})
	if err != nil {
		rsp.Status = h.proceedError(err, "Error getting account", map[string]interface{}{
			"ids": req.AccountIds,
		})
		return err
	}

	rsp.Status = tradingAPI.Status_OK

	for _, a := range accounts {
		client, err := common.NewExchangeClientFromName(a.Exchange, exchange.ClientOptions{})
		if err != nil {
			rsp.Status = h.proceedError(err, "Error getting exchange client", map[string]interface{}{
				"account": a,
			})
			continue
		}

		_ = client.Authenticate(a.ID, a.Credentials)

		order, err := domain.OrderFromDomain(dto.OrderFromNewDTO(req, *a))
		if err != nil {
			_ = h.proceedError(err, "Error parsing order", map[string]interface{}{
				"order": order,
			})
			continue
		}

		err = client.SubmitOrder(ctx, order)
		if err != nil {
			_ = h.proceedError(err, "Error submitting order", map[string]interface{}{
				"order": order,
			})
			continue
		}
	}

	return nil
}

func (h *handler) GetOrders(ctx context.Context, req *tradingAPI.OrderRequest, rsp *tradingAPI.OrderResponse) error {
	criteria := domain.OrderCriteria{AccountIds: req.AccountIds}
	orders, err := h.app.OrderRepository.GetMatching(ctx, criteria)
	if err != nil {
		rsp.Status = h.proceedError(err, "Error submitting order", map[string]interface{}{
			"criteria": criteria,
		})
		return err
	}

	rsp.Status = tradingAPI.Status_OK
	for _, o := range orders {
		rsp.Orders = append(rsp.Orders, dto.OrderToDTO(*o))
	}

	return nil
}
