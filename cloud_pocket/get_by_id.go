package cloud_pocket

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h handler) GetById(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
