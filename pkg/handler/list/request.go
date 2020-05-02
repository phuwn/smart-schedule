package list

import (
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo"

	"github.com/phuwn/smart-schedule/pkg/model"
	"github.com/phuwn/tools/errors"
)

// ParseCreateListRequest - get create list request info
func ParseCreateListRequest(c echo.Context) (*model.List, error) {
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return nil, errors.Customize(400, "failed to read request body", err)
	}

	var list model.List
	err = json.Unmarshal(b, &list)
	if err != nil {
		return nil, errors.Customize(400, "failed to parse request body into list model", err)
	}

	return &list, nil
}
