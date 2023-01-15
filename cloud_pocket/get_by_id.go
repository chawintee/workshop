package pocket

import (
	"errors"
	"net/http"

	"github.com/kkgo-software-engineering/workshop/mlog"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

const (
	getDetailStmt = "SELECT id, name, balance, currency, category FROM cloud_pockets WHERE id = $1;"
)

func (h handler) GetById(c echo.Context) error {
	logger := mlog.L(c)
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, errors.New("bad"))
	}

	stmt, err := h.db.Prepare(getDetailStmt)
	if err != nil {
		logger.Error("query prepare error", zap.Error(err))
	}

	rows := stmt.QueryRow(id)

	var p CloudPocketsResponse
	err = rows.Scan(&p.ID, &p.Name, &p.Balance, &p.Currency, &p.Category)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, p)
}
