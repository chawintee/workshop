package pocket

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	getDetailStmt = "SELECT * FROM cloud_pockets WHERE id = $1;"
)

func (h handler) GetById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, errors.New("bad"))
	}

	return c.JSON(http.StatusOK, nil)
}
