package collector

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"encoding/json"
	"github.com/gesiel/go-collect/collector/access"
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
)

const validAccessJson = `{ "clientId": "id", "path": "path/to/some/resource", "date" : "2018-03-11T00:00:00Z" }`
const invalidAccessJson = `{ "path": "path/to/some/resource" }`
var _ = Describe("Collector api", func() {
	Context("Collect access", func() {
		var controller *CollectAccessController
		var recorder *httptest.ResponseRecorder
		var gateway *AccessGatewayInMemory

		BeforeEach(func() {
			gateway = &AccessGatewayInMemory{
				Db: map[string]*access.Access{},
			}
			controller = &CollectAccessController{
				UseCase: &access.CollectAccessUseCase{
					Gateway: gateway,
				},
			}
			recorder = httptest.NewRecorder()
		})

		It("A valid access data", func() {
			request := httptest.NewRequest(echo.POST, "/access", strings.NewReader(validAccessJson))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
			context := echo.New().NewContext(request, recorder)

			err := controller.collect(context)

			Expect(err).Should(BeNil())
			Expect(recorder.Code).Should(Equal(http.StatusCreated))
			Expect(gateway.IdCount).Should(Equal(1))

			result := mapFromJSON(recorder.Body.Bytes())

			Expect(result["id"]).Should(Equal("1"))
			Expect(result["clientId"]).Should(Equal("id"))
			Expect(result["path"]).Should(Equal("path/to/some/resource"))
			Expect(result["date"]).Should(Equal("2018-03-11T00:00:00Z"))
		})

		It("A invalid access data", func() {
			request := httptest.NewRequest(echo.POST, "/access", strings.NewReader(invalidAccessJson))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
			context := echo.New().NewContext(request, recorder)

			err := controller.collect(context)

			Expect(err).Should(BeNil())
			Expect(recorder.Code).Should(Equal(http.StatusBadRequest))

			result := mapFromJSON(recorder.Body.Bytes())
			Expect(result["error"]).Should(Equal("Access missing field: ClientId"))
		})
	})
})

func mapFromJSON(data []byte) map[string]interface{} {
	var result interface{}
	json.Unmarshal(data, &result)
	return result.(map[string]interface{})
}

/* =========== MOCKS =========== */

/* =========== GATEWAY ===========*/

type AccessGatewayInMemory struct {
	IdCount int
	Db      map[string]*access.Access
}

func (this *AccessGatewayInMemory) Save(access *access.Access) error {
	this.IdCount++
	accessId := strconv.Itoa(this.IdCount)
	this.Db[accessId] = access
	access.Id = accessId
	return nil
}

func (this *AccessGatewayInMemory) Get(id int) *access.Access {
	return this.Db[strconv.Itoa(this.IdCount)]
}
