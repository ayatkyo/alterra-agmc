package main

import (
	"fmt"

	"github.com/ayatkyo/alterra-agcm/day-3/config"
	"github.com/ayatkyo/alterra-agcm/day-3/middlewares"
	"github.com/ayatkyo/alterra-agcm/day-3/routes"
	"github.com/spf13/viper"
)

func main() {
	// Initialize DB
	config.InitDB()
	config.MigrateDB()

	// Before route middlewares

	// Setup routes
	e := routes.New()

	// After route middlewares
	middlewares.LogMiddlewares(e)

	// Start App
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", viper.Get("APP_HOST"), viper.Get("APP_PORT"))))
}
