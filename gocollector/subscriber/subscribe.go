package subscriber

import (
	"errors"
	"github.com/gesiel/go-collect/gocollector/utils"
)

var MissingClientIdError = errors.New("subscriber missing field: ClientId")
var MissingNameError = errors.New("subscriber missing field: Name")
var MissingEmailError = errors.New("subscriber missing field: Email")

type SubscribeUseCase struct {
	Gateway Gateway
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
	if !utils.IsValidValue(input.GetClientId()) {
		return MissingClientIdError
	}
	if !utils.IsValidValue(input.GetName()) {
		return MissingNameError
	}
	if !utils.IsValidValue(input.GetEmail()) {
		return MissingEmailError
	}
	return nil
}

func createSubscriberFor(input SubscribeInput) *Subscriber {
	return &Subscriber{
		ClientId: input.GetClientId(),
		Name:     input.GetName(),
		Email:    input.GetEmail(),
	}
}

type SubscribeInput interface {
	GetClientId() string
	GetName() string
	GetEmail() string
}

type SubscribeResponse struct {
	Subscriber *Subscriber
}
