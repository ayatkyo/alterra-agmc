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

func GetUserByID(ID int) (any, error) {
	var user models.User

	if e := config.DB.Where("id = ?", ID).First(&user).Error; e != nil {
		return nil, e
	}

	return user, nil
}

func CreateUser(user models.User) (any, error) {
	if e := config.DB.Create(&user).Error; e != nil {
		return nil, e
	}

	return user, nil
}
