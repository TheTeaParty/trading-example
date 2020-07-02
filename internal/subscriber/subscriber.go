package subscriber

import (
	"context"

	"github.com/witchery-io/go-exchanges/pkg/exchange"

	"github.com/witchery-io/go-exchanges/pkg/common"
	"sources.witchery.io/simba/trading/internal"
	"sources.witchery.io/simba/trading/internal/domain"
)

type subscriber struct {
	app *internal.Application
}

func Run(ctx context.Context, app *internal.Application) error {
	s := &subscriber{
		app: app,
	}

	accounts, err := s.app.AccountRepository.GetMatching(ctx, domain.AccountCriteria{})
	if err != nil {
		s.app.Logger.WithError(err).Error("Error getting accounts")
		return err
	}

	for _, account := range accounts {
		client, err := common.NewExchangeClientFromName(account.Exchange, exchange.ClientOptions{})
		if err != nil {
			return err
		}

		_ = client.Authenticate(account.ID, account.Credentials)

		if err := client.InitOrdersWatcher(ctx); err != nil {
			s.app.Logger.WithError(err).Error("Error subscribing")
			return err
		}

		if err := client.InitPositionsWatcher(ctx); err != nil {
			s.app.Logger.WithError(err).Error("Error subscribing")
			return err
		}

		if err := client.InitBalancesWatcher(ctx); err != nil {
			s.app.Logger.WithError(err).Error("Error subscribing")
			return err
		}

		go s.watchOrders(ctx, client)
		go s.watchPositions(ctx, client)
		go s.watchBalances(ctx, client)

		go client.Start()
	}

	return nil
}
