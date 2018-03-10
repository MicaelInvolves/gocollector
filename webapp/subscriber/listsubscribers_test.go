package subscriber_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/gesiel/go-collect/webapp/subscriber"
)

var _ = Describe("List subscribers", func() {
	var (
		useCase *subscriber.ListSubscribersAccessDataUseCase
		gateway *SubscriberGatewayMock
	)

	BeforeEach(func() {
		gateway = &SubscriberGatewayMock{}
		useCase = &subscriber.ListSubscribersAccessDataUseCase{
			Gateway: gateway,
		}
	})

	It("No subscribers found", func() {
		gateway.AllData = []*subscriber.SubscribersAccessData{}

		response, err := useCase.List()

		Expect(gateway.AllCount).Should(Equal(1))
		Expect(response).ShouldNot(BeNil())
		Expect(len(response.Subscribers)).Should(Equal(0))
		Expect(response.Subscribers).Should(Equal(gateway.AllData))
		Expect(err).Should(BeNil())
	})
})
