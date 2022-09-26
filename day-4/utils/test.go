package utils

import (
	"encoding/json"
	"net/http/httptest"
	"strings"

	"github.com/ayatkyo/alterra-agmc/day-4/config"
	"github.com/ayatkyo/alterra-agmc/day-4/models"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupTest() (*echo.Echo, *gorm.DB) {
	var err error
	DB := config.DB

	// Setup test database
	if DB == nil {
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
		pass, _ := BcryptMake("rahasia")
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
	}

	// Setup echo
	e := echo.New()
	e.Validator = EchoCustomValidator

	return e, DB
}

func CreateTestContext(method string, path string, body string) (echo.Context, *httptest.ResponseRecorder, *echo.Echo) {
	e, _ := SetupTest()

	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	return c, res, e
}

func ParseResponseJSON(body string) map[string]any {
	res := map[string]any{}

	json.Unmarshal([]byte(body), &res)

	return res
}
