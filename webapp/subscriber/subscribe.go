package subscriber

import (
	"errors"
	"github.com/gesiel/go-collect/webapp/utils"
)

var MissingClientIdError = errors.New("subscriber missing field: ClientId")
var MissingNameError = errors.New("subscriber missing field: Name")
var MissingEmailError = errors.New("subscriber missing field: Email")

type SubscribeUseCase struct {
	Gateway SubscriberGateway
}

func (this *SubscribeUseCase) Subscribe(input SubscribeInput) (*SubscribeResponse, error) {
	err := validateInput(input)
	if err != nil {
		return nil, err
	}

	subscriber := createSubscriberFor(input)
	err = this.Gateway.Save(subscriber)
	if err != nil {
		return nil, err
	}

	return &SubscribeResponse{
		Subscriber: subscriber,
	}, nil
}

func validateInput(input SubscribeInput) error {
	if !utils.IsValidValue(input.ClientId()) {
		return MissingClientIdError
	}
	if !utils.IsValidValue(input.Name()) {
		return MissingNameError
	}
	if !utils.IsValidValue(input.Email()) {
		return MissingEmailError
	}
	return nil
}

func createSubscriberFor(input SubscribeInput) *Subscriber {
	return &Subscriber{
		ClientId: input.ClientId(),
		Name:     input.Name(),
		Email:    input.Email(),
	}
}

type SubscribeInput interface {
	ClientId() string
	Name() string
	Email() string
}

type SubscribeResponse struct {
	Subscriber *Subscriber
}
