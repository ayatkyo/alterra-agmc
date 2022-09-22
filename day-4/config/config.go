package config

import (
	"github.com/spf13/viper"
)

func init() {
	// Load config
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}
