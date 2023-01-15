package pocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

var (
	req = CloudPocketsRequest{
		Name:           "Test name",
		Currency:       "THB",
		InitialBalance: 400.00,
		Category:       "Vacation",
	}
)

func TestAddPocket(t *testing.T) {
	// Mock
	db, mock, _ := sqlmock.New()
	createStmt := `INSERT INTO cloud_pockets (name,balance,currency,category, account_id)
                     values ($1,$2,$3,$4,$5)
                     RETURNING name,balance,currency,category,id;`
	mockedRow := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery(regexp.QuoteMeta(createStmt)).
		WithArgs(req.Name, req.InitialBalance, req.Currency, req.Category, 1).
		WillReturnRows((mockedRow))

	// Setup
	b, err := json.Marshal(expense)
	if err != nil {
		fmt.Println(err)
		return
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewHandler(db)

	// Assertions
	if assert.NoError(t, h.CreateExpenseHandler(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}
