package handler

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const accountJSON = `{"account_id":"1", "document_number":"09365523859"}`

// TODO method not allowed
func TestCreateAccount(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/accounts", strings.NewReader(accountJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewAccount()

	if assert.NoError(t, h.createNew(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "{\"account_id\":\"1\"}\n", rec.Body.String())
	}
}

func TestFindOneAccount(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/accounts/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewAccount()

	// Assertions
	if assert.NoError(t, h.findById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"account_id\":\"1\"}\n", rec.Body.String())
	}
}
