package list

import (
	"github.com/labstack/echo"

	"github.com/phuwn/smart-schedule/pkg/model"
	"github.com/phuwn/tools/db"
)

type listPGStore struct{}

// NewStore - create new list store
func NewStore() Store {
	return &listPGStore{}
}

func (s listPGStore) Find(c echo.Context) ([]*model.List, error) {
	return nil, nil
}

func (s listPGStore) Get(c echo.Context, id string) (*model.List, error) {
	tx := db.GetTxFromCtx(c)
	var res model.List
	return &res, tx.Where("id = ?", id).First(&res).Error
}

func (s listPGStore) Create(c echo.Context, list *model.List) error {
	tx := db.GetTxFromCtx(c)
	return tx.Create(list).Error
}
