package dto

type AuthLoginRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required,min=5"`
}

type AuthLoginResponse struct {
	Token string `json:"token"`
}
