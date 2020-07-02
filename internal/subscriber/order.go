package subscriber

import (
	"context"

	"github.com/micro/go-micro/broker"
	"github.com/witchery-io/go-exchanges/pkg/exchange"
	"sources.witchery.io/simba/trading/internal/domain"
	"sources.witchery.io/simba/trading/pkg/topic"
)

func (s *subscriber) watchOrders(ctx context.Context, client exchange.Client) {
	for t := range client.OrderEvents(ctx) {
		order, err := domain.OrderToDomain(&t.Order)
		if err != nil {
			s.app.Logger.
				WithField("order", t.Order).
				WithError(err).Error("Cant't parse order")
			continue
		}

		_, err = s.app.OrderRepository.GetWithExternalID(ctx, t.Order.OrderNumber, t.Order.Exchange)
		switch err {
		case nil:
			err := s.app.OrderRepository.UpdateWithExternalID(ctx, t.Order.OrderNumber, t.Order.Exchange, &order)
			if err != nil {
				s.app.Logger.
					WithField("order", order).
					WithError(err).Error("Error updating order")
				continue
			}
		case domain.ErrOrderNotFound:
			err := s.app.OrderRepository.Create(ctx, &order)
			if err != nil {
				s.app.Logger.
					WithField("order", order).
					WithError(err).Error("Error adding order")
				continue
			}
		default:
			s.app.Logger.
				WithField("order", order).
				WithError(err).Error("Error getting order")
		}

		err = s.app.Service.Options().Broker.Publish(topic.OrderUpdateTopic, &broker.Message{
			Header: map[string]string{
				"id": order.ID,
			},
		})
		if err != nil {
			s.app.Logger.
				WithField("order", order).
				WithError(err).Error("Cant't publish " + topic.OrderUpdateTopic)
		}
	}
}
