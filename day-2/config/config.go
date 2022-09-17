package config

import "github.com/joho/godotenv"

func init() {
	// 	Load config
	godotenv.Load(".env")
}
