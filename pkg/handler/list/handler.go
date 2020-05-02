package list

import (
	"github.com/labstack/echo"

	"github.com/phuwn/smart-schedule/pkg/model"
	"github.com/phuwn/smart-schedule/pkg/server"
	"github.com/phuwn/tools/errors"
)

// CreateList - create new list dta handler
func CreateList(c echo.Context, list *model.List) error {
	cfg := server.GetServerCfg()
	return errors.Customize(0, "failed to create new list", cfg.Store().List.Create(c, list))
}
