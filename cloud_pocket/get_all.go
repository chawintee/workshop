package pocket

import (
	"net/http"

	"github.com/kkgo-software-engineering/workshop/mlog"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

const (
	cStmt = "INSERT INTO cloud_pockets (balance, name, category, currency) VALUES ($1, $2, $3, $4) RETURNING id;"
)

func (h handler) GetAll(c echo.Context) error {
	logger := mlog.L(c)
	stmt, err := h.db.Prepare("SELECT id, name, balance, category, currency  FROM cloud_pockets")
	if err != nil {
		logger.Error("query prepare error", zap.Error(err))
	}
	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	cloudPockets := []CloudPocketsResponse{}
	for rows.Next() {
		var p CloudPocketsResponse
		err = rows.Scan(&p.ID, &p.Name, &p.Category, &p.Currency, &p.Category)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		cloudPockets = append(cloudPockets, p)
	}
	return c.JSON(http.StatusOK, cloudPockets)

}
