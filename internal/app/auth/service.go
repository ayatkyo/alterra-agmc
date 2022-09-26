package auth

import (
	"context"
	"errors"

	"github.com/ayatkyo/alterra-agcm/day-10/internal/dto"
	"github.com/ayatkyo/alterra-agcm/day-10/internal/factory"
	"github.com/ayatkyo/alterra-agcm/day-10/internal/repository"
	"github.com/ayatkyo/alterra-agcm/day-10/pkg/utils"
)

type service struct {
	UserRepository repository.User
}

type Service interface {
	Login(c context.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error)
	Signup(c context.Context, payload *dto.AuthSignupRequest) (*dto.AuthSignupReponse, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		UserRepository: f.UserRepository,
	}
}

func (s *service) Login(c context.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error) {
	user, err := s.UserRepository.FirstByUsernameOrEmail(c, payload.Username)
	if err != nil {
		return nil, err
	}

	if !utils.BcryptCheck(payload.Password, user.Password) {
		return nil, errors.New("username or password is incorrect")
	}

	// create token
	token, err := utils.JWTCreateToken(int(user.ID))
	if err != nil {
		return nil, errors.New("cannot create token")
	}

	res := &dto.AuthLoginResponse{
		Token: token,
	}

	return res, nil
}

func (s *service) Signup(c context.Context, payload *dto.AuthSignupRequest) (*dto.AuthSignupReponse, error) {
	// check unique email and username
	if !s.UserRepository.IsUnique(c, payload.Username, payload.Email, 0) {
		return nil, errors.New("username or email already exists")
	}

	// Create password hash
	pass, err := utils.BcryptMake(payload.Password)
	if err != nil {
		return nil, err
	}

	// Set password hash
	payload.Password = pass

	// create user
	user, err := s.UserRepository.Store(c, payload)
	if err != nil {
		return nil, err
	}

	res := &dto.AuthSignupReponse{
		ID:        user.ID,
		Fullname:  user.Fullname,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return res, nil
}
