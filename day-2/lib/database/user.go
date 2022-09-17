package database

import (
	"github.com/ayatkyo/alterra-agcm/day-2/config"
	"github.com/ayatkyo/alterra-agcm/day-2/models"
)

func GetUsers() (any, error) {
	var users []models.User

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}

	return users, nil
}
