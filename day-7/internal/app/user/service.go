package user

import (
	"context"
	"errors"

	"github.com/ayatkyo/alterra-agcm/day-7/internal/dto"
	"github.com/ayatkyo/alterra-agcm/day-7/internal/factory"
	"github.com/ayatkyo/alterra-agcm/day-7/internal/models"
	"github.com/ayatkyo/alterra-agcm/day-7/internal/repository"
	"github.com/ayatkyo/alterra-agcm/day-7/pkg/utils"
)

type service struct {
	UserRepository repository.User
}

type Service interface {
	FindAll(c context.Context) ([]models.User, error)
	FindByID(c context.Context, ID uint) (models.User, error)
	Update(c context.Context, ID uint, payload *dto.UserUpdateRequest) (*models.User, error)
	Destroy(c context.Context, ID uint) error
}

func NewService(f *factory.Factory) Service {
	return &service{
		UserRepository: f.UserRepository,
	}
}

func (s *service) FindAll(c context.Context) ([]models.User, error) {
	users, err := s.UserRepository.FindAll(c)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *service) FindByID(c context.Context, ID uint) (models.User, error) {
	user, err := s.UserRepository.FindByID(c, ID)
	return user, err
}

func (s *service) Update(c context.Context, ID uint, payload *dto.UserUpdateRequest) (*models.User, error) {
	// Find user
	user, err := s.UserRepository.FindByID(c, ID)
	if err != nil {
		return nil, err
	}

	// check unique email and username
	if !s.UserRepository.IsUnique(c, payload.Username, payload.Email, int(user.ID)) {
		return nil, errors.New("username or email already exists")
	}

	// update data
	user.Fullname = payload.Fullname
	user.Email = payload.Email
	user.Username = payload.Username

	// check password change
	if payload.Password != "" {
		passwordHash, err := utils.BcryptMake(payload.Password)
		if err != nil {
			return nil, err
		}

		// set new password
		user.Password = passwordHash
	}

	err = s.UserRepository.Update(c, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *service) Destroy(c context.Context, ID uint) error {
	var user models.User

	// Find user
	user, err := s.UserRepository.FindByID(c, ID)
	if err != nil {
		return err
	}

	if user.ID == 0 {
		return errors.New("user not found")
	}

	err = s.UserRepository.Destroy(c, ID)
	return err
}
