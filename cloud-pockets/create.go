package cloud_pockets

import (
	"database/sql"
	"github.com/kkgo-software-engineering/workshop/config"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance"`
}

type CloudPockets struct {
	Name           string  `json:"name"`
	Currency       string  `json:"currency"`
	InitialBalance float64 `json:"initial_balance"`
}

type handler struct {
	cfg config.FeatureFlag
	db  *sql.DB
}

func New(cfgFlag config.FeatureFlag, db *sql.DB) *handler {
	return &handler{cfgFlag, db}
}

const (
	cStmt         = "INSERT INTO accounts (balance) VALUES ($1) RETURNING id;"
	cBalanceLimit = 10000
)

var (
	hErrBalanceLimitExceed = echo.NewHTTPError(http.StatusBadRequest,
		"create account balance exceed limitation")
)

func (h handler) Create(c echo.Context) error {
	//logger := mlog.L(c)
	//ctx := c.Request().Context()
	var cp CloudPockets
	err := c.Bind(&cp)
	if err != nil {
		//logger.Error("bad request body", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, "bad request body", err.Error())
	}

	//logger.Info("create successfully", zap.Int64("id", 1))
	//cp.ID = 1
	return c.JSON(http.StatusCreated, cp)
}
