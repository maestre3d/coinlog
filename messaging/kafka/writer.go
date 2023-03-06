package kafka

import (
	"context"

	"github.com/maestre3d/coinlog/messaging"
	"github.com/segmentio/kafka-go"
)

type Writer struct {
	kafkaWriter *kafka.Writer
}

var _ messaging.Writer = Writer{}

func NewWriter(cfg Config) (Writer, func()) {
	kafkaWriter := &kafka.Writer{
		Addr:                   kafka.TCP(cfg.Address),
		Balancer:               &kafka.CRC32Balancer{}, // guarantees ordering and exactly-once semantics
		MaxAttempts:            10,
		WriteBackoffMin:        0,
		WriteBackoffMax:        0,
		BatchSize:              0,
		BatchBytes:             0,
		BatchTimeout:           0,
		ReadTimeout:            0,
		WriteTimeout:           cfg.WriterTimeout,
		RequiredAcks:           kafka.RequireAll,
		Async:                  false,
		Completion:             nil,
		Compression:            kafka.Gzip,
		Logger:                 nil,
		ErrorLogger:            nil,
		Transport:              nil,
		AllowAutoTopicCreation: cfg.EnableTopicGeneration,
	}
	cleanup := func() {
		_ = kafkaWriter.Close()
	}
	return Writer{
		kafkaWriter: kafkaWriter,
	}, cleanup
}

func (w Writer) Write(ctx context.Context, msgs []messaging.Message) error {
	return w.kafkaWriter.WriteMessages(ctx, newMessageBatch(msgs)...)
}
