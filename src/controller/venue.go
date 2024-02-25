package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	params "github.com/fucso/locos-only-api/src/controller/params/venue"
	"github.com/fucso/locos-only-api/src/usecase"
)

type VenueController struct {
	usecase *usecase.VenueUsecase
}

func NewVenueController(usecase *usecase.VenueUsecase) *VenueController {
	return &VenueController{
		usecase: usecase,
	}
}

func (con *VenueController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var r params.CreateRequest
		if err := c.Bind(&r); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		venue, err := con.usecase.Create(r)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, venue)
	}
}
