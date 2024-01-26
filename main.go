package main

import (
	"net/http"

	"github.com/fucso/locos-only-api/src/controller"
	"github.com/fucso/locos-only-api/src/infrastructure"
	"github.com/fucso/locos-only-api/src/router"
	"github.com/fucso/locos-only-api/src/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db, err := infrastructure.NewDatabase()
	if err != nil {
		e.Logger.Fatal(err)
	}

	router.InitEventRouter(e, func() *controller.EventController {
		usecase := usecase.NewEventUsecase(db)
		return controller.NewEventController(usecase)
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
