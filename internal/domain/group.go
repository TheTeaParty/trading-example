package domain

import (
	"context"
	"errors"
)

var (
	ErrGroupNotFound = errors.New("group not found")
)

type Group struct {
	ID         string
	Name       string
	AccountIDs []string
}

type GroupCriteria struct {
}

type GroupRepository interface {
	GetByID(ctx context.Context, groupID string) (*Group, error)
	Create(ctx context.Context, group *Group) error
	Update(ctx context.Context, groupID string, group *Group) error
	Delete(ctx context.Context, groupID string) error
	GetMatching(ctx context.Context, criteria GroupCriteria) ([]*Group, error)
}
