package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password,omitempty" form:"password" validate:"required_without=id"`
	Token    string `json:"token,omitempty" form:"token"`
}

func (u User) AsResponse() map[string]any {
	return map[string]any{
		"id":         u.ID,
		"fullname":   u.Fullname,
		"username":   u.Username,
		"email":      u.Email,
		"created_at": u.CreatedAt,
		"updated_at": u.UpdatedAt,
	}
}
