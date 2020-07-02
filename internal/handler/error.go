package handler

import (
	"sources.witchery.io/simba/trading/internal/domain"
	tradingAPI "sources.witchery.io/simba/trading/pkg/api/proto"
)

var errorStatuses = map[error]tradingAPI.Status{
	domain.ErrPositionNotFound: tradingAPI.Status_NOT_FOUND,
	domain.ErrOrderNotFound:    tradingAPI.Status_NOT_FOUND,
	domain.ErrBalanceNotFound:  tradingAPI.Status_NOT_FOUND,
	domain.ErrGroupNotFound:    tradingAPI.Status_NOT_FOUND,
	domain.ErrAccountNotFound:  tradingAPI.Status_NOT_FOUND,
}

func (h *handler) proceedError(err error, message string, data map[string]interface{}) tradingAPI.Status {
	log := h.app.Logger.WithError(err).WithFields(data)

	if err == nil {
		return tradingAPI.Status_OK
	}

	if status, ok := errorStatuses[err]; ok {
		log.Debug(message)
		return status
	}

	log.Error(message)
	return tradingAPI.Status_SERVER_ERROR
}
