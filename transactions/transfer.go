package transactions

import (
	"net/http"

	"github.com/kkgo-software-engineering/workshop/mlog"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

const (
	sStmt       = "UPDATE cloud_pockets;"
	dStmt       = "UPDATE cloud_pockets;"
	historyStmt = "INSERT INTO transactions (source_cloud_pocket_id, destination_cloud_pocket_id, amount, description) VALUES ($1, $2, $3, $4) RETURNING transaction_id;"
)

func (h handler) Transfer(c echo.Context) error {
	logger := mlog.L(c)
	ctx := c.Request().Context()
	var t Transaction
	err := c.Bind(&t)
	if err != nil {
		logger.Error("bad request body", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, "bad request body", err.Error())
	}

	var lastInsertId int64
	err = h.db.QueryRowContext(ctx, historyStmt, t.SourceCloudPocketID, t.DestinationCloudPocketID, t.Amount, t.Desciption).Scan(&lastInsertId)
	if err != nil {
		logger.Error("query row error", zap.Error(err))
		return err
	}

	logger.Info("transfer successfully", zap.Int64("transaction_id", lastInsertId))
	t.TransactionID = lastInsertId
	return c.JSON(http.StatusCreated, t)
}
