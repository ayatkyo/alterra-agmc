package repository

import (
	"context"

	"github.com/ayatkyo/alterra-agmc/day-10/internal/dto"
	"github.com/ayatkyo/alterra-agmc/day-10/internal/models"
	"gorm.io/gorm"
)

type User interface {
	FindAll(c context.Context) ([]models.User, error)
	FirstByUsernameOrEmail(c context.Context, username string) (models.User, error)
	FindByID(c context.Context, ID uint) (models.User, error)
	Store(c context.Context, data *dto.AuthSignupRequest) (models.User, error)
	Update(c context.Context, user *models.User) error
	Destroy(c context.Context, ID uint) error
	IsUnique(c context.Context, username string, email string, ignoreID int) bool
}

type user struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) User {
	return &user{db}
}

func (r *user) FirstByUsernameOrEmail(c context.Context, username string) (models.User, error) {
	var user models.User
	err := r.DB.WithContext(c).Where("username = @username OR email = @username", map[string]any{"username": username}).First(&user).Error
	return user, err
}

func (r *user) FindAll(c context.Context) ([]models.User, error) {
	var users []models.User
	err := r.DB.WithContext(c).Model(&models.User{}).Find(&users).Error
	return users, err
}

func (r *user) FindByID(c context.Context, ID uint) (models.User, error) {
	var user models.User
	err := r.DB.WithContext(c).Where("id = ?", ID).First(&user).Error
	return user, err
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

func (r *user) Update(c context.Context, user *models.User) error {
	err := r.DB.WithContext(c).Save(&user).Error
	return err
}

func (r *user) Destroy(c context.Context, ID uint) error {
	err := r.DB.WithContext(c).Where("id = ?", ID).Delete(&models.User{}).Error
	return err
}

func (r *user) IsUnique(c context.Context, username string, email string, ignoreID int) bool {
	q := "(username = @username OR email = @email)"
	if ignoreID != 0 {
		q += " AND id != @id"
	}

	var count int64
	if err := r.DB.WithContext(c).Model(&models.User{}).Where(q, map[string]any{
		"username": username,
		"email":    email,
		"id":       ignoreID,
	}).Count(&count).Error; err != nil {
		return false
	}

	return count == 0
}
