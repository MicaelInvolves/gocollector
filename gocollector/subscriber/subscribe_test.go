package subscriber

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"errors"
)

var _ = Describe("Subscribe Use Case", func() {
	var (
		useCase *SubscribeUseCase
		input   *SubscribeInputMock
		gateway *SubscriberGatewayMock
	)

	BeforeEach(func() {
		gateway = &SubscriberGatewayMock{}
		useCase = &SubscribeUseCase{
			Gateway: gateway,
		}
	})

	Context("Subscribing", func() {

		It("Should save a valid subscriber", func() {
			input = NewSubscribeInputMock("id", "Jay", "valid@email.com")

			response, err := useCase.Subscribe(input)

			Expect(response).ShouldNot(BeNil())
			Expect(gateway.SaveCount).Should(Equal(1))

			Expect(response.Subscriber).ShouldNot(BeNil())
			Expect(response.Subscriber).Should(Equal(gateway.SavedSubscriber))
			Expect(response.Subscriber.ClientId).Should(Equal(input.clientId))
			Expect(response.Subscriber.Name).Should(Equal(input.name))
			Expect(response.Subscriber.Email).Should(Equal(input.email))

			Expect(err).Should(BeNil())
		})

		It("Should propagate gateway error on save fail", func() {
			input = NewSubscribeInputMock("id", "Jay", "valid@email.com")
			saveErr := errors.New("BOOM!")
			gateway.Err = saveErr

			response, err := useCase.Subscribe(input)

			Expect(response).Should(BeNil())
			Expect(err).Should(Equal(saveErr))
		})

		It("Should validate missing ClientId", func() {
			input = NewSubscribeInputMock("", "Jay", "valid@email.com")
			response, err := useCase.Subscribe(input)
			assertResponseError(response, err, MissingClientIdError, "subscriber missing field: ClientId")

			input = NewSubscribeInputMock("  ", "Jay", "valid@email.com")
			response, err = useCase.Subscribe(input)
			assertResponseError(response, err, MissingClientIdError, "subscriber missing field: ClientId")
		})

		It("Should validate missing Name", func() {
			input = NewSubscribeInputMock("id", "", "valid@email.com")
			response, err := useCase.Subscribe(input)
			assertResponseError(response, err, MissingNameError, "subscriber missing field: Name")

			input = NewSubscribeInputMock("id", " ", "valid@email.com")
			response, err = useCase.Subscribe(input)
			assertResponseError(response, err, MissingNameError, "subscriber missing field: Name")
		})

		It("Should validate missing Email", func() {
			input = NewSubscribeInputMock("id", "Jay", "")
			response, err := useCase.Subscribe(input)
			assertResponseError(response, err, MissingEmailError, "subscriber missing field: Email")

			input = NewSubscribeInputMock("id", "Jay", "")
			response, err = useCase.Subscribe(input)
			assertResponseError(response, err, MissingEmailError, "subscriber missing field: Email")
		})
	})
})

func assertResponseError(response *SubscribeResponse, err, expectedErr error, expectedMessage string) {
	Expect(response).Should(BeNil())
	Expect(err).To(Equal(expectedErr))
	Expect(err.Error()).To(Equal(expectedMessage))
}

/* =============== MOCKS =============== */

/* ======== INPUT ======== */

type SubscribeInputMock struct {
	clientId string
	name     string
	email    string
}

func (this *SubscribeInputMock) GetClientId() string {
	return this.clientId
}

func (this *SubscribeInputMock) GetName() string {
	return this.name
}

func (this *SubscribeInputMock) GetEmail() string {
	return this.email
}

func NewSubscribeInputMock(clientId, name, email string) *SubscribeInputMock {
	return &SubscribeInputMock{
		clientId: clientId,
		name:     name,
		email:    email,
	}
}

/* ======== GATEWAY ======== */

type SubscriberGatewayMock struct {
	SavedSubscriber *Subscriber
	SaveCount       int

	AllCount int
	AllData  []*SubscribersAccessData

	Err error
}

func (this *SubscriberGatewayMock) Save(subscriber *Subscriber) error {
	this.SaveCount++
	if this.Err != nil {
		return this.Err
	}
	this.SavedSubscriber = subscriber
	return nil
}

func (this *SubscriberGatewayMock) All() ([]*SubscribersAccessData, error) {
	this.AllCount++
	if this.Err != nil {
		return nil, this.Err
	}
	return this.AllData, nil
}
