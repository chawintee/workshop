package cloud_pocket

////go:build unitget

import (

	// "errors"

	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kkgo-software-engineering/workshop/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	entity1 := &PocketResponse{
		ID:       1,
		Name:     "test-name",
		Balance:  100.00,
		Category: "test-category",
		Currency: "test-currency"}
	entity2 := &PocketResponse{
		ID:       2,
		Name:     "test-name",
		Balance:  100.00,
		Category: "test-category",
		Currency: "test-currency"}
	entities := []*PocketResponse{}
	entities = append(entities, entity1)
	entities = append(entities, entity2)
	newsMockRows := sqlmock.NewRows([]string{"id", "name", "balance", "category", "currency"}).
		AddRow(entity1.ID, entity1.Name, entity1.Balance, entity1.Category, entity1.Currency).
		AddRow(entity2.ID, entity2.Name, entity2.Balance, entity2.Category, entity2.Currency)
	mock.ExpectPrepare(regexp.QuoteMeta("SELECT id, name, balance, category, currency FROM cloud_pockets")).ExpectQuery().WillReturnRows(newsMockRows)
	// entitiesJson, err := json.Marshal(entities)
	// assert.NoError(t, err)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/cloud-pockets", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// db, err := tc.sqlFn()
	h := New(config.FeatureFlag{}, db)
	// Assertions
	assert.NoError(t, err)
	if assert.NoError(t, h.GetAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// assert.JSONEq(t, string(entitiesJson), rec.Body.String())
	}
}
