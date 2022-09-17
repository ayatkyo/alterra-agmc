package config

import (
	"fmt"
	"os"

	"github.com/ayatkyo/alterra-agcm/day-2/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// 	Create connection string
	c := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?parseTime=True`,
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	//	Connect
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
