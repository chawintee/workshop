//go:build unit

package pocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kkgo-software-engineering/workshop/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	req = CloudPocketsRequest{
		Name:           "Test name",
		Currency:       "THB",
		InitialBalance: 400.00,
		Category:       "Vacation",
	}
)

func TestCreateCloudPocketsSuccess(t *testing.T) {
	// Mock
	db, mock, _ := sqlmock.New()
	createStmt := `INSERT INTO cloud_pockets (name,balance,currency,category, account_id)
	                 values ($1,$2,$3,$4,$5)
	                 RETURNING name,balance,currency,category,id;`
	mockedRow := sqlmock.NewRows([]string{"name", "balance", "currency", "category", "id"}).
		AddRow("test name", 100.00, "THB", "food", 1)
	//&res.Name, &res.Balance, &res.Currency, &res.Category, &res.ID

	mock.ExpectQuery(regexp.QuoteMeta(createStmt)).
		WithArgs(req.Name, req.InitialBalance, req.Currency, req.Category, 1).
		WillReturnRows((mockedRow))

	// Setup
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := New(config.FeatureFlag{}, db)

	// Assertions
	if assert.NoError(t, h.CreateCloudPockets(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}
