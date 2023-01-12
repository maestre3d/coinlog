package domain

import "context"

type EventBus interface {
	Publish(ctx context.Context, events ...any) error
}
