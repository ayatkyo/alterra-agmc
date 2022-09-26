package config

import (
	"fmt"

	"github.com/ayatkyo/alterra-agmc/day-3/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// Create connection string
	c := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?parseTime=True`,
		viper.Get("DB_USERNAME"),
		viper.Get("DB_PASSWORD"),
		viper.Get("DB_HOST"),
		viper.Get("DB_PORT"),
		viper.Get("DB_DATABASE"),
	)

	// Connect
	var err error
	DB, err = gorm.Open(mysql.Open(c), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return DB
}

func MigrateDB() {
	DB.AutoMigrate(&models.User{})
}
