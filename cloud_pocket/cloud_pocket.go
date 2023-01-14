package cloud_pocket

import (
	"database/sql"
	"net/http"

	"github.com/kkgo-software-engineering/workshop/config"
	"github.com/labstack/echo/v4"
)

type CloudPocket struct {
	ID       int64   `json:"id"`
	Balance  float64 `json:"balance"`
	Name     float64 `json:"name"`
	Category string  `json:"category"`
	Currency string  `json:"currency"`
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
