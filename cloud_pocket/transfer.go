package pocket

import (
	"net/http"

	"github.com/kkgo-software-engineering/workshop/mlog"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

const (
	bStmt = "SELECT balance from cloud_pockets WHERE id = $1;"
	sStmt = "UPDATE cloud_pockets SET balance = $2 WHERE id = $1"
	hStmt = "INSERT INTO transactions (source_cloud_pocket_id, destination_cloud_pocket_id, amount, description, status) VALUES ($1, $2, $3, $4, $5) RETURNING transaction_id, status;"
)

var (
	hErrNotEnoughBalance = echo.NewHTTPError(http.StatusBadRequest,
		TransactionErr{ErrorMessage: "Not enough balance in the source cloud pocket", Status: "Failed"})
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

	var sourceBalance float64
	err = h.db.QueryRowContext(ctx, bStmt, t.SourceCloudPocketID).Scan(&sourceBalance)
	if err != nil {
		logger.Error("query row error", zap.Error(err))
		return err
	}

	if sourceBalance < t.Amount {
		logger.Error("Not enough balance in the source cloud pocket", zap.Error(hErrNotEnoughBalance))
		return hErrNotEnoughBalance
	}

	var destBalance float64
	err = h.db.QueryRowContext(ctx, bStmt, t.DestinationCloudPocketID).Scan(&destBalance)
	if err != nil {
		logger.Error("query row error", zap.Error(err))
		return err
	}

	// Get a Tx for making transaction requests.
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		logger.Error("error", zap.Error(err))
		return err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	var lastInsertId int64
	minusAmount := MinusFloat(sourceBalance, t.Amount)
	tx.Prepare(sStmt)
	_, err = tx.ExecContext(ctx, sStmt, t.SourceCloudPocketID, minusAmount)
	if err != nil {
		logErr := tx.QueryRowContext(ctx, hStmt, t.SourceCloudPocketID, t.DestinationCloudPocketID, t.Amount, t.Desciption, "Failed").Scan(&lastInsertId, &t.Status)
		if logErr != nil {
			logger.Error("log error", zap.Error(logErr))
			return logErr
		}
		t.TransactionID = lastInsertId
		return c.JSON(http.StatusInternalServerError, TransactionErr{ErrorMessage: "update err", Status: t.Status})
	}

	newAmount := AddFloat(destBalance, t.Amount)
	_, err = tx.ExecContext(ctx, sStmt, t.DestinationCloudPocketID, newAmount)
	if err != nil {
		logErr := tx.QueryRowContext(ctx, hStmt, t.SourceCloudPocketID, t.DestinationCloudPocketID, t.Amount, t.Desciption, "Failed").Scan(&lastInsertId, &t.Status)
		if logErr != nil {
			logger.Error("log error", zap.Error(logErr))
			return logErr
		}
		t.TransactionID = lastInsertId
		return c.JSON(http.StatusInternalServerError, TransactionErr{ErrorMessage: "update err", Status: t.Status})
	}

	err = tx.QueryRowContext(ctx, hStmt, t.SourceCloudPocketID, t.DestinationCloudPocketID, t.Amount, t.Desciption, "Success").Scan(&lastInsertId, &t.Status)
	if err != nil {
		logger.Error("log error", zap.Error(err))
		return err
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		logger.Error("error", zap.Error(err))
		return err
	}

	logger.Info("transfer successfully", zap.Int64("transaction_id", lastInsertId))
	t.TransactionID = lastInsertId
	return c.JSON(http.StatusOK, t)
}
