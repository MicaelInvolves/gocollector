package access

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"errors"
	"time"
)

var _ = Describe("Collect Access Use Case", func() {
	var (
		useCase    *CollectAccessUseCase
		input      *CollectAccessInputMock
		gateway    *AccessGatewayMock
		mockedDate time.Time
	)

	Context("Collecting a new access", func() {

		BeforeEach(func() {
			gateway = &AccessGatewayMock{}
			useCase = &CollectAccessUseCase{
				Gateway: gateway,
			}
			mockedDate = time.Date(2018, time.March, 10, 14, 14, 0, 0, time.Local)
		})

		It("Should save a new valid access", func() {
			input = NewCollectAccessInputMock("id", "path/to/resource", mockedDate)

			response, err := useCase.Collect(input)

			Expect(response).ShouldNot(BeNil())
			Expect(gateway.SaveCount).Should(Equal(1))

			Expect(response.Access).ShouldNot(BeNil())
			Expect(response.Access).Should(Equal(gateway.SavedAccess))
			Expect(response.Access.ClientId).Should(Equal(input.clientId))
			Expect(response.Access.Path).Should(Equal(input.path))
			Expect(response.Access.Date).Should(Equal(input.date))

			Expect(err).Should(BeNil())
		})

		It("Should propagate gateway error on save fail", func() {
			input = NewCollectAccessInputMock("id", "path/to/resource", mockedDate)

			saveErr := errors.New("BOOM!")
			gateway.Err = saveErr

			response, err := useCase.Collect(input)

			Expect(response).Should(BeNil())
			Expect(err).Should(Equal(saveErr))
		})

		It("Should validate missing ClientId", func() {
			input = NewCollectAccessInputMock("", "path/to/resource", mockedDate)
			response, err := useCase.Collect(input)
			assertErrorResponse(response, err, MissingClientIdError, "Access missing field: ClientId")

			input = NewCollectAccessInputMock("   ", "path/to/resource", mockedDate)
			response, err = useCase.Collect(input)
			assertErrorResponse(response, err, MissingClientIdError, "Access missing field: ClientId")
		})

		It("Should validate missing Path", func() {
			input = NewCollectAccessInputMock("ClientId", "", mockedDate)
			response, err := useCase.Collect(input)
			assertErrorResponse(response, err, MissingPathError, "Access missing field: Path")

			input = NewCollectAccessInputMock("ClientId", "  ", mockedDate)
			response, err = useCase.Collect(input)
			assertErrorResponse(response, err, MissingPathError, "Access missing field: Path")
		})
	})
})

func assertErrorResponse(response *CollectAccessResponse, err, expectedError error, expectedMessage string) {
	Expect(response).Should(BeNil())
	Expect(err).Should(Equal(expectedError))
	Expect(err.Error()).Should(Equal(expectedMessage))
}

/* =============== MOCKS =============== */

/* ======== INPUT ======== */

type CollectAccessInputMock struct {
	clientId string
	path     string
	date     time.Time
}

func (this *CollectAccessInputMock) GetClientId() string {
	return this.clientId
}

func (this *CollectAccessInputMock) GetPath() string {
	return this.path
}

func (this *CollectAccessInputMock) GetDate() time.Time {
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
	SavedAccess *Access
	SaveCount   int
	Err         error
}

func (this *AccessGatewayMock) Save(access *Access) error {
	if this.Err != nil {
		return this.Err
	}
	this.SaveCount++
	this.SavedAccess = access
	return nil
}
