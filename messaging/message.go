package messaging

import "time"

type Message struct {
	ID          string    `json:"message_id"`
	Time        time.Time `json:"message_time"`
	TimeUsec    int64     `json:"message_time_microseconds"`
	StreamName  string    `json:"stream_name"`
	Key         string    `json:"message_key"`
	ContentType string    `json:"content_type"`
	Data        []byte    `json:"data"`
}
