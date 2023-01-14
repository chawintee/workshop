package cloud_pocket

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kkgo-software-engineering/workshop/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetById(t *testing.T) {
	tests := []struct {
		name    string
		cfgFlag config.FeatureFlag
		sqlFn   func() (*sql.DB, error)
		reqBody string
		wantErr error
	}{}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/cloud-pockets", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// db, _ := tc.sqlFn()
			h := New(tc.cfgFlag, nil)

			resp := h.GetById(c)
			// Assertions
			assert.Equal(t, http.StatusOK, resp)
		})
	}
}
