package kafka

import (
	"github.com/maestre3d/coinlog/messaging"
	"github.com/segmentio/kafka-go"
)

func newMessageBatch(msgs []messaging.Message) []kafka.Message {
	buf := make([]kafka.Message, 0, len(msgs))
	for _, msg := range msgs {
		buf = append(buf, kafka.Message{
			Topic: msg.StreamName,
			Key:   []byte(msg.Key),
			Value: msg.Data,
			Headers: []kafka.Header{
				{
					Key:   "message_id",
					Value: []byte(msg.ID),
				},
				{
					Key:   "content_type",
					Value: []byte(msg.ContentType),
				},
			},
			Time: msg.Time,
		})
	}
	return buf
}

func unmarshalMessage(msg kafka.Message) messaging.Message {
	headers := make(map[string]string)
	for _, h := range msg.Headers {
		headers[h.Key] = string(h.Value)
	}
	return messaging.Message{
		ID:          headers["message_id"],
		Time:        msg.Time,
		TimeUsec:    msg.Time.UnixMicro(),
		StreamName:  msg.Topic,
		Key:         string(msg.Key),
		ContentType: headers["content_type"],
		Data:        msg.Value,
	}
}
