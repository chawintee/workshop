package cloud_pocket

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h handler) GetById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, errors.New("bad"))
	}

	return c.JSON(http.StatusOK, nil)
}
