package domain

import (
	"context"
)

type Criteria struct {
	Limit     int       `json:"limit"`
	PageToken PageToken `json:"page_token"`
}

type BasicService[T any] interface {
	Create(ctx context.Context, args any) error
	Update(ctx context.Context, args any) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (T, error)
	List(ctx context.Context, cr Criteria) ([]T, PageToken, error)
}
