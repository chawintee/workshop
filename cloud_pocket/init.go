package pocket

import (
	"database/sql"

	"github.com/kkgo-software-engineering/workshop/config"
)

type CloudPocketsResponse struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance"`
}

type CloudPocketsRequest struct {
	Name           string  `json:"name"`
	Currency       string  `json:"currency"`
	InitialBalance float64 `json:"initial_balance"`
	Category       string  `json:"category"`
}

type handler struct {
	cfg config.FeatureFlag
	db  *sql.DB
}

func New(cfgFlag config.FeatureFlag, db *sql.DB) *handler {
	return &handler{cfgFlag, db}
}

func AddFloat(a, b float64) float64 {
	x := a * 100
	y := b * 100
	z := x + y
	z = z / 100

	return z
}

func MinusFloat(a, b float64) float64 {
	x := a * 100
	y := b * 100
	z := x - y
	z = z / 100

	return z
}
