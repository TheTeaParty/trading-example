package subscriber

import (
	"context"

	"github.com/micro/go-micro/broker"
	"github.com/witchery-io/go-exchanges/pkg/exchange"
	"sources.witchery.io/simba/trading/internal/domain"
	"sources.witchery.io/simba/trading/pkg/topic"
)

func (s *subscriber) watchBalances(ctx context.Context, client exchange.Client) {
	for t := range client.BalanceEvents(ctx) {
		balance := domain.BalanceToDomain(t.Balance)

		_, err := s.app.BalanceRepository.GetByCurrencyExchangeAndName(ctx,
			balance.Currency, balance.Exchange, balance.Name)
		switch err {
		case nil:
			err := s.app.BalanceRepository.Create(ctx, &balance)
			if err != nil {
				s.app.Logger.WithError(err).Error("Error adding balance")
				continue
			}
		case domain.ErrBalanceNotFound:
			err = s.app.BalanceRepository.UpdateByCurrencyExchangeAndName(ctx,
				balance.Currency, balance.Exchange, balance.Name, &balance)
			if err != nil {
				s.app.Logger.WithError(err).Error("Error updating balance")
				continue
			}
		default:
			s.app.Logger.
				WithField("balance", balance).
				WithError(err).Error("Cant't get balance")
			continue
		}

		_ = s.app.Service.Options().Broker.Publish(topic.PositionUpdateTopic, &broker.Message{
			Header: map[string]string{
				"accountID": balance.AccountID,
			},
		})
	}
}
