package main

import (
	"fmt"

	"github.com/ayatkyo/alterra-agmc/day-10/database"
	"github.com/ayatkyo/alterra-agmc/day-10/internal/factory"
	"github.com/ayatkyo/alterra-agmc/day-10/internal/http"
	"github.com/ayatkyo/alterra-agmc/day-10/internal/middlewares"
	"github.com/ayatkyo/alterra-agmc/day-10/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	config.LoadConfig(".env")

	database.InitDB()
	database.MigrateDB()

	f := factory.NewFactory()
	e := echo.New()

	middlewares.LogMiddleware(e)

	http.NewHttp(e, f)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", viper.Get("APP_HOST"), viper.Get("APP_PORT"))))
}
