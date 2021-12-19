package message

import "time"

const (
	Topic string = "effective_bassoon"
)

// Message is used as body published for the queue
type Message struct {
	Name      string
	Content   string
	Timestamp int64
}

func New(name, content string) Message {
	return Message{
		Name:      name,
		Content:   content,
		Timestamp: time.Now().Unix(),
	}
}
