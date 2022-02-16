package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"yuno/echo-example/database"
	"yuno/echo-example/response"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func init() {
	database.ConnectPostgres()
}

func TestCreateUser(t *testing.T) {
	studentJSON := `{
		"first_name": "Test",
		"last_name": "Name",
		"born_date": "2000-11-14T19:03:44-04:56",
		"country_id": 1
	}`

	studentJSONResponse := `{
		"first_name": "Test",
		"last_name": "Name",
		"born_date": "2000-11-14T19:03:44-04:56",
		"country_name": "Colombia",
		"contry_iso": "CO"
	}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/student", strings.NewReader(studentJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)

	if assert.NoError(t, CreateStudent(context)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		
		// Marshall
		responseExpect := new(response.UserResponse)
		responseActual := new(response.UserResponse)
		json.Unmarshal([]byte(studentJSONResponse), &responseExpect)
		json.Unmarshal(rec.Body.Bytes(), &responseActual)

		assert.Equal(t, responseExpect.CountryName, responseActual.CountryName)
		assert.Equal(t, responseExpect.FirstName, responseActual.FirstName)
		assert.Equal(t, responseExpect.LastName, responseActual.LastName)
	}
}
