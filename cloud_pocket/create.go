package cloud_pocket

import (
	"github.com/kkgo-software-engineering/workshop/mlog"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

const (
	createStmt = `INSERT INTO cloud_pockets (name,balance,currency,category, account_id) 
                      values ($1,$2,$3,$4,$5) 
                      RETURNING name,balance,currency,category,id;`
	cBalanceLimit = 10000
)

//var (
//	hErrBalanceLimitExceed = echo.NewHTTPError(http.StatusBadRequest,
//		"create account balance exceed limitation")
//)

func (h handler) Create(c echo.Context) error {
	logger := mlog.L(c)
	var req RequestCloudPockets
	var res ResponseCloudPockets

	err := c.Bind(&req)
	if err != nil {
		logger.Error("bad request body", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, "bad request body", err.Error())
	}

	err = h.db.QueryRow(createStmt, req.Name, req.InitialBalance, req.Currency, req.Category, 1).
		Scan(&res.Name, &res.Balance, &res.Currency, &res.Category, &res.ID)
	if err != nil {
		logger.Error("internal error", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, "Database Error ", err.Error())
	}

	logger.Info("create successfully", zap.Int64("id", res.ID))
	return c.JSON(http.StatusCreated, res)
}
