package dto

type UserUpdateRequest struct {
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password,omitempty" form:"password" validate:"omitempty,required,min=5"`
}
