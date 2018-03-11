package controllers

import (
	"github.com/gesiel/go-collect/gocollector/subscriber"
	"github.com/labstack/echo"
	"net/http"
)

type SubscribeController struct {
	UseCase *subscriber.SubscribeUseCase
}

func (this *SubscribeController) Subscribe(context echo.Context) error {
	viewModel := &SubscriberViewModel{}

	if err := context.Bind(viewModel); err != nil {
		return err
	}

	response, err := this.UseCase.Subscribe(viewModel)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return context.JSON(http.StatusCreated, &SubscriberViewModel{
		ClientId: response.Subscriber.ClientId,
		Name:     response.Subscriber.Name,
		Email:    response.Subscriber.Email,
	})
}

type SubscriberViewModel struct {
	ClientId string `json:"clientId" form:"clientId"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
}

func (this *SubscriberViewModel) GetClientId() string {
	return this.ClientId
}

func (this *SubscriberViewModel) GetName() string {
	return this.Name
}

func (this *SubscriberViewModel) GetEmail() string {
	return this.Email
}
