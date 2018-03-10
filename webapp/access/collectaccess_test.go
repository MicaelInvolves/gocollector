package access_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"errors"
	"github.com/gesiel/go-collect/webapp/access"
	"time"
)

var _ = Describe("Collect Access Use Case", func() {
	var (
		useCase    *access.CollectAccessUseCase
		input      *CollectAccessInputMock
		gateway    *AccessGatewayMock
		mockedDate time.Time
	)

	Context("Collecting a new access", func() {

		BeforeEach(func() {
			gateway = &AccessGatewayMock{}
			useCase = &access.CollectAccessUseCase{
				Gateway: gateway,
			}
			mockedDate = time.Date(2018, time.March, 10, 14, 14, 0, 0, time.Local)
		})

		It("Should save a new valid access", func() {
			input = NewCollectAccessInputMock("id", "path/to/resource", mockedDate)

			response, err := useCase.Collect(input)

			Expect(response).To(Not(BeNil()))
			Expect(gateway.SaveCount).To(Equal(1))

			Expect(response.Access).To(Not(BeNil()))
			Expect(response.Access).To(Equal(gateway.SavedAccess))
			Expect(response.Access.ClientId).To(Equal(input.clientId))
			Expect(response.Access.Path).To(Equal(input.path))
			Expect(response.Access.Date).To(Equal(input.date))

			Expect(err).To(BeNil())
		})

		It("Should propagate gateway error on save fail", func() {
			input = NewCollectAccessInputMock("id", "path/to/resource", mockedDate)

			saveErr := errors.New("BOOM!")
			gateway.Err = saveErr

			response, err := useCase.Collect(input)

			Expect(response).To(BeNil())
			Expect(err).To(Equal(saveErr))
		})

		It("Should validate missing ClientId", func() {
			input = NewCollectAccessInputMock("", "path/to/resource", mockedDate)
			response, err := useCase.Collect(input)
			assertErrorResponse(response, err, access.MissingClientIdError, "Access missing field: ClientId")

			input = NewCollectAccessInputMock("   ", "path/to/resource", mockedDate)
			response, err = useCase.Collect(input)
			assertErrorResponse(response, err, access.MissingClientIdError, "Access missing field: ClientId")
		})

		It("Should validate missing Path", func() {
			input = NewCollectAccessInputMock("ClientId", "", mockedDate)
			response, err := useCase.Collect(input)
			assertErrorResponse(response, err, access.MissingPathError, "Access missing field: Path")

			input = NewCollectAccessInputMock("ClientId", "  ", mockedDate)
			response, err = useCase.Collect(input)
			assertErrorResponse(response, err, access.MissingPathError, "Access missing field: Path")
		})
	})
})

func assertErrorResponse(response *access.CollectAccessResponse, err, expectedError error, expectedMessage string) {
	Expect(response).To(BeNil())
	Expect(err).To(Equal(expectedError))
	Expect(err.Error()).To(Equal(expectedMessage))
}

/* =============== MOCKS =============== */

/* ======== INPUT ======== */

type CollectAccessInputMock struct {
	clientId string
	path     string
	date     time.Time
}

func (this *CollectAccessInputMock) ClientId() string {
	return this.clientId
}

func (this *CollectAccessInputMock) Path() string {
	return this.path
}

func (this *CollectAccessInputMock) Date() time.Time {
	return this.date
}

func NewCollectAccessInputMock(id, path string, date time.Time) *CollectAccessInputMock {
	return &CollectAccessInputMock{
		clientId: id,
		path:     path,
		date:     date,
	}
}

/* ======== GATEWAY ======== */

type AccessGatewayMock struct {
	SavedAccess *access.Access
	SaveCount   int
	Err         error
}

func (this *AccessGatewayMock) Save(access *access.Access) error {
	if this.Err != nil {
		return this.Err
	}
	this.SaveCount++
	this.SavedAccess = access
	return nil
}
