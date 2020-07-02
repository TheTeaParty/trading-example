package internal

import (
	"sources.witchery.io/packages/infrastructure/micro"
	"sources.witchery.io/simba/trading/internal/domain"
)

type Application struct {
	*micro.Application

	AccountRepository  domain.AccountRepository
	PositionRepository domain.PositionRepository
	OrderRepository    domain.OrderRepository
	GroupRepository    domain.GroupRepository
	BalanceRepository  domain.BalanceRepository
}
