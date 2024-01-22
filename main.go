package main

import (
	"net/http"

	"github.com/fucso/locos-only-api/src/infrastructure"
	"github.com/fucso/locos-only-api/src/repository"
	"github.com/fucso/locos-only-api/src/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db, err := infrastructure.NewDatabase()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})

	e.GET("/events", func(c echo.Context) error {
		repo := repository.NewEventRepository(db)
		usecase := usecase.NewEventUsecase(repo)
		events, err := usecase.FindAll()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, events)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
