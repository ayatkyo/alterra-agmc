package main

import "github.com/ayatkyo/alterra-agcm/day-2/routes"

func main() {
	e := routes.New()
	e.Logger.Fatal(e.Start(":15000"))
}
