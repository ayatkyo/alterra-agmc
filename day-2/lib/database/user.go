package database

import (
	"errors"

	"github.com/ayatkyo/alterra-agmc/day-2/config"
	"github.com/ayatkyo/alterra-agmc/day-2/models"
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

func UpdateUser(user models.User) (any, error) {
	if e := config.DB.Save(&user).Error; e != nil {
		return nil, e
	}

	return user, nil
}

func DeleteUser(ID int) error {
	var user models.User

	//	Find user
	e := config.DB.Where("id = ?", ID).First(&user).Error
	if e != nil {
		return e
	}

	if user.ID == 0 {
		return errors.New("User not found")
	}

	//	Delete
	e = config.DB.Delete(&user).Error
	if e != nil {
		return e
	}

	return nil
}
