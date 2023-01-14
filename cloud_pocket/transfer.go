package cloud_pocket

import (
	"net/http"

	"github.com/kkgo-software-engineering/workshop/mlog"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

const (
	balanceStmt = "SELECT balance from cloud_pockets WHERE id $1;"
	sStmt       = "UPDATE cloud_pockets SET balance = (balance - $2) WHERE id = $1 RETURNING *;"
	dStmt       = "UPDATE cloud_pockets SET balance = (balance + $2) WHERE id = $1 RETURNING *;"
	historyStmt = "INSERT INTO transactions (source_cloud_pocket_id, destination_cloud_pocket_id, amount, description, status) VALUES ($1, $2, $3, $4, $5) RETURNING transaction_id;"
)

var (
	hErrNotEnoughBalance = echo.NewHTTPError(http.StatusBadRequest,
		"Not enough balance in the source cloud pocket")
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

	var balance float64
	err = h.db.QueryRowContext(ctx, sStmt, t.Amount, t.SourceCloudPocketID).Scan(&balance)
	if err != nil {
		logger.Error("query row error", zap.Error(err))
		return err
	}

	if balance < t.Amount {
		logger.Error("Not enough balance in the source cloud pocket", zap.Error(hErrNotEnoughBalance))
		return hErrNotEnoughBalance
	}

	var sourcePocket CloudPocket
	err = h.db.QueryRowContext(ctx, sStmt, t.Amount, t.SourceCloudPocketID).Scan(&sourcePocket)
	if err != nil {
		logger.Error("query row error", zap.Error(err))
		return err
	}

	var destinationPocket CloudPocket
	err = h.db.QueryRowContext(ctx, dStmt, t.Amount, t.DestinationCloudPocketID).Scan(&destinationPocket)
	if err != nil {
		logger.Error("query row error", zap.Error(err))
		return err
	}

	var lastInsertId int64
	err = h.db.QueryRowContext(ctx, historyStmt, t.SourceCloudPocketID, t.DestinationCloudPocketID, t.Amount, t.Desciption, "Success").Scan(&lastInsertId)
	if err != nil {
		logger.Error("query row error", zap.Error(err))
		return err
	}

	logger.Info("transfer successfully", zap.Int64("transaction_id", lastInsertId))
	t.TransactionID = lastInsertId
	return c.JSON(http.StatusCreated, t)
}
