package main

import (
	"context"

	"github.com/Rhymond/go-money"
	"github.com/witchery-io/go-exchanges/pkg/domain"
	"sources.witchery.io/packages/infrastructure/micro"
	"sources.witchery.io/simba/trading/internal"
	"sources.witchery.io/simba/trading/internal/domain/account"
	"sources.witchery.io/simba/trading/internal/domain/balance"
	"sources.witchery.io/simba/trading/internal/domain/group"
	"sources.witchery.io/simba/trading/internal/domain/order"
	"sources.witchery.io/simba/trading/internal/domain/position"
	"sources.witchery.io/simba/trading/internal/handler"
	"sources.witchery.io/simba/trading/internal/subscriber"
	tradingAPI "sources.witchery.io/simba/trading/pkg/api/proto"
)

const (
	appName     = "simba"
	serviceName = "simba-trading"
)

func main() {
	for _, c := range domain.AllCurrencies {
		money.AddCurrency(c.String(), c.String(), "$1", ".", ",", 8)
	}

	money.AddCurrency("BTC", "BTC", "$1", ".", ",", 8)
	money.AddCurrency("ETH", "ETH", "$1", ".", ",", 8)

	app := &internal.Application{
		Application:        micro.NewService(appName, serviceName),
		AccountRepository:  account.NewAccountMemoryRepository(),
		PositionRepository: position.NewPositionMemoryRepository(),
		OrderRepository:    order.NewOrderMemoryRepository(),
		GroupRepository:    group.NewGroupMemoryRepository(),
		BalanceRepository:  balance.NewBalanceMemoryRepository(),
	}

	// Handle panics
	defer func() {
		if err := recover(); err != nil {
			app.Logger.WithError(err.(error)).Error("Service panicked")
		}
	}()

	if err := app.ConnectBroker(); err != nil {
		app.Logger.WithError(err).Fatal("Error connecting to broker")
	}

	if err := subscriber.Run(context.Background(), app); err != nil {
		app.Logger.WithError(err).Fatal("Error initiating subscriber")
	}

	if err := tradingAPI.RegisterTradingServiceHandler(app.Service.Server(), handler.NewTradingHandler(app)); err != nil {
		app.Logger.WithError(err).Fatal("Error initiating handler")
	}

	if err := app.Run(); err != nil {
		app.Logger.WithError(err).Fatal("Error starting service")
	}
}
