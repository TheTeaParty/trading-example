package handler

import (
	"sources.witchery.io/simba/trading/internal"
	tradingAPI "sources.witchery.io/simba/trading/pkg/api/proto"
)

type handler struct {
	app *internal.Application
}

func NewTradingHandler(app *internal.Application) tradingAPI.TradingServiceHandler {
	return &handler{app: app}
}
