package main

import (
	"github.com/gesiel/go-collect/collector/access"
	"github.com/gesiel/go-collect/collector/controllers"
	"github.com/gesiel/go-collect/collector/database"
	"github.com/gesiel/go-collect/collector/subscriber"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"os"
	"github.com/labstack/echo/middleware"
)

const databaseHost = "DATABASE_URL"
const port = "PORT"

func main() {
	host := os.Getenv(databaseHost)
	port := os.Getenv(port)

	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collectAccessCtrl := &controllers.CollectAccessController{
		UseCase: &access.CollectAccessUseCase{
			Gateway: database.NewAccessGateway(session),
		},
	}
	subscribeCtrl := &controllers.SubscribeController{
		UseCase: &subscriber.SubscribeUseCase{
			Gateway: database.NewSubscriberGateway(session),
		},
	}
	listSubscribersCtrl := &controllers.ListSubscribersController{
		UseCase: &subscriber.ListSubscribersAccessDataUseCase{
			Gateway: database.NewSubscriberGateway(session),
		},
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.POST("/access", collectAccessCtrl.Collect)
	e.POST("/subscribe", subscribeCtrl.Subscribe)
	e.GET("/subscribers", listSubscribersCtrl.List)

	e.Start(":" + port)
}
