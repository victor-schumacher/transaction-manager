package handler

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"transaction-manager/internal/mock"
)

func TestCreateTransaction(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/transactions", strings.NewReader(accountJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	transacation := NewTransaction(mock.TransactionRepo{}, e)

	if assert.NoError(t, transacation.createNew(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
	}
}
