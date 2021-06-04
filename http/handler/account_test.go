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


	table := []struct {
		name         string
		body         string
		method       string
		responseBody string
		status       int
		wantErr      bool
	}{
		{
			name:         "success case",
			body:         accountJSON,
			method:       http.MethodPost,
			responseBody: "{\"message\":\"account successfully created\"}\n",
			status:       http.StatusCreated,
			wantErr:      false,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/accounts", strings.NewReader(tt.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := NewAccount(mock.AccountRepo{}, e)
			if !tt.wantErr {
				assert.NoError(t, h.createNew(c))
			}
			assert.Equal(t, tt.status, rec.Code)
			assert.Equal(t, tt.responseBody, rec.Body.String())
		})
	}

}