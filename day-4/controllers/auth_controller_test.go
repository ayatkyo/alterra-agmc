package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ayatkyo/alterra-agcm/day-4/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func authTestContextWithBody(body string) (echo.Context, *httptest.ResponseRecorder) {
	e, _ := utils.SetupTest()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	return c, res
}

func TestLoginSuccess(t *testing.T) {
	c, res := authTestContextWithBody(`{
		"username": "ayat",
		"password": "rahasia"
	}`)

	if assert.NoError(t, AuthLogin(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
	}
}

func TestLoginInvalid(t *testing.T) {
	c, res := authTestContextWithBody(`{
		"username": "random",
		"password": "random"
	}`)

	if assert.NoError(t, AuthLogin(c)) {
		assert.Equal(t, http.StatusBadRequest, res.Code)
	}
}

func TestLoginWrongPass(t *testing.T) {
	c, res := authTestContextWithBody(`{
		"username": "ayat",
		"password": "inisalah"
	}`)

	if assert.NoError(t, AuthLogin(c)) {
		assert.Equal(t, http.StatusBadRequest, res.Code)
	}
}

func TestLoginValidation(t *testing.T) {
	c, res := authTestContextWithBody(``)

	if assert.NoError(t, AuthLogin(c)) {
		assert.Equal(t, http.StatusBadRequest, res.Code)
	}
}
