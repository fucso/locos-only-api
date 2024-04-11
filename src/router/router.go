package router

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type initializer[C any] func() *C

type executer[C any] func(c C) echo.HandlerFunc
type mapper[C any] map[string]executer[C]
type handler[C any] struct {
	m mapper[C]
}

func (h handler[C]) handle(i initializer[C], method string) echo.HandlerFunc {
	return func(c echo.Context) error {
		controller := i()
		executer := h.m[method]
		fmt.Println("executer", executer)
		if executer == nil {
			return c.String(404, "Not Found")
		}

		return executer(*controller)(c)
	}
}
