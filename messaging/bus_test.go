package messaging

import (
	"context"
	"testing"

	"github.com/maestre3d/coinlog/domain"
)

type NoopWriter struct {
}

var _ Writer = NoopWriter{}

func (n NoopWriter) Write(_ context.Context, _ []Message) error {
	return nil
}

type NoopReader struct {
}

func (n NoopReader) Read(_ context.Context, _ ReaderTask) error {
	return nil
}

var _ Reader = NoopReader{}

type fooEvent struct {
	ID string
}

func (f fooEvent) GetKey() string {
	return "123"
}

func TestBus(t *testing.T) {
	b := NewBus(NoopWriter{}, NoopReader{})
	_ = b.Register("coinlog.foo", fooEvent{})
	_ = b.Publish(context.TODO(), []domain.Event{fooEvent{}})
}
