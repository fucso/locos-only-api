package router

import (
	"github.com/fucso/locos-only-api/src/controller"
	"github.com/labstack/echo/v4"
)

type initializer func() *controller.EventController

type Method string

const (
	FindAll Method = "FindAll"
	Create  Method = "Create"
)

func InitEventRouter(e *echo.Echo, i initializer) {
	e.GET("/events", handle(i, FindAll))
	e.POST("/event", handle(i, Create))
}

func handle(i initializer, method Method) echo.HandlerFunc {
	return func(c echo.Context) error {
		controller := i()
		handler := func() echo.HandlerFunc {
			switch method {
			case FindAll:
				return controller.FindAll()
			case Create:
				return controller.Create()
			default:
				return nil
			}
		}()
		if handler == nil {
			return c.String(404, "Not Found")
		}

		return handler(c)
	}
}
