package group

import (
	"context"

	"github.com/google/uuid"
	"sources.witchery.io/simba/trading/internal/domain"
)

type memoryRepository struct {
	groups []*domain.Group
}

func (r *memoryRepository) Create(ctx context.Context, group *domain.Group) error {
	group.ID = uuid.New().String()
	r.groups = append(r.groups, group)

	return nil
}

func (r *memoryRepository) Update(ctx context.Context, groupID string, group *domain.Group) error {
	for i, g := range r.groups {
		if g.ID == groupID {
			group.ID = groupID
			r.groups[i] = group
			return nil
		}
	}

	return domain.ErrGroupNotFound
}

func (r *memoryRepository) Delete(ctx context.Context, groupID string) error {
	for i, g := range r.groups {
		if g.ID == groupID {
			r.groups = append(r.groups[:i], r.groups[i+1:]...)
			return nil
		}
	}

	return domain.ErrGroupNotFound
}

func (r *memoryRepository) GetByID(ctx context.Context, groupID string) (*domain.Group, error) {
	for i, g := range r.groups {
		if g.ID == groupID {
			return r.groups[i], nil
		}
	}

	return nil, domain.ErrGroupNotFound
}

func (r *memoryRepository) GetMatching(ctx context.Context, criteria domain.GroupCriteria) ([]*domain.Group, error) {
	return r.groups, nil
}

func NewGroupMemoryRepository() domain.GroupRepository {
	return &memoryRepository{groups: []*domain.Group{
		{
			ID:         uuid.New().String(),
			Name:       "Testing group",
			AccountIDs: []string{"3da3c70d-bfbf-47d7-a818-47baf8aad7f6", "5a46e66d-0293-4a6e-93f2-c0ef28009b1d"},
		},
	}}
}
