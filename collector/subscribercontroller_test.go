package collector

import (
	"github.com/gesiel/go-collect/collector/subscriber"
	"github.com/labstack/echo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
)

const validSubscriberJson = `{ "clientId": "id", "name": "Jay", "email": "valid@email.com" }`
const invalidSubscriberJson = `{ "clientId": "id", "name": "Jay" }`

var _ = Describe("Subscribe api", func() {
	Context("Subscribe", func() {
		var controller *SubscribeController
		var recorder *httptest.ResponseRecorder
		var gateway *SubscriberGatewayInMemory

		BeforeEach(func() {
			gateway = &SubscriberGatewayInMemory{
				Db: map[string]*subscriber.Subscriber{},
			}
			controller = &SubscribeController{
				UseCase: &subscriber.SubscribeUseCase{
					Gateway: gateway,
				},
			}
			recorder = httptest.NewRecorder()
		})

		It("A valid subscriber data", func() {
			request := httptest.NewRequest(echo.POST, "/subscribe", strings.NewReader(validSubscriberJson))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
			context := echo.New().NewContext(request, recorder)

			err := controller.subscribe(context)

			Expect(err).Should(BeNil())
			Expect(recorder.Code).Should(Equal(http.StatusCreated))
			Expect(gateway.IdCount).Should(Equal(1))

			result := mapFromJSON(recorder.Body.Bytes())

			Expect(result["clientId"]).Should(Equal("id"))
			Expect(result["name"]).Should(Equal("Jay"))
			Expect(result["email"]).Should(Equal("valid@email.com"))
		})

		It("A invalid subscriber data", func() {
			request := httptest.NewRequest(echo.POST, "/subscribe", strings.NewReader(invalidSubscriberJson))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
			context := echo.New().NewContext(request, recorder)

			err := controller.subscribe(context)

			Expect(err).Should(BeNil())
			Expect(recorder.Code).Should(Equal(http.StatusBadRequest))

			result := mapFromJSON(recorder.Body.Bytes())
			Expect(result["error"]).Should(Equal("subscriber missing field: Email"))
		})

	})
})

/* =========== MOCKS =========== */

/* =========== GATEWAY ===========*/

type SubscriberGatewayInMemory struct {
	IdCount int
	Db      map[string]*subscriber.Subscriber
	AllData []*subscriber.SubscribersAccessData
}

func (this *SubscriberGatewayInMemory) Save(subscriber *subscriber.Subscriber) error {
	this.IdCount++
	subscriberId := strconv.Itoa(this.IdCount)
	this.Db[subscriberId] = subscriber
	return nil
}

func (this *SubscriberGatewayInMemory) Get(id int) *subscriber.Subscriber {
	return this.Db[strconv.Itoa(this.IdCount)]
}

func (this *SubscriberGatewayInMemory) All() ([]*subscriber.SubscribersAccessData, error) {
	return this.AllData, nil
}
