package messaging

import (
	"context"
)

type Writer interface {
	Write(ctx context.Context, msgs []Message) error
}
