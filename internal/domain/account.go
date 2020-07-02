package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrAccountNotFound = errors.New("account not found")
)

type Account struct {
	ID          string            `json:"id" bson:"_id"`
	Name        string            `json:"name" bson:"name"`
	Exchange    string            `json:"exchange" bson:"exchange"`
	Credentials map[string]string `json:"credentials" bson:"credentials"`
	CreatedAt   time.Time         `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt" bson:"updatedAt"`
}

type AccountCriteria struct {
	Ids []string
}

type AccountRepository interface {
	GetMatching(ctx context.Context, criteria AccountCriteria) ([]*Account, error)
}
