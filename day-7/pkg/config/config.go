package config

import (
	"github.com/spf13/viper"
)

func LoadConfig(path string) {
	// Load config
	viper.SetConfigFile(path)
	viper.ReadInConfig()
	viper.AutomaticEnv()
}
