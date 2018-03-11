package controllers

import (
	"github.com/gesiel/go-collect/gocollector/access"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type CollectAccessController struct {
	UseCase *access.CollectAccessUseCase
}

func (this *CollectAccessController) Collect(context echo.Context) error {
	viewModel := new(AccessViewModel)

	if err := context.Bind(viewModel); err != nil {
		return err
	}

	response, err := this.UseCase.Collect(viewModel)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return context.JSON(http.StatusCreated, &AccessViewModel{
		Id:       response.Access.Id,
		ClientId: response.Access.ClientId,
		Path:     response.Access.Path,
		Date:     response.Access.Date,
	})
}

type AccessViewModel struct {
	Id       string    `json:"id" form:"id"`
	ClientId string    `json:"clientId" form:"clientId"`
	Path     string    `json:"path" form:"path"`
	Date     time.Time `json:"date" form:"date"`
}

func (this *AccessViewModel) GetClientId() string {
	return this.ClientId
}

func (this *AccessViewModel) GetPath() string {
	return this.Path
}

func (this *AccessViewModel) GetDate() time.Time {
	return this.Date
}
