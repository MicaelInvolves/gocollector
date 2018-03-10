package access

import (
	"errors"
	"github.com/gesiel/go-collect/webapp/utils"
	"time"
)

var MissingClientIdError = errors.New("Access missing field: ClientId")
var MissingPathError = errors.New("Access missing field: Path")

type CollectAccessUseCase struct {
	Gateway AccessGateway
}

func (this *CollectAccessUseCase) Collect(input CollectAccessInput) (*CollectAccessResponse, error) {
	err := validateInput(input)
	if err != nil {
		return nil, err
	}

	access := createAccessFor(input)
	err = this.Gateway.Save(access)
	if err != nil {
		return nil, err
	}

	return &CollectAccessResponse{
		Access: access,
	}, nil
}

type CollectAccessResponse struct {
	Access *Access
}

type CollectAccessInput interface {
	ClientId() string
	Path() string
	Date() time.Time
}

func validateInput(input CollectAccessInput) error {
	if !utils.IsValidValue(input.ClientId()) {
		return MissingClientIdError
	}

	if !utils.IsValidValue(input.Path()) {
		return MissingPathError
	}

	return nil
}

func createAccessFor(input CollectAccessInput) *Access {
	access := &Access{
		ClientId: input.ClientId(),
		Path:     input.Path(),
		Date:     input.Date(),
	}
	return access
}
