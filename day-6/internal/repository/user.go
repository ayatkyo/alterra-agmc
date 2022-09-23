package repository

import (
	"context"

	"github.com/ayatkyo/alterra-agcm/day-6/internal/dto"
	"github.com/ayatkyo/alterra-agcm/day-6/internal/models"
	"gorm.io/gorm"
)

type User interface {
	FirstByUsernameOrEmail(c context.Context, username string) (models.User, error)
	IsUnique(c context.Context, username string, email string, ignoreID int) bool
	Store(c context.Context, data *dto.AuthSignupRequest) (models.User, error)
}

type user struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *user {
	return &user{db}
}

func (r *user) FirstByUsernameOrEmail(c context.Context, username string) (models.User, error) {
	var user models.User
	err := r.DB.WithContext(c).Where("username = @username OR email = @username", map[string]any{"username": username}).First(&user).Error
	return user, err
}

func (r *user) IsUnique(c context.Context, username string, email string, ignoreID int) bool {
	q := "(username = @username OR email = @email)"
	if ignoreID != 0 {
		q += " AND id != @id"
	}

	var count int64
	if err := r.DB.Model(&models.User{}).Where(q, map[string]any{
		"username": username,
		"email":    email,
		"id":       ignoreID,
	}).Count(&count).Error; err != nil {
		return false
	}

	return count == 0
}

func (r *user) Store(c context.Context, data *dto.AuthSignupRequest) (models.User, error) {
	newUser := models.User{
		Fullname: data.Fullname,
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	}

	if err := r.DB.Create(&newUser).Error; err != nil {
		return newUser, err
	}

	return newUser, nil
}
