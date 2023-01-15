//go:build unit

package pocket

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kkgo-software-engineering/workshop/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetById(t *testing.T) {
	tests := []struct {
		name       string
		cfgFlag    config.FeatureFlag
		sqlFn      func() (*sql.DB, error)
		id         string
		wantStatus int
		wantBody   string
	}{
		{"get pocket detail success fully",
			config.FeatureFlag{},
			func() (*sql.DB, error) {
				db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
				if err != nil {
					return nil, err
				}
				row := sqlmock.NewRows([]string{"id", "balance", "name", "category", "currency"}).AddRow(1, 100.00, "Junk food", "food", "THB")
				mock.ExpectQuery(getDetailStmt).WithArgs(1).WillReturnRows(row)
				return db, err
			},
			`1`,
			http.StatusOK,
			`{"id": 1, "balance": 100.00, "name": "Junk food", "category": "food", "currency": "THB"}`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/cloud-pockets/", strings.NewReader(""))

			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(tc.id)

			db, _ := tc.sqlFn()
			h := New(tc.cfgFlag, db)

			if assert.NoError(t, h.GetById(c)) {
				assert.Equal(t, tc.wantStatus, rec.Code)
			}
		})
	}
}

func TestGetById_No_Param(t *testing.T) {
	tests := []struct {
		name       string
		cfgFlag    config.FeatureFlag
		sqlFn      func() (*sql.DB, error)
		id         string
		wantStatus int
		wantBody   string
	}{
		{"test by id",
			config.FeatureFlag{},
			func() (*sql.DB, error) {
				db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
				if err != nil {
					return nil, err
				}
				row := sqlmock.NewRows([]string{"id", "balance", "name", "category", "currency"}).AddRow(1, 100.00, "Junk food", "food", "THB")
				mock.ExpectQuery(getDetailStmt).WithArgs(1).WillReturnRows(row)
				return db, err
			},
			`1`,
			http.StatusCreated,
			`{"id": 1, "balance": 100.00, "name": "Junk food", "category": "food", "currency": "THB"}`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/cloud-pockets/", strings.NewReader(""))

			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues("")

			// db, _ := tc.sqlFn()
			h := New(tc.cfgFlag, nil)

			if assert.NoError(t, h.GetById(c)) {
				assert.Equal(t, http.StatusBadRequest, rec.Code)
			}
		})
	}
}
