package access_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/gesiel/go-collect/webapp/access"
	"time"
)

var _ = Describe("Access Use Case", func() {
	var (
		useCase    *access.CollectAccessUseCase
		input      *CollectAccessInputMock
		gateway    *AccessGatewayMock
		mockedDate time.Time
	)

	BeforeEach(func() {
		gateway = &AccessGatewayMock{}
		useCase = &access.CollectAccessUseCase{
			Gateway: gateway,
		}
		mockedDate = time.Date(2018, time.March, 10, 14, 14, 0, 0, time.Local)
	})

	Context("Collecting a new access", func() {
		It("Should save a new valid access", func() {
			input = NewCollectAccessInputMock("id", "path/to/resource", mockedDate)

			response, err := useCase.Collect(input)

			Expect(response).To(Not(BeNil()))
			Expect(response.Access).To(Not(BeNil()))
			Expect(response.Access.Id).To(Equal("dbId"))
			Expect(response.Access.ClientId).To(Equal("id"))
			Expect(response.Access.Path).To(Equal("path/to/resource"))
			Expect(response.Access.Date).To(Equal(mockedDate))
			Expect(err).To(BeNil())
		})
	})
})

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
}

func (*AccessGatewayMock) Save(access *access.Access) error {
	access.Id = "dbId"
	return nil
}
