package auth

import (
	"context"
	"errors"

	"github.com/ayatkyo/alterra-agcm/day-6/internal/dto"
	"github.com/ayatkyo/alterra-agcm/day-6/internal/factory"
	"github.com/ayatkyo/alterra-agcm/day-6/internal/repository"
	"github.com/ayatkyo/alterra-agcm/day-6/pkg/utils"
)

type service struct {
	UserRepository repository.User
}

type Service interface {
	Login(c context.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error)
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
		return nil, errors.New("Username or password is incorrect")
	}

	// create token
	token, err := utils.JWTCreateToken(int(user.ID))
	if err != nil {
		return nil, errors.New("Cannot create token")
	}

	res := &dto.AuthLoginResponse{
		Token: token,
	}

	return res, nil
}
