package cloud_pocket

import (
<<<<<<< HEAD
	"github.com/kkgo-software-engineering/workshop/mlog"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
=======
>>>>>>> f852357add66335a098c330ad6b9c780ec95e09c
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseCloudPockets struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance"`
}

type RequestCloudPockets struct {
	Name           string  `json:"name"`
	Currency       string  `json:"currency"`
	InitialBalance float64 `json:"initial_balance"`
	Category       string  `json:"category"`
}

const (
	cStmt         = "INSERT INTO cloud_pockets (name,balance,currency,category) values ($1,$2,$3,$4) RETURNING name,balance,currency,category,id;"
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


	err = h.db.QueryRow(cStmt, req.Name, req.InitialBalance, req.Currency, req.Category).
		Scan(&res.Name, &res.Balance, &res.Currency, &res.Category, &res.ID)


	logger.Info("create successfully", zap.Int64("id", res.ID))
	return c.JSON(http.StatusCreated, res)
}
