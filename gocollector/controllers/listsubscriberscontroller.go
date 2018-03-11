package controllers

import (
	"github.com/gesiel/go-collect/gocollector/subscriber"
	"github.com/labstack/echo"
	"net/http"
)

type ListSubscribersController struct {
	UseCase *subscriber.ListSubscribersAccessDataUseCase
}

func (this *ListSubscribersController) List(context echo.Context) error {
	response, err := this.UseCase.List()
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, viewModel(response))
}

type SubscriberAccessDataViewModel struct {
	ClientId string `json:"clientId" form:"clientId"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Visits   int    `json:"visits" form:"visits"`
	Pages    string `json:"pages" form:"pages"`
}

func viewModel(response *subscriber.ListSubscribersAccessDataResponse) []*SubscriberAccessDataViewModel {
	result := make([]*SubscriberAccessDataViewModel, len(response.SubscribersAccessData))
	for i, accessData := range response.SubscribersAccessData {
		result[i] = viewModelFor(accessData)
	}
	return result
}

func viewModelFor(accessData *subscriber.SubscribersAccessData) *SubscriberAccessDataViewModel {
	return &SubscriberAccessDataViewModel{
		ClientId: accessData.Subscriber.ClientId,
		Name:     accessData.Subscriber.Name,
		Email:    accessData.Subscriber.Email,
		Visits:   accessData.AccessCount,
		Pages:    extractPages(accessData),
	}
}

func extractPages(accessData *subscriber.SubscribersAccessData) string {
	pages := accessData.AccessPaths[0]
	for i := 1; i < len(accessData.AccessPaths); i++ {
		pages += ", " + accessData.AccessPaths[i]
	}
	return pages
}
