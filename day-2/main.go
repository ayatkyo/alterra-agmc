package main

import (
	"fmt"
	"os"

	"github.com/ayatkyo/alterra-agcm/day-2/config"
	"github.com/ayatkyo/alterra-agcm/day-2/routes"
)

func main() {
	//	Initialize DB
	config.InitDB()
	config.MigrateDB()

	e := routes.New()
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))))
}
