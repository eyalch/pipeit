package usecase

import (
	"log"

	"github.com/eyalch/pipeit/backend/domain"
)

type codeUsecase struct {
	repo   domain.CodeRepository
	pubsub domain.CodePubSub

	message domain.MessageUsecase
}

// NewCodeUsecase creates a new code usecase
func NewCodeUsecase(r domain.CodeRepository, ps domain.CodePubSub, m domain.MessageUsecase) domain.CodeUsecase {
	return &codeUsecase{r, ps, m}
}

func (uc *codeUsecase) SendRandomCodeAndWaitForPair(cw domain.CodeWriter) error {
	code, err := uc.repo.GetRandomCode()
	if err != nil {
		return err
	}

	// Send the code to the client
	err = cw.Write(code)
	if err != nil {
		return err
	}

	// Subscribe to a channel to be able to respond to pair events
	for range uc.pubsub.Subscribe(code) {
		log.Println("Someone has paired with code", code)
	}

	return nil
}

func (uc *codeUsecase) Pair(code string) error {
	return nil
}
