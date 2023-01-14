package cloud_pocket

import (
	"database/sql"
	"github.com/kkgo-software-engineering/workshop/config"
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
