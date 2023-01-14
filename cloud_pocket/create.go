package cloud_pocket

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

type PocketResponse struct {
	ID       string  `json:"id"`
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

//const (
//	cStmt         = "INSERT INTO accounts (balance) VALUES ($1) RETURNING id;"
//	cBalanceLimit = 10000
//)

//var (
//	hErrBalanceLimitExceed = echo.NewHTTPError(http.StatusBadRequest,
//		"create account balance exceed limitation")
//)

func (h handler) Create(c echo.Context) error {
	//logger := mlog.L(c)
	//ctx := c.Request().Context()
	var cp CloudPockets

	err := c.Bind(&cp)
	if err != nil {
		//logger.Error("bad request body", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, "bad request body", err.Error())
	}

	err = h.db.QueryRow("INSERT INTO expenses (title,amount,note,tags) values ($1,$2,$3,$4) RETURNING title,amount,note,tags,id", req.Title, req.Amount, req.Note, pq.Array(req.Tags)).
		Scan(&res.Title, &res.Amount, &res.Note, pq.Array(&res.Tags), &res.ID)

	pr := PocketResponse{
		ID:       "246810",
		Name:     "Travel Fund",
		Category: "Vacation",
		Currency: "THB",
		Balance:  100.00,
	}
	//logger.Info("create successfully", zap.Int64("id", 1))
	//cp.ID = 1
	return c.JSON(http.StatusCreated, pr)
}
