package access

import (
	"errors"
	"strings"
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
	this.Gateway.Save(access)

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
	if !isValidString(input.ClientId()) {
		return MissingClientIdError
	}

	if !isValidString(input.Path()) {
		return MissingPathError
	}

	return nil
}

func isValidString(value string) bool {
	trim := strings.Trim(value, " ")
	return len(trim) > 0
}

func createAccessFor(input CollectAccessInput) *Access {
	access := &Access{
		ClientId: input.ClientId(),
		Path:     input.Path(),
		Date:     input.Date(),
	}
	return access
}
