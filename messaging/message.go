package messaging

type Message struct {
	ID   string `json:"message_id"`
	Data []byte `json:"data"`
}
