package access

import "time"

type CollectAccessUseCase struct {
	Gateway AccessGateway
}

func (this *CollectAccessUseCase) Collect(input CollectAccessInput) (response *CollectAccessResponse, err error) {
	access := createAccessFor(input)

	this.Gateway.Save(access)

	return &CollectAccessResponse{
		Access: access,
	}, nil
}

func createAccessFor(input CollectAccessInput) *Access {
	access := &Access{
		ClientId: input.ClientId(),
		Path:     input.Path(),
		Date:     input.Date(),
	}
	return access
}

type CollectAccessResponse struct {
	Access *Access
}

type CollectAccessInput interface {
	ClientId() string
	Path() string
	Date() time.Time
}
