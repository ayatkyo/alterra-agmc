package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ayatkyo/alterra-agcm/day-4/config"
	"github.com/ayatkyo/alterra-agcm/day-4/models"
	"github.com/ayatkyo/alterra-agcm/day-4/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	e  *echo.Echo
	DB *gorm.DB
)

func init() {
	var err error

	// Setup test database
	DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	// Migrate
	DB.AutoMigrate(&models.User{})

	// Replace DB
	config.DB = DB

	// Seed test data
	pass, _ := utils.BcryptMake("rahasia")
	var users = []models.User{
		{
			Fullname: "Hidayatullah",
			Email:    "ayat@gmail.com",
			Username: "ayat",
			Password: pass,
		},
		{
			Fullname: "Muhammad Al-Fath",
			Email:    "alfath@gmail.com",
			Username: "alfath",
			Password: pass,
		},
		{
			Fullname: "Siti Mawaddah",
			Email:    "mawad@gmail.com",
			Username: "mawad",
			Password: pass,
		},
	}
	DB.Create(&users)

	// Setup echo
	e = echo.New()
	e.Validator = utils.EchoCustomValidator
}

func createContextWithBody(body string) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	return c, res
}

func TestLoginSuccess(t *testing.T) {
	c, res := createContextWithBody(`{
		"username": "ayat",
		"password": "rahasia"
	}`)

	if assert.NoError(t, AuthLogin(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
	}
}

func TestLoginInvalid(t *testing.T) {
	c, res := createContextWithBody(`{
		"username": "random",
		"password": "random"
	}`)

	if assert.NoError(t, AuthLogin(c)) {
		assert.Equal(t, http.StatusBadRequest, res.Code)
	}
}

func TestLoginWrongPass(t *testing.T) {
	c, res := createContextWithBody(`{
		"username": "ayat",
		"password": "inisalah"
	}`)

	if assert.NoError(t, AuthLogin(c)) {
		assert.Equal(t, http.StatusBadRequest, res.Code)
	}
}

func TestLoginValidation(t *testing.T) {
	c, res := createContextWithBody(``)

	if assert.NoError(t, AuthLogin(c)) {
		assert.Equal(t, http.StatusBadRequest, res.Code)
	}
}
