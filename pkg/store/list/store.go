package list

import (
	"github.com/labstack/echo"

	"github.com/phuwn/smart-schedule/pkg/model"
)

// Store - list store interface
type Store interface {
	Find(c echo.Context) ([]*model.List, error)
	Get(c echo.Context, id string) (*model.List, error)
	Create(c echo.Context, list *model.List) error
}
