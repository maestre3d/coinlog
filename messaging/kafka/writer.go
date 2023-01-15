package kafka

import (
	"context"

	"github.com/maestre3d/coinlog/messaging"
)

type Writer struct {
}

var _ messaging.Writer = Writer{}

func (w Writer) Write(_ context.Context, _ ...any) error {
	return nil
}
