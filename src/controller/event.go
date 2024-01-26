package controller

import (
	"net/http"

	"github.com/fucso/locos-only-api/src/usecase"
	"github.com/labstack/echo/v4"
)

type EventController struct {
	usecase *usecase.EventUsecase
}

func NewEventController(usecase *usecase.EventUsecase) *EventController {
	return &EventController{
		usecase: usecase,
	}
}

func (con *EventController) FindAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		events, err := con.usecase.FindAll()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, events)
	}
}
