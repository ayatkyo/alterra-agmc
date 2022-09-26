package constants

import "github.com/spf13/viper"

var JWT_SECRET = viper.GetString("JWT_SECRET")
