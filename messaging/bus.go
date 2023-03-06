package messaging

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/maestre3d/coinlog/codec"
	"github.com/maestre3d/coinlog/domain"
	"github.com/maestre3d/coinlog/identifier"
	"github.com/modern-go/reflect2"
	"github.com/rs/zerolog/log"
)

var (
	ErrInvalidSchemaType = errors.New("bus: invalid schema type")
	ErrSchemaNotFound    = errors.New("bus: schema not found")
)

type schemaMetadata struct {
	StreamName string
}

type Bus struct {
	writer            Writer
	reader            Reader
	baseCtx           context.Context
	baseCtxCancelFunc context.CancelFunc

	schemaRegistry   map[string]schemaMetadata
	schemaRegistryMu sync.RWMutex

	readerRegistry   []ReaderTask
	readerRegistryMu sync.Mutex

	encodeFunc  codec.EncodeFunc
	decodeFunc  codec.DecodeFunc
	contentType string

	identifierFunc identifier.FactoryFunc
}

func NewBus(w Writer, r Reader) *Bus {
	return &Bus{
		writer:           w,
		reader:           r,
		baseCtx:          nil,
		schemaRegistry:   map[string]schemaMetadata{},
		schemaRegistryMu: sync.RWMutex{},
		encodeFunc:       codec.EncodeJSON,
		decodeFunc:       codec.DecodeJSON,
		contentType:      codec.ContentTypeJSON,
		identifierFunc:   identifier.NewKSUID,
	}
}

func (b *Bus) Register(stream string, v any) error {
	typeOf := reflect2.TypeOf(v)
	if _, ok := typeOf.(reflect2.StructType); !ok {
		return ErrInvalidSchemaType
	}

	b.schemaRegistryMu.Lock()
	defer b.schemaRegistryMu.Unlock()
	b.schemaRegistry[typeOf.String()] = schemaMetadata{
		StreamName: stream,
	}
	return nil
}

func (b *Bus) Subscribe(stream, group string, subscriberFunc SubscriberFunc) {
	b.readerRegistryMu.Lock()
	defer b.readerRegistryMu.Unlock()
	b.readerRegistry = append(b.readerRegistry, ReaderTask{
		Stream:  stream,
		GroupID: group,
		Handler: subscriberFunc,
	})
}

func (b *Bus) Start() {
	b.baseCtx, b.baseCtxCancelFunc = context.WithCancel(context.Background())
	for _, t := range b.readerRegistry {
		go func(task ReaderTask) {
			if err := b.reader.Read(b.baseCtx, task); err != nil {
				log.Err(err).Msg("bus: failed to read")
			}
		}(t)
	}
}

func (b *Bus) Shutdown() {
	b.baseCtxCancelFunc()
}

func (b *Bus) newMessageBatch(events []domain.Event) ([]Message, error) {
	buf := make([]Message, 0, len(events))
	for _, e := range events {
		metadata, ok := b.schemaRegistry[reflect2.TypeOf(e).String()]
		if !ok {
			return nil, ErrSchemaNotFound
		}
		data, err := b.encodeFunc(e)
		if err != nil {
			return nil, err
		}
		id, err := b.identifierFunc()
		if err != nil {
			return nil, err
		}
		now := time.Now().UTC()

		buf = append(buf, Message{
			ID:          id,
			Time:        now,
			TimeUsec:    now.UnixMicro(),
			StreamName:  metadata.StreamName,
			Key:         e.GetKey(),
			ContentType: b.contentType,
			Data:        data,
		})
	}
	return buf, nil
}

func (b *Bus) Publish(ctx context.Context, events []domain.Event) error {
	b.schemaRegistryMu.RLock()
	defer b.schemaRegistryMu.RUnlock()

	batch, err := b.newMessageBatch(events)
	if err != nil {
		return err
	}
	return b.writer.Write(ctx, batch)
}
