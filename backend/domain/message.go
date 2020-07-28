package domain

type MessageType uint8

const (
	Text MessageType = iota
	File
)

// Message is a message which one client sends to another
type Message struct {
	Type    MessageType
	Payload string
}

type MessageUsecase interface {
	CreateQueue() (string, error)
	SendMessage(queueName string, message Message) error
}

type MessageQueue interface {
	Create() (string, error)
	// Read() <-chan string
	Write(queueName string, json []byte) error
}
