package messaging

import "context"

// Bus set of one or more brokers used by distributed systems to communicate each other in an asynchronous way.
type Bus struct {
	w       Writer
	baseCtx context.Context
}

func NewBus(w Writer) *Bus {
	return &Bus{
		w: w,
	}
}

func (b *Bus) Write(ctx context.Context, data ...any) error {
	//TODO implement me
	panic("implement me")
}

func (b *Bus) RegisterReader(streamName string, readerFunc ReaderFunc) {
	//TODO implement me
	panic("implement me")
}

func (b *Bus) Start() error {
	//TODO implement me
	panic("implement me")
}

func (b *Bus) Shutdown(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return nil
	}
}
