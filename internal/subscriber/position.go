package subscriber

import (
	"context"

	"github.com/micro/go-micro/broker"
	"github.com/witchery-io/go-exchanges/pkg/exchange"
	"sources.witchery.io/simba/trading/internal/domain"
	"sources.witchery.io/simba/trading/pkg/topic"
)

func (s *subscriber) watchPositions(ctx context.Context, client exchange.Client) {
	for t := range client.PositionEvents(ctx) {
		position := domain.PositionToDomain(t.Position)

		// Making sure position does not duplicate itself in platform
		_, err := s.app.PositionRepository.GetWithExternalID(ctx, t.Position.ID, t.Position.Exchange)
		switch err {
		case nil:
			// Position already exist update it
			err := s.app.PositionRepository.UpdateWithExternalID(ctx, t.Position.ID, t.Position.Exchange, &position)
			if err != nil {
				s.app.Logger.
					WithField("position", position).
					WithError(err).Error("Error adding position")
				continue
			}
		case domain.ErrPositionNotFound:
			// Position does not exit create it
			err := s.app.PositionRepository.Create(ctx, &position)
			if err != nil {
				s.app.Logger.
					WithField("position", position).
					WithError(err).Error("Error updating position")
				continue
			}
		default:
			s.app.Logger.
				WithField("position", position).
				WithError(err).Error("Cant't get position")
			continue
		}

		err = s.app.Service.Options().Broker.Publish(topic.PositionUpdateTopic, &broker.Message{
			Header: map[string]string{
				"id": position.ID,
			},
		})
		if err != nil {
			s.app.Logger.
				WithField("position", position).
				WithError(err).Error("Cant't publish " + topic.PositionUpdateTopic)
		}
	}
}
