package messaging

import "context"

type SubscriberFunc func(ctx context.Context, message Message) error

type ReaderTask struct {
	Stream  string
	GroupID string
	Handler SubscriberFunc
}

type Reader interface {
	Read(ctx context.Context, task ReaderTask) error
}
