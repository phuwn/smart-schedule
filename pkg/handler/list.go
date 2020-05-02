package handler

import (
	"github.com/labstack/echo"
	"github.com/phuwn/smart-schedule/pkg/handler/list"
)

func listRoutes(r *echo.Echo) {
	g := r.Group("/list")
	{
		g.POST("", newList)
	}
}

func newList(c echo.Context) error {
	l, err := list.ParseCreateListRequest(c)
	if err != nil {
		return err
	}
	return list.CreateList(c, l)
}
