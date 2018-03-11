package subscriber

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"errors"
)

var _ = Describe("List subscribers", func() {
	var (
		useCase *ListSubscribersAccessDataUseCase
		gateway *SubscriberGatewayMock
	)

	BeforeEach(func() {
		gateway = &SubscriberGatewayMock{}
		useCase = &ListSubscribersAccessDataUseCase{
			Gateway: gateway,
		}
	})

	It("Should ask for subscribers to gateway", func() {
		gateway.AllData = []*SubscribersAccessData{}

		response, err := useCase.List()

		Expect(gateway.AllCount).Should(Equal(1))
		Expect(response).ShouldNot(BeNil())
		Expect(len(response.SubscribersAccessData)).Should(Equal(0))
		Expect(response.SubscribersAccessData).Should(Equal(gateway.AllData))

		Expect(err).Should(BeNil())
	})

	It("Should propagate gateway error", func() {
		gateway.Err = errors.New("BOOM!")

		response, err := useCase.List()

		Expect(gateway.AllCount).Should(Equal(1))
		Expect(response).Should(BeNil())
		Expect(err).ShouldNot(BeNil())
	})
})
