package controllers

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"encoding/json"
	"github.com/gesiel/go-collect/collector/subscriber"
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("List subscribers api", func() {
	Context("List subscribers", func() {
		var (
			controller *ListSubscribersController
			recorder   *httptest.ResponseRecorder
			gateway    *SubscriberGatewayInMemory
		)

		BeforeEach(func() {
			gateway = &SubscriberGatewayInMemory{
				Db: map[string]*subscriber.Subscriber{},
			}
			controller = &ListSubscribersController{
				UseCase: &subscriber.ListSubscribersAccessDataUseCase{
					Gateway: gateway,
				},
			}
			recorder = httptest.NewRecorder()
		})

		It("No subscribers found", func() {
			gateway.AllData = []*subscriber.SubscribersAccessData{}

			request := httptest.NewRequest(echo.GET, "/subscribers", nil)
			context := echo.New().NewContext(request, recorder)

			err := controller.List(context)

			Expect(err).Should(BeNil())
			Expect(recorder.Code).Should(Equal(http.StatusOK))
			Expect(recorder.Body.String()).Should(Equal("[]"))
		})

		It("No subscribers found", func() {
			gateway.AllData = []*subscriber.SubscribersAccessData{
				{
					Subscriber: &subscriber.Subscriber{
						Name:     "Jay",
						ClientId: "id1",
						Email:    "valid1@email.com",
					},
					AccessCount: 10,
					AccessPaths: []string{"home", "contact"},
				},
				{
					Subscriber: &subscriber.Subscriber{
						Name:     "John",
						ClientId: "id2",
						Email:    "valid2@email.com",
					},
					AccessCount: 1,
					AccessPaths: []string{"contact"},
				},
			}

			request := httptest.NewRequest(echo.GET, "/subscribers", nil)
			context := echo.New().NewContext(request, recorder)

			err := controller.List(context)

			Expect(err).Should(BeNil())
			Expect(recorder.Code).Should(Equal(http.StatusOK))

			results := sliceFromJSON(recorder.Body.Bytes())

			result := results[0].(map[string]interface{})
			Expect(result["clientId"]).Should(Equal("id1"))
			Expect(result["name"]).Should(Equal("Jay"))
			Expect(result["email"]).Should(Equal("valid1@email.com"))
			Expect(result["visits"]).Should(Equal(float64(10)))
			Expect(result["pages"]).Should(Equal("home, contact"))

			result = results[1].(map[string]interface{})
			Expect(result["clientId"]).Should(Equal("id2"))
			Expect(result["name"]).Should(Equal("John"))
			Expect(result["email"]).Should(Equal("valid2@email.com"))
			Expect(result["visits"]).Should(Equal(float64(1)))
			Expect(result["pages"]).Should(Equal("contact"))
		})
	})
})

func sliceFromJSON(data []byte) []interface{} {
	var result interface{}
	json.Unmarshal(data, &result)
	return result.([]interface{})
}
