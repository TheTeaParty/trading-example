package position

import (
	"context"
	"time"

	"github.com/google/uuid"
	"sources.witchery.io/simba/trading/internal/domain"
)

type memoryRepository struct {
	positions []*domain.Position
}

func (r *memoryRepository) GetWithExternalID(ctx context.Context,
	positionID string, exchange string) (*domain.Position, error) {
	for i, p := range r.positions {
		if p.ExternalID == positionID && p.Exchange == exchange {
			return r.positions[i], nil
		}
	}

	return nil, domain.ErrPositionNotFound
}

func (r *memoryRepository) Create(ctx context.Context, position *domain.Position) error {
	position.ID = uuid.New().String()
	position.CreatedAt = time.Now()
	position.UpdatedAt = time.Now()

	r.positions = append(r.positions, position)

	return nil
}

func (r *memoryRepository) UpdateWithExternalID(ctx context.Context, positionID string,
	exchange string, position *domain.Position) error {
	for i, p := range r.positions {
		if p.ExternalID == positionID && p.Exchange == exchange {
			position.UpdatedAt = time.Now()
			r.positions[i] = position
			return nil
		}
	}

	return domain.ErrPositionNotFound
}

func (r *memoryRepository) DeleteWithExternalID(ctx context.Context, positionID string, exchange string) error {
	for i, p := range r.positions {
		if p.ExternalID == positionID && p.Exchange == exchange {
			r.positions[len(r.positions)-1], r.positions[i] = r.positions[i], r.positions[len(r.positions)-1]
			r.positions = r.positions[:len(r.positions)-1]
			return nil
		}
	}

	return domain.ErrPositionNotFound
}

func (r *memoryRepository) GetMatching(ctx context.Context,
	criteria domain.PositionCriteria) ([]*domain.Position, error) {
	positions := make([]*domain.Position, 0)
	for _, p := range r.positions {
		isFound := false
		for _, a := range criteria.AccountIDs {
			if a == p.AccountID {
				isFound = true
				break
			}
		}

		if !isFound && len(criteria.AccountIDs) > 0 {
			continue
		}

		positions = append(positions, p)
	}

	return positions, nil
}

func NewPositionMemoryRepository() domain.PositionRepository {
	return &memoryRepository{positions: []*domain.Position{}}
}
