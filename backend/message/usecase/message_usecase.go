package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/eyalch/pipeit/backend/domain"
)

type messageUsecase struct {
	queue domain.MessageQueue
}

func NewMessageUsecase(q domain.MessageQueue) domain.MessageUsecase {
	return &messageUsecase{q}
}

func (uc *messageUsecase) CreateQueue() (string, error) {
	return uc.queue.Create()
}

func (uc *messageUsecase) SendMessage(queueName string, message domain.Message) error {
	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}
	fmt.Println(string(payload))

	return uc.queue.Write(queueName, payload)
}
