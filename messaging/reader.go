package messaging

import (
	"context"
)

type ReaderFunc func(ctx context.Context)
