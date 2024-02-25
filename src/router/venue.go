package router

import (
	"github.com/fucso/locos-only-api/src/controller"
	"github.com/labstack/echo/v4"
)

func InitVenueRouter(e *echo.Echo, i initializer[controller.VenueController]) {
	m := map[string]executer[controller.VenueController]{
		"Create": func(c controller.VenueController) echo.HandlerFunc { return c.Create() },
	}
	h := handler[controller.VenueController]{m: m}

	e.POST("/venue", h.handle(i, "Create"))
}
