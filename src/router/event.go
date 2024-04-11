package router

import (
	"github.com/fucso/locos-only-api/src/controller"
	"github.com/labstack/echo/v4"
)

func InitEventRouter(e *echo.Echo, i initializer[controller.EventController]) {
	m := map[string]executer[controller.EventController]{
		"FindAll": func(c controller.EventController) echo.HandlerFunc { return c.FindAll() },
		"Create":  func(c controller.EventController) echo.HandlerFunc { return c.Create() },
	}
	h := handler[controller.EventController]{m: m}

	e.GET("/events", h.handle(i, "FindAll"))
	e.POST("/event", h.handle(i, "Create"))
}
