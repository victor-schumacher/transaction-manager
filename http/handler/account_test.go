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

const accountJSON = `{"account_id":"1fe89750-676e-47ca-949d-f97e56111a02", "document_number":"09365523859"}`

func TestCreateAccount(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/accounts", strings.NewReader(accountJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewAccount(mock.AccountRepo{}, e)

	table := []struct {
		name   string
		body   string
		method string
	}{
		{
			name:   "success case",
			body:   accountJSON,
			method: http.MethodPost,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			if assert.NoError(t, h.createNew(c)) {
				assert.Equal(t, http.StatusCreated, rec.Code)
				assert.Equal(t, "{\"message\":\"account successfully created\"}\n", rec.Body.String())
			}
		})
	}


}

func TestFindOneAccount(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/accounts/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewAccount(mock.AccountRepo{}, e)

	// Assertions
	if assert.NoError(t, h.findById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"account_id\":\"1\"}\n", rec.Body.String())
	}
}
