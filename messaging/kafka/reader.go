package kafka

import (
	"context"

	"github.com/maestre3d/coinlog/messaging"
	"github.com/segmentio/kafka-go"
)

type Reader struct {
	cfg Config
}

var _ messaging.Reader = Reader{}

func NewReader(cfg Config) Reader {
	return Reader{
		cfg: cfg,
	}
}

func (r Reader) Read(ctx context.Context, task messaging.ReaderTask) error {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:                []string{r.cfg.Address},
		GroupID:                task.GroupID,
		GroupTopics:            nil,
		Topic:                  task.Stream,
		Partition:              0,
		Dialer:                 nil,
		QueueCapacity:          0,
		MinBytes:               0,
		MaxBytes:               0,
		MaxWait:                0,
		ReadBatchTimeout:       0,
		ReadLagInterval:        0,
		GroupBalancers:         nil,
		HeartbeatInterval:      0,
		CommitInterval:         0,
		PartitionWatchInterval: 0,
		WatchPartitionChanges:  false,
		SessionTimeout:         0,
		RebalanceTimeout:       0,
		JoinGroupBackoff:       0,
		RetentionTime:          0,
		StartOffset:            0,
		ReadBackoffMin:         0,
		ReadBackoffMax:         0,
		Logger:                 nil,
		ErrorLogger:            nil,
		IsolationLevel:         0,
		MaxAttempts:            0,
		OffsetOutOfRangeError:  false,
	})

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			break
		}
		if err = task.Handler(ctx, unmarshalMessage(msg)); err != nil {
			// TODO: Avoid commit on errors
			break
		}
	}

	return reader.Close()
}
