package queue

import (
	"log"

	"github.com/eyalch/pipeit/backend/domain"
	"github.com/streadway/amqp"
)

type messageQueue struct {
	ch *amqp.Channel
}

func NewMessageQueue(ch *amqp.Channel) domain.MessageQueue {
	return &messageQueue{ch}
}

func (mq *messageQueue) Create() (string, error) {
	q, err := mq.ch.QueueDeclare("", false, false, true, false, nil)
	return q.Name, err
}

// func (mq *messageQueue) Read() (string, error) {
// 	mq.ch.Consume()
// }

func (mq *messageQueue) Write(queueName string, json []byte) error {
	go func() {
		c := make(chan amqp.Return)
		for r := range mq.ch.NotifyReturn(c) {
			log.Println("Return reply-code:", r.ReplyCode)
		}
	}()

	msg := amqp.Publishing{
		ContentType: "application/json",
		Body:        json,
	}
	return mq.ch.Publish("", queueName, true, true, msg)
}
