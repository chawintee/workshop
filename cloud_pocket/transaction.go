package pocket

// GET /cloud-pockets/:id/transactions

// POST /cloud-pockets/transfer

type Transaction struct {
	TransactionID            int64   `json:"transaction_id"`
	SourceCloudPocketID      int64   `json:"source_cloud_pocket_id"`
	DestinationCloudPocketID int64   `json:"destination_cloud_pocket_id"`
	Amount                   float64 `json:"amount"`
	Desciption               string  `json:"description"`
	Status                   string  `json:"status"`
}

type TransactionErr struct {
	ErrorMessage string `json:"error_message"`
	Status       string `json:"status"`
}
