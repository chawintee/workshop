package transactions

import (
	"database/sql"

	"github.com/kkgo-software-engineering/workshop/config"
)

// GET /cloud-pockets/:id/transactions



// POST /cloud-pockets/transfer

type Transaction struct {
	TransactionID            int64   `json:"transaction_id"`
	SourceCloudPocketID      int64   `json:"source_cloud_pocket_id"`
	DestinationCloudPocketID int64   `json:"destination_cloud_pocket_id"`
	Amount                   float64 `json:"amount"`
	Desciption               string  `json:"description"`
}

type handler struct {
	cfg config.FeatureFlag
	db  *sql.DB
}

func New(cfgFlag config.FeatureFlag, db *sql.DB) *handler {
	return &handler{cfgFlag, db}
}
