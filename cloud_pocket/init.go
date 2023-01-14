package cloud_pocket

import (
	"database/sql"
	"github.com/kkgo-software-engineering/workshop/config"
)

type handler struct {
	cfg config.FeatureFlag
	db  *sql.DB
}

func New(cfgFlag config.FeatureFlag, db *sql.DB) *handler {
	return &handler{cfgFlag, db}
}
