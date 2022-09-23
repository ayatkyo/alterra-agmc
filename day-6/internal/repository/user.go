package repository

import (
	"context"

	"github.com/ayatkyo/alterra-agcm/day-6/internal/models"
	"gorm.io/gorm"
)

type User interface {
	FirstByUsernameOrEmail(ctx context.Context, username string) (models.User, error)
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
